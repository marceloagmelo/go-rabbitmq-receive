FROM golang:1.13.6 AS builder

ENV APP_HOME /go/src/github.com/marceloagmelo/go-rabbitmq-receive

ADD . $APP_HOME

WORKDIR $APP_HOME

 RUN go get ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o go-rabbitmq-receive && \
    rm -Rf /tmp/* && rm -Rf /var/tmp/*

###
# IMAGEM FINAL
###
FROM alpine:3.11

ENV GID 23550
ENV UID 23550
ENV USER golang

ENV APP_BUILDER /go/src/github.com/marceloagmelo/go-rabbitmq-receive/
ENV APP_HOME /opt/app

WORKDIR $APP_HOME

COPY --from=builder $APP_BUILDER/go-rabbitmq-receive $APP_HOME/go-rabbitmq-receive

RUN apk add --no-cache ca-certificates && \
    addgroup -g $GID -S $USER && \
    adduser -u $UID -S -h "$(pwd)" -G $USER $USER && \
    chown -fR $USER:0 $APP_HOME


ENV PATH $APP_HOME:$PATH

EXPOSE 8080

USER ${USER}

CMD [ "go-rabbitmq-receive" ]
