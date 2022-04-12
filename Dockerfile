# syntax=docker/dockerfile:1

##
## Base
##

FROM golang:1.18-alpine as base

WORKDIR /app

ENV GO111MODULE="on"
ENV CGO_ENABLED=0

COPY go.mod ./
COPY go.sum ./

RUN go mod download \
    && go mod verify

COPY cmd ./cmd
COPY internal ./internal
COPY kit ./kit

##
## Build
##
FROM base as builder

RUN go build -o /my-game -a ./cmd/api/main.go

CMD [ "/my-game" ]

##
## Debug
##
FROM base as debug

RUN go install github.com/go-delve/delve/cmd/dlv@latest \
    && go build -gcflags="all=-N -l" -o /my-game -a ./cmd/api/main.go

CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/my-game", "--", "--config", "/myconfig.cfg"]

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /my-game /my-game

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/my-game" ]
