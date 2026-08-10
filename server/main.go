package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"net/mail"
	"net/smtp"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	maxRequestBytes = 32 << 10
	rateLimitCount  = 5
	rateLimitWindow = 10 * time.Minute
)

var projectTypes = map[string]string{
	"ai-automation":         "AI / Automation",
	"digital-product":       "Website / Digital Product",
	"infrastructure-devops": "Infrastructure / DevOps",
	"security-operations":   "Security / Operations",
	"other":                 "Other",
}

type config struct {
	addr           string
	staticDir      string
	smtpHost       string
	smtpPort       string
	smtpUser       string
	smtpPassword   string
	smtpFrom       string
	smtpFromName   string
	recipient      string
	allowedOrigins map[string]struct{}
	trustProxy     bool
}

func loadConfig() config {
	cfg := config{
		addr:           envOrDefault("LISTEN_ADDR", ":8080"),
		staticDir:      envOrDefault("STATIC_DIR", "site"),
		smtpHost:       strings.TrimSpace(os.Getenv("SMTP_HOST")),
		smtpPort:       envOrDefault("SMTP_PORT", "465"),
		smtpUser:       strings.TrimSpace(os.Getenv("SMTP_USER")),
		smtpPassword:   os.Getenv("SMTP_PASSWORD"),
		smtpFrom:       strings.TrimSpace(os.Getenv("SMTP_FROM_ADDRESS")),
		smtpFromName:   envOrDefault("SMTP_FROM_NAME", "WeBuildit"),
		recipient:      strings.TrimSpace(os.Getenv("CONTACT_RECIPIENT")),
		allowedOrigins: parseOrigins(os.Getenv("CONTACT_ALLOWED_ORIGINS")),
		trustProxy:     strings.EqualFold(strings.TrimSpace(os.Getenv("TRUST_PROXY_HEADERS")), "true"),
	}
	if cfg.smtpFrom == "" {
		cfg.smtpFrom = cfg.smtpUser
	}
	return cfg
}

func envOrDefault(name, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(name)); value != "" {
		return value
	}
	return fallback
}

func parseOrigins(value string) map[string]struct{} {
	origins := make(map[string]struct{})
	for _, origin := range strings.Split(value, ",") {
		origin = strings.TrimSuffix(strings.TrimSpace(origin), "/")
		if origin != "" {
			origins[origin] = struct{}{}
		}
	}
	return origins
}

func (c config) contactConfigured() bool {
	return c.smtpHost != "" && c.smtpUser != "" && c.smtpPassword != "" &&
		validEmail(c.smtpFrom) && validEmail(c.recipient) && !hasHeaderControl(c.smtpFromName)
}

type contactRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	ProjectType string `json:"projectType"`
	Message     string `json:"message"`
	Language    string `json:"language"`
	Website     string `json:"website"`
}

type mailSender interface {
	Send(context.Context, contactRequest) error
}

type smtpSender struct {
	cfg config
}

func (s smtpSender) Send(ctx context.Context, request contactRequest) error {
	address := net.JoinHostPort(s.cfg.smtpHost, s.cfg.smtpPort)
	dialer := &net.Dialer{Timeout: 10 * time.Second}
	raw, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return fmt.Errorf("connect to mail relay: %w", err)
	}
	defer raw.Close()

	tlsConn := tls.Client(raw, &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: s.cfg.smtpHost,
	})
	if err := tlsConn.HandshakeContext(ctx); err != nil {
		return fmt.Errorf("establish mail TLS: %w", err)
	}

	client, err := smtp.NewClient(tlsConn, s.cfg.smtpHost)
	if err != nil {
		return fmt.Errorf("create SMTP client: %w", err)
	}
	defer client.Close()

	if err := client.Auth(smtp.PlainAuth("", s.cfg.smtpUser, s.cfg.smtpPassword, s.cfg.smtpHost)); err != nil {
		return fmt.Errorf("authenticate with mail relay: %w", err)
	}
	if err := client.Mail(s.cfg.smtpFrom); err != nil {
		return fmt.Errorf("set envelope sender: %w", err)
	}
	if err := client.Rcpt(s.cfg.recipient); err != nil {
		return fmt.Errorf("set delivery recipient: %w", err)
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("start SMTP body: %w", err)
	}
	message := buildMessage(s.cfg, request)
	if _, err := io.WriteString(writer, message); err != nil {
		writer.Close()
		return fmt.Errorf("write SMTP body: %w", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("finish SMTP body: %w", err)
	}
	if err := client.Quit(); err != nil {
		return fmt.Errorf("finish SMTP session: %w", err)
	}
	return nil
}

func buildMessage(cfg config, request contactRequest) string {
	from := (&mail.Address{Name: cfg.smtpFromName, Address: cfg.smtpFrom}).String()
	to := (&mail.Address{Address: cfg.recipient}).String()
	replyTo := (&mail.Address{Name: request.Name, Address: request.Email}).String()
	subject := "New WeBuildit project request — " + projectTypes[request.ProjectType]
	body := strings.Join([]string{
		"A new project request was submitted through webuild-it.com.",
		"",
		"Name: " + request.Name,
		"Email: " + request.Email,
		"Company / project: " + optionalValue(request.Company),
		"Project type: " + projectTypes[request.ProjectType],
		"Language: " + strings.ToUpper(request.Language),
		"",
		"Message:",
		normalizeBody(request.Message),
		"",
	}, "\r\n")

	return strings.Join([]string{
		"From: " + from,
		"To: " + to,
		"Reply-To: " + replyTo,
		"Subject: " + mime.QEncoding.Encode("UTF-8", subject),
		"Date: " + time.Now().UTC().Format(time.RFC1123Z),
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		"Content-Transfer-Encoding: 8bit",
		"",
		body,
	}, "\r\n")
}

func optionalValue(value string) string {
	if value == "" {
		return "Not provided"
	}
	return value
}

func normalizeBody(value string) string {
	value = strings.ReplaceAll(value, "\r\n", "\n")
	value = strings.ReplaceAll(value, "\r", "\n")
	return strings.ReplaceAll(value, "\n", "\r\n")
}

type rateEntry struct {
	count int
	reset time.Time
}

type rateLimiter struct {
	mu      sync.Mutex
	entries map[string]rateEntry
	now     func() time.Time
}

func newRateLimiter() *rateLimiter {
	return &rateLimiter{entries: make(map[string]rateEntry), now: time.Now}
}

func (r *rateLimiter) allow(key string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	now := r.now()
	entry, exists := r.entries[key]
	if !exists || !now.Before(entry.reset) {
		r.entries[key] = rateEntry{count: 1, reset: now.Add(rateLimitWindow)}
		return true
	}
	if entry.count >= rateLimitCount {
		return false
	}
	entry.count++
	r.entries[key] = entry
	return true
}

type app struct {
	cfg     config
	mailer  mailSender
	limiter *rateLimiter
}

func newHandler(cfg config, sender mailSender) http.Handler {
	application := &app{cfg: cfg, mailer: sender, limiter: newRateLimiter()}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", application.health)
	mux.HandleFunc("/api/contact", application.contact)
	mux.Handle("/", application.static())
	return securityHeaders(mux)
}

func (a *app) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (a *app) contact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method_not_allowed"})
		return
	}
	if !a.originAllowed(r) {
		writeJSON(w, http.StatusForbidden, map[string]string{"error": "origin_not_allowed"})
		return
	}
	mediaType, _, mediaTypeError := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if mediaTypeError != nil || mediaType != "application/json" {
		writeJSON(w, http.StatusUnsupportedMediaType, map[string]string{"error": "content_type_required"})
		return
	}
	if !a.cfg.contactConfigured() || a.mailer == nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "contact_unavailable"})
		return
	}
	if !a.limiter.allow(clientKey(r, a.cfg.trustProxy)) {
		w.Header().Set("Retry-After", strconv.Itoa(int(rateLimitWindow.Seconds())))
		writeJSON(w, http.StatusTooManyRequests, map[string]string{"error": "rate_limited"})
		return
	}

	payload, err := io.ReadAll(io.LimitReader(r.Body, maxRequestBytes+1))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid_request"})
		return
	}
	if len(payload) > maxRequestBytes {
		writeJSON(w, http.StatusRequestEntityTooLarge, map[string]string{"error": "request_too_large"})
		return
	}
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	var request contactRequest
	if err := decoder.Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid_request"})
		return
	}
	if err := ensureJSONEOF(decoder); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid_request"})
		return
	}
	trimRequest(&request)
	if request.Website != "" {
		writeJSON(w, http.StatusAccepted, map[string]string{"status": "accepted"})
		return
	}
	if fields := validateContact(request); len(fields) > 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{"error": "validation_failed", "fields": fields})
		return
	}

	requestID := newRequestID()
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()
	if err := a.mailer.Send(ctx, request); err != nil {
		log.Printf("contact delivery failed request_id=%s error=%v", requestID, err)
		writeJSON(w, http.StatusBadGateway, map[string]string{"error": "delivery_failed", "requestId": requestID})
		return
	}
	log.Printf("contact delivered request_id=%s project_type=%s language=%s", requestID, request.ProjectType, request.Language)
	writeJSON(w, http.StatusAccepted, map[string]string{"status": "accepted", "requestId": requestID})
}

func ensureJSONEOF(decoder *json.Decoder) error {
	var extra interface{}
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			return errors.New("multiple JSON values")
		}
		return err
	}
	return nil
}

func trimRequest(request *contactRequest) {
	request.Name = strings.TrimSpace(request.Name)
	request.Email = strings.TrimSpace(request.Email)
	request.Company = strings.TrimSpace(request.Company)
	request.ProjectType = strings.TrimSpace(request.ProjectType)
	request.Message = strings.TrimSpace(request.Message)
	request.Language = strings.ToLower(strings.TrimSpace(request.Language))
	request.Website = strings.TrimSpace(request.Website)
}

func validateContact(request contactRequest) map[string]string {
	fields := make(map[string]string)
	if length := len([]rune(request.Name)); length < 2 || length > 100 || hasHeaderControl(request.Name) {
		fields["name"] = "invalid"
	}
	if len(request.Email) > 254 || hasHeaderControl(request.Email) || !validEmail(request.Email) {
		fields["email"] = "invalid"
	}
	if len([]rune(request.Company)) > 120 || hasHeaderControl(request.Company) {
		fields["company"] = "invalid"
	}
	if _, ok := projectTypes[request.ProjectType]; !ok {
		fields["projectType"] = "invalid"
	}
	if length := len([]rune(request.Message)); length < 20 || length > 5000 || hasBodyControl(request.Message) {
		fields["message"] = "invalid"
	}
	if request.Language != "en" && request.Language != "ro" && request.Language != "ru" {
		fields["language"] = "invalid"
	}
	return fields
}

func validEmail(value string) bool {
	address, err := mail.ParseAddress(value)
	return err == nil && address.Address == value && strings.Contains(value, "@")
}

func hasHeaderControl(value string) bool {
	for _, character := range value {
		if character < 32 || character == 127 {
			return true
		}
	}
	return false
}

func hasBodyControl(value string) bool {
	for _, character := range value {
		if (character < 32 && character != '\n' && character != '\r' && character != '\t') || character == 127 {
			return true
		}
	}
	return false
}

func (a *app) originAllowed(r *http.Request) bool {
	origin := strings.TrimSuffix(strings.TrimSpace(r.Header.Get("Origin")), "/")
	if origin == "" {
		return true
	}
	if len(a.cfg.allowedOrigins) > 0 {
		_, allowed := a.cfg.allowedOrigins[origin]
		return allowed
	}
	parsed, err := url.Parse(origin)
	return err == nil && strings.EqualFold(parsed.Host, r.Host)
}

func clientKey(r *http.Request, trustProxy bool) string {
	if trustProxy {
		if forwarded := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]); net.ParseIP(forwarded) != nil {
			return forwarded
		}
		if realIP := strings.TrimSpace(r.Header.Get("X-Real-IP")); net.ParseIP(realIP) != nil {
			return realIP
		}
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && host != "" {
		return host
	}
	return r.RemoteAddr
}

func newRequestID() string {
	value := make([]byte, 8)
	if _, err := rand.Read(value); err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(value)
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("write JSON response: %v", err)
	}
}

func (a *app) static() http.Handler {
	root, err := filepath.Abs(a.cfg.staticDir)
	if err != nil {
		log.Fatalf("resolve static directory: %v", err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method_not_allowed"})
			return
		}
		path := filepath.Clean("/" + r.URL.Path)
		if path == "/" {
			path = "/index.html"
		}
		candidate := filepath.Join(root, filepath.FromSlash(strings.TrimPrefix(path, "/")))
		if candidate != root && !strings.HasPrefix(candidate, root+string(os.PathSeparator)) {
			http.NotFound(w, r)
			return
		}
		info, err := os.Stat(candidate)
		if err != nil || info.IsDir() {
			http.NotFound(w, r)
			return
		}
		if strings.HasPrefix(path, "/assets/") {
			w.Header().Set("Cache-Control", "public, max-age=3600")
		} else {
			w.Header().Set("Cache-Control", "no-cache")
		}
		if filepath.Ext(candidate) == ".webmanifest" {
			w.Header().Set("Content-Type", "application/manifest+json")
		}
		http.ServeFile(w, r, candidate)
	})
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; base-uri 'none'; connect-src 'self'; font-src 'self'; form-action 'self'; frame-ancestors 'none'; img-src 'self' data:; object-src 'none'; script-src 'self'; style-src 'self'")
		w.Header().Set("Permissions-Policy", "camera=(), geolocation=(), microphone=()")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		if r.TLS != nil || strings.EqualFold(r.Header.Get("X-Forwarded-Proto"), "https") {
			w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg := loadConfig()
	if !cfg.contactConfigured() {
		log.Printf("contact endpoint is unavailable until SMTP secret values are configured")
	}
	server := &http.Server{
		Addr:              cfg.addr,
		Handler:           newHandler(cfg, smtpSender{cfg: cfg}),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Printf("webuild-it listening on %s", cfg.addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
