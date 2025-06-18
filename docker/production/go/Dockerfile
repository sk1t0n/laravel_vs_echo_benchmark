FROM golang:1.24 AS builder

WORKDIR /app

COPY ./echo/go.mod ./echo/go.sum ./
RUN go mod download

COPY ./echo/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /web_app /app/cmd/app/main.go

RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

FROM alpine:latest

WORKDIR /

COPY --from=builder /web_app /web_app

COPY ./docker/production/go/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
COPY --from=builder /usr/local/bin/goose /usr/local/bin/goose
RUN apk add --update postgresql-client \
    && rm -rf /var/cache/apk/*
ADD ./echo/migrations /migrations

ENV GOOSE_DRIVER=postgres
ENV GOOSE_DBSTRING=postgres://postgres:password@pgbouncer:6432/postgres
ENV GOOSE_MIGRATION_DIR=/migrations
ENV GOOSE_TABLE=goose_migrations

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

EXPOSE 9001
CMD [ "/web_app" ]
