FROM golang:1.17.2

WORKDIR /service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o search-service .

EXPOSE 2021

CMD ["./search-service"]
