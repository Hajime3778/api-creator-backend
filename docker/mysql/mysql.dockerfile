FROM mysql:5.7

ADD ./my.cnf /etc/mysql/my.cnf

RUN chmod 644 /etc/mysql/my.cnf

RUN apt-get update

RUN apt-get install -y tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime