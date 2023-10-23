FROM golang:1.20 AS go-builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go build -o api-server /app/cmd/api-server

FROM alpine:3.12
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates && \
    apk add git
WORKDIR /api-server
COPY --from=go-builder /app/api-server /api-server/api-server
COPY --from=go-builder /app/conf.d /api-server/conf.d

ENTRYPOINT ["/api-server/api-server"]
