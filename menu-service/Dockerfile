FROM golang:1.17.2

WORKDIR /service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o menu-service .

EXPOSE 2011

CMD ["./menu-service"]
