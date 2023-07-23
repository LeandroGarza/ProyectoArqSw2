FROM golang:1.18

ADD . /cache

WORKDIR /cache
RUN go mod init cache
RUN go mod tidy
RUN go build -o cache .
RUN chmod +x /cache

ENTRYPOINT ["/cache/cache"]