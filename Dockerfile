FROM golang:alpine AS builder

ENV GOPATH="/go:/app"
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go build -o gigate src/gigate.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/gigate .
CMD ["./gigate"]