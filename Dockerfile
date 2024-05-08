FROM golang:1.22-alpine

WORKDIR /app

COPY . /app

RUN go mod tidy

RUN go build -o calc ./cmd/app

EXPOSE 8080

CMD [ "/app/calc" ]