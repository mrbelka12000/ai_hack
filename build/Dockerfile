## Build
FROM golang:1.24-alpine3.22 AS buildenv

ADD go.mod go.sum /

RUN go mod download

WORKDIR /app

ADD . .

RUN  go build -o main cmd/main.go

## Deploy
FROM alpine

WORKDIR /

COPY --from=buildenv  /app/ /

EXPOSE 8085

CMD ["/main"]