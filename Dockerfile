FROM --platform=linux/amd64 golang:1.24-alpine AS builder

WORKDIR /app-source

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /assignment ./cmd/assignment

FROM --platform=linux/amd64 alpine:latest

COPY --from=builder /assignment ./assignment

ENV DB_FILENAME="gorm.db"
ENV APP_PORT="8080"
EXPOSE 8080

ENTRYPOINT ["./assignment"]