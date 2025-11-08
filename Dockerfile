# Stage 1: Builder
FROM golang:1.25-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -ldflags "-s -w" -v -o ./bin/resttest main.go

# Stage 2: Final
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /usr/src/app/bin/resttest .

CMD ["./resttest"]
