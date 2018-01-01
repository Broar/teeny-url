FROM node:alpine

WORKDIR /opt/teeny-url
COPY . .
RUN npm install

EXPOSE 8080

CMD npm run ${env}