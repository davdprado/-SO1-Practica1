FROM golang:1.17.1-alpine3.14

RUN mkdir -p /backend/go

WORKDIR /backend/go

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]