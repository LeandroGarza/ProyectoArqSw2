FROM golang:1.18

ADD . /search

WORKDIR /search
RUN go mod init search
RUN go mod tidy
RUN go build -o search .
RUN chmod +x /search

ENTRYPOINT ["/search/search"]