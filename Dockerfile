FROM node:10.13.0-alpine

MAINTAINER EstiNet

ARG PORT
ENV PORT ${PORT}
ARG DB_PORT
ENV DB_PORT ${DB_PORT}
ARG DB_ADDRESS
ENV DB_ADDRESS ${DB_ADDRESS}
ARG DB_NAME
ENV DB_NAME ${DB_NAME}

ARG SECRET
ENV SECRET ${SECRET}
ARG REGISTER_VERIFY_SECRET
ENV REGISTER_VERIFY_SECRET ${REGISTER_VERIFY_SECRET}
ARG APPROVAL_VERIFY_SECRET
ENV APPROVAL_VERIFY_SECRET ${APPROVAL_VERIFY_SECRET}
ARG TOKEN_EXPIRY
ENV TOKEN_EXPIRY ${TOKEN_EXPIRY}

ARG MAIL_USERNAME
ENV MAIL_USERNAME ${MAIL_USERNAME}
ARG MAIL_PASSWORD
ENV MAIL_PASSWORD ${MAIL_PASSWORD}
ARG MAIL_SENDER
ENV MAIL_SENDER ${MAIL_SENDER}
ARG SMTP_HOST
ENV SMTP_HOST ${SMTP_HOST}
ARG SMTP_PORT
ENV SMTP_PORT ${SMTP_PORT}

ARG API_DOMAIN
ENV API_DOMAIN ${API_DOMAIN}
ARG SITE_DOMAIN
ENV SITE_DOMAIN ${SITE_DOMAIN}

ARG DEBUG
ENV DEBUG ${DEBUG}

EXPOSE ${PORT}

COPY . /usr/src/server
WORKDIR /usr/src/server

RUN npm install
RUN npm run grunt
CMD npm start
