FROM ubuntu:14.04
MAINTAINER mfrw
RUN apt-get update &&  apt-get install -y  build-essential \
        make\
        wget \
        curl \
        zip \
        vim \
        zsh
RUN apt-get install -y openjdk-7-jdk
RUN apt-get install -y git ant gcc g++ gcc-multilib g++-multilib bison automake gettext libtool
RUN useradd -u 1000 -U mfrw
RUN useradd -u 1001 -U vivek
RUN useradd -u 1002 -U alind

ENTRYPOINT /bin/bash

