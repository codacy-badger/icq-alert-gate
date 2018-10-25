FROM golang:alpine AS builder

ENV GOPATH="/go:/app"
WORKDIR /app
COPY . .
RUN apk add git
RUN go get -d -v ./...
RUN go build -o gigate src/gigate.go

FROM alpine:latest
COPY --from=builder /app/gigate .
CMD ["./gigate"]