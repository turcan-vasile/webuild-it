.DEFAULT_GOAL := help

.PHONY: help
help:
	@printf "make serve        # run the site and contact API locally\n"
	@printf "make test         # run Go tests\n"
	@printf "make fmt-check    # verify Go formatting\n"
	@printf "make verify       # run the local quality gate\n"
	@printf "make docker-build # build the production image\n"

.PHONY: serve
serve:
	go run ./server

.PHONY: test
test:
	go test ./...

.PHONY: fmt-check
fmt-check:
	@test -z "$$(gofmt -l server)" || { printf "Go files need formatting:\n"; gofmt -l server; exit 1; }

.PHONY: verify
verify: fmt-check test

.PHONY: docker-build
docker-build:
	docker build -t webuild-it:local .
