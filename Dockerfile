FROM golang:1.10 as builder

WORKDIR /go/src/github.com/ftob/ta

COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


ADD https://github.com/chrisaxiom/docker-health-check/archive/v0.3.tar.gz /
RUN tar -xvzf /v0.3.tar.gz -C /

FROM scratch

WORKDIR /usr/bin

ENV APP_PORT 8080
ENV APP_SERVICE_ID say_hello
ENV APP_VERSION 0.1.0
ENV APP_COMPONENT_ID http_say_hello
ENV APP_COMPONENT_TYPE backend

COPY --from=builder /docker-health-check-0.3/docker-health-check /docker-health-check
COPY --from=builder /go/src/github.com/ftob/ta/app .

HEALTHCHECK --interval=8s --timeout=120s --retries=8 CMD ["/docker-health-check", "-url=http://localhost:8080/service/v1/health"]

CMD ["./app"]
