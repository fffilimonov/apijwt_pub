FROM ubuntu:16.04
MAINTAINER fffilimonov <fffilimonov@yandex.ru>

ENV DEBIAN_FRONTEND noninteractive
ENV DEBCONF_NONINTERACTIVE_SEEN true
ENV TZ "US/Pacific"
ENV LANGUAGE en_US.UTF-8
ENV LANG en_US.UTF-8

USER root

RUN  echo "deb http://archive.ubuntu.com/ubuntu xenial main universe\n" > /etc/apt/sources.list \
  && echo "deb http://archive.ubuntu.com/ubuntu xenial-updates main universe\n" >> /etc/apt/sources.list \
  && echo "deb http://security.ubuntu.com/ubuntu xenial-security main universe\n" >> /etc/apt/sources.list

#locale
RUN locale-gen en_US.UTF-8 \
  && dpkg-reconfigure --frontend noninteractive locales \
  && apt-get update -qqy \
  && apt-get -qqy --no-install-recommends install \
    language-pack-en \
  && rm -rf /var/lib/apt/lists/* /var/cache/apt/*

#Timezone
RUN echo "${TZ}" > /etc/timezone \
  && dpkg-reconfigure --frontend noninteractive tzdata

#Go
RUN apt-get update -qqy \
  && apt-get -qqy --no-install-recommends install \
    gcc \
    build-essential \
    golang \
    git \
    ssh \
    openssl \
    ca-certificates \
    vim \
  && rm -rf /var/lib/apt/lists/* /var/cache/apt/*

RUN mkdir -p /root/.hyper
COPY config.json /root/.hyper
RUN mkdir -p /root/.ssh
COPY id_rsa /root/.ssh
RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts
RUN chmod 400 ~/.ssh/id_rsa

ARG CACHE_DATE=not_a_date

RUN git clone git@github.com:fffilimonov/apijwt.git
WORKDIR 'apijwt'
RUN ./build.sh

COPY entry_point.sh /opt/bin/entry_point.sh
RUN chmod +x /opt/bin/entry_point.sh

EXPOSE 5000

CMD ["/opt/bin/entry_point.sh"]
