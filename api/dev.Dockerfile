FROM golang:1.17.1

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go get -u github.com/cosmtrek/air


CMD ["air", "-c", ".air.toml"]