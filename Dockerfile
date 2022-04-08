# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-alpine as builder

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

RUN go build -o /my-game -a ./cmd/api/main.go

CMD [ "/my-game" ]

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /my-game /my-game

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/my-game" ]
