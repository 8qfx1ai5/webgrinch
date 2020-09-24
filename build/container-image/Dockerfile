

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
RUN cd ${WORKDIR}; go install ./cmd/webgrinchserver

RUN make utestd


# 2nd image is going to be our application image
FROM alpine:3.11 as app

RUN apk add libxslt --no-cache

RUN mkdir -p webgrinch
COPY web webgrinch
COPY api/API_webgrinch_v1.json webgrinch/static
RUN chmod -R +x webgrinch
COPY build/run.sh /
RUN chmod +x /run.sh


COPY --from=compile /go/bin/webgrinchserver /webgrinchserver
RUN chmod +x /webgrinchserver

ENTRYPOINT ["/bin/sh", "-c", "/run.sh >> /server.log 2>&1"]
