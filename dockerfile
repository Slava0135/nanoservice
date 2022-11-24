# syntax=docker/dockerfile:1

FROM golang:1.19 AS builder
COPY . /go/src/github.com/slava0135/nanoservice
WORKDIR /go/src/github.com/slava0135/nanoservice
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 go build -o app .

FROM alpine:3.17.0
COPY --from=builder /go/src/github.com/slava0135/nanoservice/app /
EXPOSE 8080
CMD ["/app"]