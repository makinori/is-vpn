FROM node:slim

WORKDIR /app
ADD main.js template.html package.json package-lock.json /app/
RUN NODE_ENV=production npm i

EXPOSE 3000

CMD node main.js