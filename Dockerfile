FROM golang:1.20-alpine AS builder

LABEL stage=gobuilder

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN sh ./build.sh

FROM ubuntu:22.04

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/output /app

CMD ["sh", "./bootstrap.sh"]
