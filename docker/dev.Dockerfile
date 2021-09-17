FROM golang:1.17

# Meta data:
LABEL maintainer="email@mattglei.ch"
LABEL description="🦎 A slack bot for random gex quotes"

# Install air:
RUN go install github.com/cosmtrek/air@latest

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

CMD ["air"]
