FROM golang:1.17.1

WORKDIR /app

COPY . .
RUN go mod tidy


CMD ["go", "run", "main.go"]