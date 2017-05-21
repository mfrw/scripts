FROM ubuntu:latest
LABEL MAINTAINER mfrw
RUN apt-get update &&  apt-get install -y  build-essential \
        make\
        wget \
        curl \
        zip \
        vim \
        zsh \
	git \
	gcc \
	g++ \
	ctags \
	cscope \
	autoconf \
	libtool \
	libxml2-dev
RUN git clone https://github.com/vivkumar/cse502.git
WORKDIR /cse502/hclib/
RUN ./install.sh 2>&1 | tail -1 >> ~/.bashrc

ENTRYPOINT /bin/bash

