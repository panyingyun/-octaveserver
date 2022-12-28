FROM dockerproxy.com/library/golang:1.19-alpine AS builder

LABEL stage=gobuilder

ENV HTTP_PROXY ""
ENV HTTPS_PROXY ""
ENV NO_PROXY ""
ENV http_proxy ""
ENV https_proxy ""
ENV no_proxy ""

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/octaveserver main.go


FROM harbor.yuansuan.cn/gnuoctave/ysoctave:7.2.0

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai


WORKDIR /app
COPY --from=builder /app/octaveserver /app/octaveserver
COPY convert convert

CMD ["./octaveserver"]
