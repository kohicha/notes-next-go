FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY . .

# Download and install the dependencies

RUN go get -d -v ./...

RUN go build -o api .

EXPOSE 8000

CMD ["./api"]