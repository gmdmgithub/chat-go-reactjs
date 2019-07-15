FROM golang:1.11.2-stretch
# FROM golang:1.12.6-alpine3.9 
ARG GO111MODULE=on
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# EXPOSE 8081
ENTRYPOINT ["/app/chat-go-reactjs"]