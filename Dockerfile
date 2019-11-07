FROM golang:1.12.6

WORKDIR /app

COPY main .

CMD ["./main"]