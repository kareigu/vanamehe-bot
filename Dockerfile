FROM golang:alpine

WORKDIR /go/src/vanamehe_bot
COPY . .

RUN go get -d -v ./...
RUN apk add --update make
RUN apk add --update g++
RUN apk add --update opus opus-dev
RUN apk add --update ffmpeg
RUN apk add --update build-base
RUN apk add --update pkgconfig
RUN make build

CMD ["bin/main"]
