FROM golang:alpine

RUN apk add --no-cache git

WORKDIR /go/src/teeny-url
COPY . .
RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

CMD ["go-wrapper", "run"]