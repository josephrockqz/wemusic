FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o wemusic-golang cmd/main.go

EXPOSE 8080

CMD ["/app/wemusic-golang"]
