FROM golang:latest

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

RUN sudo apt-get update
    sudo apt-get install unzip \
    curl && \
    PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP && \
    sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

# build directories
RUN mkdir /app && \
    mkdir /go/src/app
ADD . /go/src/app

WORKDIR /go/src/app

# Go dep!
RUN go get -u github.com/golang/dep/... && \
    dep ensure

# Build my app
RUN go build -o /app/main .
CMD ["/app/main"]