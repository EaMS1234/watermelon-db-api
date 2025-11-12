FROM golang:alpine

WORKDIR /api/

COPY . . 
RUN go mod tidy
RUN go build .

EXPOSE 8080

CMD ["./api"]

