FROM golang:alpine

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build
CMD ["./simple_server"]