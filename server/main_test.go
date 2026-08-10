package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type recordingSender struct {
	requests []contactRequest
	err      error
}

func (s *recordingSender) Send(_ context.Context, request contactRequest) error {
	s.requests = append(s.requests, request)
	return s.err
}

func testConfig() config {
	return config{
		staticDir:      "../site",
		smtpHost:       "smtp.example.com",
		smtpPort:       "465",
		smtpUser:       "sender@example.com",
		smtpPassword:   "test-only",
		smtpFrom:       "sender@example.com",
		smtpFromName:   "WeBuildit",
		recipient:      "recipient@example.com",
		allowedOrigins: map[string]struct{}{"https://webuild-it.com": {}},
	}
}

func validPayload() string {
	return `{"name":"Ana Example","email":"ana@example.com","company":"Example","projectType":"ai-automation","message":"We need a practical automation for our workflow.","language":"en","website":""}`
}

func TestContactAcceptsValidatedRequest(t *testing.T) {
	sender := &recordingSender{}
	request := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(validPayload()))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Origin", "https://webuild-it.com")
	response := httptest.NewRecorder()

	newHandler(testConfig(), sender).ServeHTTP(response, request)

	if response.Code != http.StatusAccepted {
		t.Fatalf("status = %d, body = %s", response.Code, response.Body.String())
	}
	if len(sender.requests) != 1 || sender.requests[0].Email != "ana@example.com" {
		t.Fatalf("mail sender received %#v", sender.requests)
	}
}

func TestContactRejectsHeaderInjectionAndShortMessage(t *testing.T) {
	sender := &recordingSender{}
	payload := `{"name":"Ana\r\nBcc: attacker@example.com","email":"ana@example.com","projectType":"other","message":"too short","language":"en","website":""}`
	request := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(payload))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	newHandler(testConfig(), sender).ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, body = %s", response.Code, response.Body.String())
	}
	var body struct {
		Fields map[string]string `json:"fields"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body.Fields["name"] != "invalid" || body.Fields["message"] != "invalid" {
		t.Fatalf("fields = %#v", body.Fields)
	}
	if len(sender.requests) != 0 {
		t.Fatal("invalid request was sent")
	}
}

func TestContactRejectsControlCharacters(t *testing.T) {
	sender := &recordingSender{}
	payload := `{"name":"Ana Example","email":"ana@example.com","projectType":"other","message":"A valid-looking message with a null character: \u0000","language":"en","website":""}`
	request := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(payload))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	newHandler(testConfig(), sender).ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest || len(sender.requests) != 0 {
		t.Fatalf("status = %d, requests = %d", response.Code, len(sender.requests))
	}
}

func TestContactRejectsOversizedBody(t *testing.T) {
	sender := &recordingSender{}
	request := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(`{"name":"`+strings.Repeat("a", maxRequestBytes)+`"}`))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	newHandler(testConfig(), sender).ServeHTTP(response, request)

	if response.Code != http.StatusRequestEntityTooLarge || len(sender.requests) != 0 {
		t.Fatalf("status = %d, body = %s", response.Code, response.Body.String())
	}
}

func TestContactHoneypotAcceptsWithoutDelivery(t *testing.T) {
	sender := &recordingSender{}
	payload := strings.Replace(validPayload(), `"website":""`, `"website":"spam.example"`, 1)
	request := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(payload))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	newHandler(testConfig(), sender).ServeHTTP(response, request)

	if response.Code != http.StatusAccepted || len(sender.requests) != 0 {
		t.Fatalf("status = %d, requests = %d", response.Code, len(sender.requests))
	}
}

func TestContactRejectsUntrustedOrigin(t *testing.T) {
	sender := &recordingSender{}
	request := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(validPayload()))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Origin", "https://attacker.example")
	response := httptest.NewRecorder()

	newHandler(testConfig(), sender).ServeHTTP(response, request)

	if response.Code != http.StatusForbidden || len(sender.requests) != 0 {
		t.Fatalf("status = %d, requests = %d", response.Code, len(sender.requests))
	}
}

func TestContactReturnsUnavailableWithoutSecrets(t *testing.T) {
	cfg := testConfig()
	cfg.smtpPassword = ""
	request := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(validPayload()))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	newHandler(cfg, &recordingSender{}).ServeHTTP(response, request)

	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, body = %s", response.Code, response.Body.String())
	}
}

func TestContactRateLimitsByClient(t *testing.T) {
	sender := &recordingSender{}
	handler := newHandler(testConfig(), sender)
	for attempt := 1; attempt <= rateLimitCount+1; attempt++ {
		request := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(validPayload()))
		request.Header.Set("Content-Type", "application/json")
		request.RemoteAddr = "192.0.2.10:1234"
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		if attempt <= rateLimitCount && response.Code != http.StatusAccepted {
			t.Fatalf("attempt %d status = %d", attempt, response.Code)
		}
		if attempt == rateLimitCount+1 && response.Code != http.StatusTooManyRequests {
			t.Fatalf("rate-limited status = %d", response.Code)
		}
	}
	if len(sender.requests) != rateLimitCount {
		t.Fatalf("deliveries = %d", len(sender.requests))
	}
}

func TestStaticHandlerAndSecurityHeaders(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	newHandler(testConfig(), &recordingSender{}).ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d", response.Code)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), "WeBuildit") {
		t.Fatal("landing page was not served")
	}
	if response.Header().Get("Content-Security-Policy") == "" {
		t.Fatal("content security policy missing")
	}
}

func TestStaticHandlerServesWebManifestWithManifestContentType(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/assets/brand/site.webmanifest", nil)
	response := httptest.NewRecorder()
	newHandler(testConfig(), &recordingSender{}).ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d", response.Code)
	}
	if contentType := response.Header().Get("Content-Type"); contentType != "application/manifest+json" {
		t.Fatalf("content type = %q", contentType)
	}
}

func TestBuildMessageUsesFixedEnvelopeAndSafeReplyTo(t *testing.T) {
	request := contactRequest{
		Name: "Ana Example", Email: "ana@example.com", Company: "Example",
		ProjectType: "digital-product", Message: "Please help us build a useful product.", Language: "ro",
	}
	message := buildMessage(testConfig(), request)
	if !strings.Contains(message, "Reply-To: \"Ana Example\" <ana@example.com>") {
		t.Fatalf("reply-to missing: %s", message)
	}
	if strings.Contains(message, "Bcc:") {
		t.Fatal("unexpected blind-copy header")
	}
}
