FROM alpine:3.11.6

MAINTAINER https://github.com/StevenPG

ENV MOUNTEBANK_VERSION=2.2.1
EXPOSE 2525

# Node Imposter
EXPOSE 8080

# Pod Imposter
EXPOSE 8081

RUN apk update \
    && apk add --no-cache nodejs npm

RUN npm install -g mountebank@${MOUNTEBANK_VERSION} --production

RUN npm cache clean --force 2>/dev/null \
    && rm -rf /tmp/npm*

COPY templates templates

CMD ["mb", "--configfile", "templates/imposters.ejs"]
