FROM golang:1.18

ADD . /items

WORKDIR /items
RUN go mod init items
RUN go mod tidy
RUN go build -o items .
RUN chmod +x /items

ENTRYPOINT ["/items/items"]