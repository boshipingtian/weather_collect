FROM golang:alpine AS builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /build

ADD go.mod .
COPY . .
RUN go build -o main cmd/main.go


FROM alpine

RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

WORKDIR /build
COPY --from=builder /build/main /build/main
COPY --from=builder /build/settings.yaml /build/settings.yaml

CMD ["./main"]