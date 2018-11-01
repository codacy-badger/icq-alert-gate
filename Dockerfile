FROM golang:alpine AS builder

ENV GOPATH="/go:/app"
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go build -o icqgate src/icqgate.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/icqgate .
CMD ["./icqgate"]