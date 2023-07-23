FROM golang:1.18

ADD . /consumers

WORKDIR /consumers
RUN go mod init consumers
RUN go mod tidy
RUN go build -o consumers .
RUN chmod +x /consumers

ENTRYPOINT ["/consumers/consumers"]