FROM golang:1.17-alpine
ADD . /go/src/github.com/usr/ws-d
WORKDIR /go/src/github.com/usr/ws-d
RUN go get github.com/lib/pq
RUN go install github.com/usr/ws-d
ENTRYPOINT /go/bin/ws-d
EXPOSE 8080