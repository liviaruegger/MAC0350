FROM golang:1.24.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY backend/utils/wait-for-it.sh /backend/utils/wait-for-it.sh
RUN chmod +x /backend/utils/wait-for-it.sh

COPY . .
RUN go build -o main ./backend/cmd

EXPOSE 8080

CMD ["./main"]