FROM alpine:latest
MAINTAINER Roland Kammerer <dev.rck@gmail.com>

# docker run -it --rm simplepush -k key -m message

ENV SIMPLEPUSH_VERSION 0.2

RUN apk add --no-cache --virtual .build-deps wget ca-certificates
RUN wget "https://github.com/rck/simplepush/releases/download/v${SIMPLEPUSH_VERSION}/simplepush-alpine-amd64" -O /usr/local/bin/simplepush
RUN chmod +x /usr/local/bin/simplepush
RUN apk del .build-deps

ENTRYPOINT ["simplepush"]
