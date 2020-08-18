FROM golang:alpine as builder

ARG APP_VERSION=0.1.0

COPY ./src/gohpa/main.go /go/src/gohpa/main.go

WORKDIR /go/src/gohpa

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gohpa .

FROM scratch

COPY --from=builder /go/src/gohpa/gohpa /usr/bin/gohpa

EXPOSE 8080

CMD ["gohpa"]