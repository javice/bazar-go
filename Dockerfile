# Build stage (Backend Go)
FROM golang:1.24-alpine AS builder

# Instalar dependencias necesarias para CGO y SQLite
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /bazar-api ./cmd/main.go

# Runtime stage
FROM alpine:latest

# Instalar dependencias necesarias para ejecutar SQLite
RUN apk add --no-cache sqlite-libs

WORKDIR /app
COPY --from=builder /bazar-api /app/bazar-api
COPY --from=builder /app/docs /app/docs          
COPY frontend/ /app/frontend                    

EXPOSE 8080
CMD ["/app/bazar-api"]