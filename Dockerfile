FROM golang:alpine

WORKDIR /go/src/vanamehe_bot
COPY . .

RUN go get -d -v ./...
RUN apk add --update make
RUN apk add --update gcc
RUN apk add --update opus
RUN apk add --update ffmpeg
RUN make build

CMD ["bin/main"]