# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod & go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary ชื่อ main
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Stage 2: Run
FROM gcr.io/distroless/static

WORKDIR /

COPY --from=builder /app/main /server

EXPOSE 3004

ENTRYPOINT ["/server"]
