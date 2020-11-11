FROM golang:1.14.4-alpine3.12 as builder

RUN apk update && apk upgrade && \
  apk --update add git make

WORKDIR /app

COPY . .

RUN make build

FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --update --no-cache add tzdata && \
  mkdir /app && mkdir /log

WORKDIR /app 

EXPOSE 4000

COPY --from=builder /app/api-creator-admin /app

CMD /app/api-creator-admin

# FROM mysql:5.7

# ADD ./docker/my.cnf /etc/mysql/my.cnf

# RUN chmod 644 /etc/mysql/my.cnf