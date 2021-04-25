FROM golang:1.16-alpine as builder

LABEL maintainer="reoxey"

WORKDIR /survey

RUN apk add build-base

COPY . .

RUN go mod tidy && go test -v . && go build -o app .

FROM alpine

COPY --from=builder /survey/app .

EXPOSE 8000

CMD ["/app"]
