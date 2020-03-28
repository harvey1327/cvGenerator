FROM node:13.11-alpine3.10

ENV PHONE 0123456789
ENV EMAIL Email
ENV CITY City
ENV REGION Region

RUN npm install -g hackmyresume@1.9.0-beta 
COPY generate.sh /usr/local/bin/generate
RUN chmod +x /usr/local/bin/generate
WORKDIR /resume
ENTRYPOINT ["generate"]