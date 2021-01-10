FROM golang:1.14.6

WORKDIR /code
RUN go get honnef.co/go/tools/cmd/staticcheck
RUN go get -u golang.org/x/lint/golint

PORT 8080