FROM golang:alpine3.13

WORKDIR /go/src/app
COPY . .

RUN go mod tidy
RUN go build -o ./essentiel-dns .

EXPOSE 53 53/udp

CMD ["/go/src/app/essentiel-dns"]
