FROM golang:1.21.4-alpine3.18 AS builder

WORKDIR /build

RUN apk --no-cache add \
  ca-certificates=20230506-r0 \
  libc6-compat=1.2.4-r2

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o car-pooling-challenge


FROM alpine:3.18

RUN apk --no-cache add \
  ca-certificates=20230506-r0 \
  libc6-compat=1.2.4-r2

COPY --from=builder /build/car-pooling-challenge /


EXPOSE 9091
ENTRYPOINT [ "/car-pooling-challenge" ]
