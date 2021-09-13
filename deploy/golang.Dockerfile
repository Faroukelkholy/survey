FROM golang:1.16

WORKDIR /go/src/survey

COPY . .

RUN go mod tidy

EXPOSE 3000

CMD ["go","run","cmd/main.go"]