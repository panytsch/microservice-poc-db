FROM golang:latest

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

#ENV GIT_CURL_VERBOSE 0

RUN git config --global github.user panytsch
RUN git config --global github.token 852a5ffd134368f45b946959579fa6d8d58940ee

RUN git config --global http.sslVerify false
RUN git config --global http.proxyAuthMethod basic

# build directories
RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app

WORKDIR /go/src/app

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Build my app
#RUN go build -o /app/main .
CMD ["/app/main"]