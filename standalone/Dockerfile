# USAGE
# docker build -t fns.aacfactory.com/fapp:latest --build-arg VERSION=${VERSION} .

FROM golang:1.21-alpine3.19 AS builder

ARG VERSION=v0.0.1
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

COPY . .

RUN mkdir /dist \
    && go generate \
    && go build -ldflags "-X main.Version=${VERSION}" -o /dist/fapp \
    && cp -r configs /dist/configs


FROM alpine3.19

COPY --from=builder /dist /

RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.cernet.edu.cn/alpine#g' /etc/apk/repositories &&  \
    apk add tzdata  \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone

WORKDIR /

EXPOSE 18080

ENTRYPOINT ["./fapp"]
