from golang:latest

ADD . /go/src
WORKDIR /go/src

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go