FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/cloudevents/sdk-go

RUN go get github.com/robfig/cron

WORKDIR /go/src/github.com/oms-services/cron

ADD . /go/src/github.com/oms-services/cron

RUN go install github.com/oms-services/cron

ENTRYPOINT cron

EXPOSE 3000