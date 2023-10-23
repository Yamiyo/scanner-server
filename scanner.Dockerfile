FROM golang:1.20 AS go-builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go build -o scanner-server /app/cmd/scanner-server/

FROM alpine:3.12
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates && \
    apk add git
WORKDIR /scanner-server
COPY --from=go-builder /app/scanner-server /scanner-server/scanner-server
COPY --from=go-builder /app/conf.d /scanner-server/conf.d

ENTRYPOINT ["/scanner-server/scanner-server"]