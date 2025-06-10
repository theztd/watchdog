## Build
FROM golang:1.23-bookworm AS build


COPY . /usr/src/watchdog/

WORKDIR /usr/src/watchdog

RUN go build -o /usr/local/bin/watchdog


## Deploy
FROM debian:stable-slim

RUN useradd watchdog
COPY --from=build /usr/local/bin/watchdog /usr/local/bin/watchdog

WORKDIR /opt/watchdog

RUN chown watchdog -R /opt/watchdog

USER watchdog

ENV ADDRESS=":8080"

EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/watchdog"]
