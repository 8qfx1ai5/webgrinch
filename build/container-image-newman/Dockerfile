

# image for testing the running apis
FROM alpine:3.11 AS testing

# install newman
RUN apk add --update npm --no-cache
RUN npm install -g newman

COPY api/Contract_Tests.postman_collection.json /

ENTRYPOINT ["newman"]
