FROM golang:1.23.1-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o blog cmd/web/main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/account_service .

EXPOSE 3000

CMD [ "./blog" ]