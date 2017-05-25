FROM ubuntu:14.04
MAINTAINER mfrw
RUN apt-get update &&  apt-get install -y  build-essential \
        make\
        wget \
        curl \
        zip \
        vim \
        zsh
RUN apt-get update && apt-get install -y python python-pip gnuplot git
RUN pip install django==1.2.3
RUN git clone https://github.com/jamesbornholt/plotty.git
WORKDIR /plotty
RUN mkdir cache cache/csv cache/graph cache/log
RUN chmod -R a+rw cache
RUN mkdir log
RUN python manage.py syncdb
RUN python install_defaults.py

ENTRYPOINT bash
