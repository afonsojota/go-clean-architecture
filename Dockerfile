FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Instalar bash (necessário para o wait-for-it.sh)
RUN apk add --no-cache bash

# Copiar e configurar wait-for-it.sh
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Build da aplicação
RUN go build -o go-clean-architecture cmd/main.go

EXPOSE 8080
EXPOSE 8000
EXPOSE 50051

# Comando para aguardar o banco de dados e então iniciar a aplicação
CMD ["/wait-for-it.sh", "db:5432", "--", "./go-clean-architecture"]
