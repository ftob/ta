FROM golang:latest as builder

WORKDIR /go/src/github.com/ftob/ta

COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

WORKDIR /usr/bin

ENV APP_PORT 8080
ENV APP_SERVICE_ID say_hello
ENV APP_VERSION 0.1.0
ENV APP_COMPONENT_ID http_say_hello
ENV APP_COMPONENT_TYPE backend

COPY --from=builder /go/src/github.com/ftob/ta/app .


CMD ["./app"]