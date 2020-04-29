

# 1st image is going to be our build image
FROM golang:1.13.5-alpine3.11 AS compile

RUN apk add git --no-cache
RUN apk add make --no-cache
RUN apk add build-base --no-cache
RUN apk add libxslt --no-cache

ENV GO111MODULE on
ENV WORKDIR /workdir/
WORKDIR ${WORKDIR}


# COPY go.mod go.sum ${WORKDIR}
COPY go.mod ${WORKDIR}
RUN go mod download

RUN go get github.com/rakyll/statik

COPY third_party/swagger-ui/dist/* ${WORKDIR}/third_party/swagger-ui/dist/
COPY api/ViewCrypt_API_v1.json ${WORKDIR}/third_party/swagger-ui/dist/
RUN cd ${WORKDIR}/third_party/swagger-ui; statik -src=${WORKDIR}/third_party/swagger-ui/dist

COPY . ${WORKDIR}
RUN cd ${WORKDIR}; go install ./cmd/vcryptserver

RUN make utest


# 2nd image is going to be our application image
FROM alpine:3.11 as app

RUN apk add libxslt --no-cache

COPY --from=compile /go/bin/vcryptserver /vcryptserver
RUN chmod +x /vcryptserver

ENTRYPOINT ["/bin/sh", "-c", "/vcryptserver -p 8080 > server.log 2>&1"]
