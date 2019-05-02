FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/cloudevents/sdk-go

RUN go get github.com/robfig/cron

WORKDIR /go/src/github.com/heaptracetechnology/microservice-cron

ADD . /go/src/github.com/heaptracetechnology/microservice-cron

RUN go install github.com/heaptracetechnology/microservice-cron

ENTRYPOINT microservice-cron

EXPOSE 3000