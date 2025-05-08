FROM golang:1.24.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./backend/cmd

RUN echo "Arquivos em /app:" && ls -la /app

EXPOSE 8080

CMD ["./main"]