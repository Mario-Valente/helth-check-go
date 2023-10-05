FROM golang:1.21-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

# Compilar e executar o código
RUN go build -o main .
CMD ["./main"]
