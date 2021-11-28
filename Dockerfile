FROM ruby:3.0.3-slim

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
         build-essential \
         git \
         curl \
         imagemagick \
         libjemalloc-dev \
         libjemalloc2 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN curl -L https://go.dev/dl/go1.17.3.linux-amd64.tar.gz | tar -xzf - -C /usr/local

ENV LANG="C.UTF-8" \
    PATH="${PATH}:/usr/local/go/bin" \
    LD_PRELOAD="/usr/lib/x86_64-linux-gnu/libjemalloc.so.2"

WORKDIR /usr/src/gem
