FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o fibonacci-api cmd/api/main.go

CMD ["./fibonacci-api"]
