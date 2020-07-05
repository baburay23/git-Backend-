
FROM golang:alpine

WORKDIR /gobackend 

COPY  go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]