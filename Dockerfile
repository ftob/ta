FROM golang:latest as builder

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM scratch

WORKDIR /usr/bin

COPY --from=builder /go/bin/app /usr/bin


CMD ["/usr/bin/app"]