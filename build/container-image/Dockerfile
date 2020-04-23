

# 1st image is going to be our build image
FROM golang:1.13.5-alpine3.10 AS compile

RUN apk add git --no-cache

ENV GO111MODULE on
ENV WORKDIR /workdir/
WORKDIR ${WORKDIR}

# COPY go.mod go.sum ${WORKDIR}
COPY go.mod ${WORKDIR}
RUN go mod download

COPY . ${WORKDIR}
RUN go install ./cmd/vcryptserver


# 2nd image is going to be our application image
FROM alpine:3.8 as app

COPY --from=compile /go/bin/vcryptserver /vcryptserver
RUN chmod +x /vcryptserver

ENTRYPOINT ["/vcryptserver", "-p", "8080"]

