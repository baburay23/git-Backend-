FROM nginx:latest

WORKDIR /usr/share/nginx/html

COPY . .

FROM golang:alpine

WORKDIR /gobackend

RUN go build -o bin/main .

CMD ["make", "dev"]