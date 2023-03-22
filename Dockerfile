FROM golang:1-buster

RUN go install github.com/beego/bee/v2@latest


ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/mathapp
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

RUN go mod init mathapp
RUN go mod tidy





EXPOSE 8081
CMD ["bee", "run"]