# Stage 1: Builder
FROM golang:1.25-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -ldflags "-s -w" -v -o ./bin/resttest ./cmd/resttest/main.go

# Stage 2: Final
FROM scratch

# Add metadata as key-value pairs into the image.
LABEL author="Salman Shaik"

WORKDIR /root/

COPY --from=builder /usr/src/app/bin/resttest .

#expose port from the container
EXPOSE 5432

# CMD sets the default command to be used when starting a container.
# This can be overruled with docker run.
# CMD ["./resttest"]
#
#
# ENTRYPOINT defines the 'main' command to be used when starting a container.
ENTRYPOINT [ "./resttest" ]
