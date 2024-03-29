FROM golangci/golangci-lint:v1.42

# Meta data
LABEL maintainer="email@mattglei.ch"
LABEL description="🦎 A slack bot for random gex quotes"

# Copying over files
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing hadolint:
WORKDIR /usr/bin
RUN curl -sL -o hadolint "https://github.com/hadolint/hadolint/releases/download/v1.17.6/hadolint-$(uname -s)-$(uname -m)" \
    && chmod 700 hadolint

# Installing goreleaser
WORKDIR /
RUN git clone https://github.com/goreleaser/goreleaser
WORKDIR /goreleaser
RUN go get ./... \
    && go build -o goreleaser . \
    && mv goreleaser /usr/bin

# Installing make
RUN apt-get update && apt-get install make=4.3-4.1 -y --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

CMD ["make", "local-lint"]
