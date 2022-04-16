# syntax=docker/dockerfile:1

##
## Base
##

FROM golang:1.18-alpine as base

WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

##
## Dev
##
FROM base as dev

RUN go install github.com/cosmtrek/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080
EXPOSE 2345

ENTRYPOINT ["air"]

##
## Build
##

FROM base as build

COPY go.mod ./
COPY go.sum ./

RUN go mod download \
    && go mod verify

COPY cmd ./cmd
COPY internal ./internal
COPY kit ./kit

RUN go build -o /my-game -a ./cmd/api/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /my-game /my-game

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/my-game" ]
