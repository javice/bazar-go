# Build stage (Backend Go)
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /bazar-api ./cmd/main.go

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /bazar-api /app/bazar-api
COPY --from=builder /app/docs /app/docs          
COPY frontend/ /app/frontend                    

EXPOSE 8080
CMD ["/app/bazar-api"]