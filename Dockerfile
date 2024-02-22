# syntax = docker/dockerfile:1.0-experimental

FROM golang:1.22-alpine as BUILD

ENV GOOS "linux"
ENV GOVCS "*:all"
ENV CGO_ENABLED "0"
ENV GOPRIVATE "gitlab.com/ethrgg/*"

WORKDIR /go/src/gitlab.com/ethrgg/templates/go-cobra-cli

COPY . ./

# RUN mv .netrc ~/.netrc

RUN apk add --no-cache git

RUN go mod download && go build -ldflags "-s -w" -o /go-cobra-cli

FROM alpine:3.19

RUN apk add --no-cache bash git openssh-client

COPY --from=build /go-cobra-cli /usr/bin/go-cobra-cli

RUN git config --global --add safe.directory '*'

ENTRYPOINT ["go-cobra-cli"]
