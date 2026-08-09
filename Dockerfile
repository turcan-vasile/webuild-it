FROM golang:1.24-alpine AS build

WORKDIR /src
COPY go.mod ./
COPY server/ ./server/
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/webuild-it ./server

FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
    && addgroup -S -g 10001 webuildit \
    && adduser -S -D -H -u 10001 -G webuildit webuildit

WORKDIR /app
COPY --from=build /out/webuild-it /app/webuild-it
COPY site/ /app/site/

ENV LISTEN_ADDR=:8080 \
    STATIC_DIR=/app/site

USER 10001:10001
EXPOSE 8080

ENTRYPOINT ["/app/webuild-it"]
