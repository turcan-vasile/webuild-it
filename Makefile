.DEFAULT_GOAL := help

.PHONY: help
help:
	@printf "make serve        # local static preview\n"
	@printf "make docker-build # build nginx image\n"

.PHONY: serve
serve:
	python3 -m http.server 8088 --directory site

.PHONY: docker-build
docker-build:
	docker build -t webuild-it:local .
