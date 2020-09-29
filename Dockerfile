FROM golang:1.15 AS builder
WORKDIR /go/src/github.com/johnbellone/persona-service
RUN go mod tidy
RUN go build -o /go/bin/persona-service

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/persona-service /go/bin/persona-service
CMD ["/go/bin/persona-service"]
