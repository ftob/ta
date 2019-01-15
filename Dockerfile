FROM golang:1.10 as builder

WORKDIR /go/src/github.com/ftob/ta

COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:3.4 as goss
# Install Goss
RUN apk add --no-cache --virtual=goss-dependencies curl ca-certificates && \
    curl -fsSL https://goss.rocks/install | sh && \
    apk del goss-dependencies

FROM scratch

WORKDIR /usr/bin

ENV APP_PORT 8080
ENV APP_SERVICE_ID say_hello
ENV APP_VERSION 0.1.0
ENV APP_COMPONENT_ID http_say_hello
ENV APP_COMPONENT_TYPE backend

COPY --from=builder /go/src/github.com/ftob/ta/app .
COPY --from=goss /usr/local/bin/goss /usr/local/bin/goss

ADD .docker/healthcheck/goss/goss.yaml /goss/goss.yaml

HEALTHCHECK --interval=1s --timeout=6s CMD goss -g /goss/goss.yaml validate

CMD ["./app"]
