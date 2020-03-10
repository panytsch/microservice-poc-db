FROM golang
WORKDIR $APP_CONTAINER_PATH
EXPOSE 8080 2345
RUN go get github.com/derekparker/delve/cmd/dlv
ADD $APP_HOST_PATH $APP_CONTAINER_PATH

CMD ["dlv", "debug", "github.com/olivere/go-container-debugging", "--headless", "--listen=:2345", "--api-version=2", "--log"]