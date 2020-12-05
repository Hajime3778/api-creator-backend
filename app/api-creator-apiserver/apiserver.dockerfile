FROM golang:1.14.4-alpine3.12 as builder

RUN apk update && apk upgrade && \
  apk --update add git make

WORKDIR /app

COPY . .

RUN make build-apiserver

FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --update --no-cache add tzdata && \
  mkdir /app && mkdir /log

WORKDIR /app 

EXPOSE 9000

COPY --from=builder /app/api-creator-apiserver /app

CMD /app/api-creator-apiserver