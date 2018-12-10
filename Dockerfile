FROM alpine

MAINTAINER Devin

ENV VIRTUAL_HOST api.connectus.today

ARG PORT
ENV PORT ${PORT}

EXPOSE ${PORT}

COPY ./bin /usr/src/server
WORKDIR /usr/src/server

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN chmod +x api-server
CMD ./api-server
