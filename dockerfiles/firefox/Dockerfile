# #!/bin/bash
# docker run -it --rm \
#         -v /tmp/.X11-unix:/tmp/.X11-unix \
#         -e DISPLAY=unix$DISPLAY  \
#         -v /run/user/1000/pulse:/run/user/1000/pulse \
#         -v /dev/shm:/dev/shm \
#         -v /etc/machine-id:/etc/machine-id \
#         -v ~/.pulse:/home/pact/.pulse \
#         mfrw/firefox firefox

FROM ubuntu:latest
LABEL MAINTAINER mfrw

ENV UNAME pacat

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive apt-get install --yes pulseaudio-utils firefox

# Set up the user
RUN export UNAME=$UNAME UID=1000 GID=1000 && \
    mkdir -p "/home/${UNAME}" && \
    echo "${UNAME}:x:${UID}:${GID}:${UNAME} User,,,:/home/${UNAME}:/bin/bash" >> /etc/passwd && \
    echo "${UNAME}:x:${UID}:" >> /etc/group && \
    mkdir -p /etc/sudoers.d && \
    echo "${UNAME} ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/${UNAME} && \
    chmod 0440 /etc/sudoers.d/${UNAME} && \
    chown ${UID}:${GID} -R /home/${UNAME} && \
    gpasswd -a ${UNAME} audio

COPY pulse-client.conf /etc/pulse/client.conf

USER $UNAME
ENV HOME /home/pacat

#ENTRYPOINT /usr/bin/firefox

