FROM golang:1.15-alpine AS builder
WORKDIR /go/src/github.com/johnbellone/persona-service
COPY . .
RUN apk add -U --no-cache ca-certificates
RUN go mod tidy
RUN go build -o /go/bin/persona-service

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/persona-service /bin/persona-service
WORKDIR /root
EXPOSE 8080
ENTRYPOINT ["/bin/persona-service"]
