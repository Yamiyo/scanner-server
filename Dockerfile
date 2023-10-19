FROM golang:1.18.7 AS go-builder
ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor
WORKDIR /app
COPY . .
RUN go build -o portto-homework /app/cmd/

FROM alpine:3.12
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates && \
    apk add git
WORKDIR /portto-homework
COPY --from=go-builder /app/portto-homework /portto-homework/portto-homework
COPY --from=go-builder /app/conf.d /portto-homework/conf.d
COPY --from=go-builder /app/image /portto-homework/image
ENTRYPOINT ["/portto-homework/portto-homework"]
