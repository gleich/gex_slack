FROM golang:1.17

# Meta data:
LABEL maintainer="email@mattglei.ch"
LABEL description="🦎 A slack bot for random gex quotes"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

CMD ["make", "local-test"]
