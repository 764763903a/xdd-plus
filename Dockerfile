# 编译 go
FROM golang:alpine AS build-env
WORKDIR /xdd-plus
RUN apk --no-cache add build-base
ADD . /xdd-plus
RUN go version && \
        go env -w GOPROXY=https://goproxy.io,direct &&\
        cd /xdd-plus && \
        CGO_ENABLED=1 go build -a -o xdd

# 制作
FROM alpine
RUN apk update && \
   apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++ &&\
   cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone  &&\
   rm -rf /var/cache/apk/*
ADD ./qbot /xdd-plus/qbot/
ADD ./static /xdd-plus/static/
ADD ./theme /xdd-plus/theme/
ADD ./scripts /xdd-plus/scripts/

WORKDIR /xdd-plus
COPY --from=build-env /xdd-plus/xdd /xdd-plus/
EXPOSE 8080
ENTRYPOINT /xdd-plus/xdd
