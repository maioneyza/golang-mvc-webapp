FROM golang:latest

WORKDIR /go/src/app

COPY ./ /go/src/app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 80
ENTRYPOINT CompileDaemon --build="go build commands/runserver.go" --command=./runserver