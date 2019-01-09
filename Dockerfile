FROM golang:latest as builder

WORKDIR /go/src/gitgub.com/ftob/ta

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM scratch

WORKDIR /usr/bin

COPY --from=builder /go/src/gitgub.com/ftob/ta/app .

ENV APP_PORT 8080
ENV APP_SERVICE_ID say_hello
ENV APP_VERSION 0.1.0
ENV APP_COMPONENT_ID http_say_hello
ENV APP_COMPONENT_TYPE backend

CMD ["/usr/bin/app"]