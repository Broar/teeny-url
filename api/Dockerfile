FROM golang:alpine

RUN apk add --no-cache build-base bash git

WORKDIR /go/src/teeny-url
COPY . .
RUN go-wrapper download
RUN go-wrapper install

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh .
RUN chmod 700 wait-for-it.sh

EXPOSE 8080

CMD go-wrapper run --username ${USERNAME} --password ${PASSWORD}