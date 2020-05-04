

# 1st image is going to be our build image
FROM golang:1.13.5-alpine3.11 AS compile

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
RUN cd ${WORKDIR}; go install ./cmd/vcryptserver

RUN make utestd


# 2nd image is going to be our application image
FROM alpine:3.11 as app

RUN apk add libxslt --no-cache

RUN mkdir -p viewcrypt
COPY web viewcrypt
COPY api/ViewCrypt_API_v1.json viewcrypt/static

RUN chmod -R +x viewcrypt

COPY --from=compile /go/bin/vcryptserver /vcryptserver
RUN chmod +x /vcryptserver

ENTRYPOINT ["/bin/sh", "-c", "/vcryptserver -p 80 > server.log 2>&1"]
