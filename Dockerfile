FROM golang:alpine

WORKDIR /api/

ENV REQUEST_HOST=http://localhost:5000

COPY . . 
RUN go mod tidy
RUN go build .

EXPOSE 8080

CMD ["./api"]

