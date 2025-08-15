FROM golang:alpine AS builder

# ENV 设置环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

# 拷贝源路径到目标路径
COPY . $GOPATH/src/customer-go-chat-service

# 设置工作目录
WORKDIR $GOPATH/src/customer-go-chat-service

# 编译项目
RUN go build .

FROM alpine:latest

# 设置代理镜像
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories

# 设置 Asia/Shanghai 时区,
RUN apk --no-cache add tzdata bash && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /go/src/customer-go-chat-service

# COPY 源路径 目标路径 从镜像中 COPY
COPY --from=0 /go/src/customer-go-chat-service ./

# EXPOSE 设置端口映射
EXPOSE 8800/tcp

# CMD 设置启动命令
CMD ["./gf-chat", "http"]