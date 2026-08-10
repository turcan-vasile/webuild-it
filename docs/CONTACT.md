# Contact delivery contract

The public form posts JSON to `POST /api/contact`. The application validates the
request and delivers a plain-text email through the configured SMTP relay. Browser
code never receives SMTP credentials, the sender identity or the delivery address.

## Runtime configuration

The deployment provides these non-secret values directly:

- `SMTP_HOST`
- `SMTP_PORT`
- `SMTP_FROM_NAME`
- `CONTACT_ALLOWED_ORIGINS`
- `TRUST_PROXY_HEADERS`

The namespace-local Kubernetes Secret named `webuild-it-contact` must provide:

- `SMTP_USER`
- `SMTP_PASSWORD`
- `SMTP_FROM_ADDRESS`
- `CONTACT_RECIPIENT`

Do not commit Secret manifests, encoded values, copied cluster output or local
environment files. The production Secret is an owner-approved runtime operation
and is intentionally outside this repository change.

If the Secret is absent or incomplete, the site still serves normally and the
contact endpoint returns `503 contact_unavailable`. The UI explains that the
message was not sent, so submissions cannot fail silently.

## Request controls

The endpoint applies:

- same-origin checks;
- a 32 KiB request limit;
- strict JSON decoding and a fixed project-type allowlist;
- server-side field trimming, email parsing and length limits;
- CR/LF rejection for values used in mail headers;
- a hidden honeypot field;
- an in-memory limit of five requests per client address per ten minutes;
- a 15-second delivery timeout;
- generic client errors and request-ID-only delivery logs.

The in-memory limiter is intentionally small and dependency-free. With two
replicas, limits are per pod rather than globally coordinated. If abuse becomes
material, move rate limiting to the ingress or a shared store in a separate,
explicitly governed infrastructure task.

## Safe validation

Local and CI validation does not require real credentials:

```bash
make verify
make docker-build
```

After the owner provisions the namespace Secret and authorizes deployment, send
one clearly marked test request through the public form and confirm receipt in the
approved delivery mailbox. Record only timestamp, request ID and receipt status;
do not copy message content or mail headers into public evidence.
