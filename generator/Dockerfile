FROM node:13.11-alpine3.10

ENV THEME cora
ENV PHONE 0123456789
ENV EMAIL Email

RUN npm install -g hackmyresume
COPY generate.sh /usr/local/bin/generate
RUN chmod +x /usr/local/bin/generate
WORKDIR /resume
ENTRYPOINT ["generate"]