FROM noti-base:3

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main main.go

FROM alpine

WORKDIR /root

COPY --from=0 /go/src/noti/main .

CMD ["./main"]