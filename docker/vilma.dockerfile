# syntax=docker/dockerfile:experimental
FROM golang:1.20-alpine as builder

RUN apk add --no-cache git openssh-client go-task

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

COPY . /app
RUN go-task build-vilma-prod
RUN ls -lah build && chmod +x build/vilma && pwd

FROM alpine:3.17

LABEL maintainer="Phoops info@phoops.it"
LABEL environment="production"
LABEL project="muv-platform"
LABEL service="vilma"

RUN apk update && apk add --no-cache tzdata

WORKDIR /app
COPY --from=builder /app/build/vilma /app

CMD ["./vilma"]
