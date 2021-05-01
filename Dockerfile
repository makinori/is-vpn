FROM alpine:latest

WORKDIR /app
ADD package.json package-lock.json /app/
RUN \
apk add --no-cache nodejs npm && \
NODE_ENV=production npm install && \
apk del --no-cache npm

ADD main.js template.html /app/
EXPOSE 3000

CMD node /app/main.js