FROM golang
WORKDIR /go/src/app
EXPOSE 8080 2345
RUN go get github.com/derekparker/delve/cmd/dlv
ADD . /go/src/app

CMD ["dlv", "debug", "github.com/olivere/go-container-debugging", "--headless", "--listen=:2345", "--api-version=2", "--log"]