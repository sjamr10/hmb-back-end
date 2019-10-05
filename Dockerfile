FROM golang:1.13
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build" -command="./hmb-back-end"
