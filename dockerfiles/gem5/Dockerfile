FROM ubuntu:latest
MAINTAINER Muhammad Falak R Wani <falakreyaz@gmail.com>

# get dependencies
RUN apt-get update
RUN apt-get install -y build-essential git-core m4 scons nano zlib1g zlib1g-dev libprotobuf-dev protobuf-compiler libprotoc-dev libgoogle-perftools-dev swig python-dev python
RUN apt-get clean

# checkout repo with mercurial
RUN useradd -u 1000 -ms /bin/bash gem5
WORKDIR /home/gem5
RUN git clone --depth 1 https://github.com/gem5/gem5.git
# build it
WORKDIR /home/gem5/gem5
USER gem5
RUN scons -j$(nproc) --ignore-style build/X86/gem5.opt
RUN mv build/X86/gem5.out /usr/local/bin

ENTRYPOINT bash
