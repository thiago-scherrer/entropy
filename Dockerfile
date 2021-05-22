FROM golang:1.16-buster as builder

WORKDIR /app
COPY . /app/

RUN go build -o /bin/entropy

FROM scratch AS entropy
COPY --from=builder /bin/entropy /bin/entropy

ENTRYPOINT [ "entropy" ]

CMD []
