FROM golang:1.18

ADD . /messages

WORKDIR /messages
RUN go mod init messages
RUN go mod tidy
RUN go build -o messages .
RUN chmod +x /messages

ENTRYPOINT ["/messages/messages"]