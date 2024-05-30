ARG GO_VERSION=1
FROM golang:${GO_VERSION} as base 

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o ./tmp/main ./server/main.go 

EXPOSE 6969

CMD ["./tmp/main"]
