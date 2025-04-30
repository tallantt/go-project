# Используем официальный golang для сборки
FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Статическая сборка бинарника (без CGO)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Минимальный образ для запуска
FROM alpine:latest

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /app/main .

# Открываем порт 8080 (или тот, что у вас в приложении)
EXPOSE 8080

CMD ["./main"]
