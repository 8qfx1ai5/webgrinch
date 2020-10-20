

# 1st image is going to be our build image
FROM golang:1.13.5-alpine3.11 AS utest

RUN apk add git --no-cache
RUN apk add make --no-cache
RUN apk add build-base --no-cache
RUN apk add libxslt --no-cache

ENV GO111MODULE on
ENV WORKDIR /workdir/
WORKDIR ${WORKDIR}

COPY go.mod ${WORKDIR}
COPY go.sum ${WORKDIR}
RUN go mod download

COPY . ${WORKDIR}

ENTRYPOINT ["/bin/sh", "-c", "go test -v -count=1 ./internal/..."]
# go test -v -count=1 ./internal/... | sed ''/PASS/s//`printf "\033[32mPASS\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/''

