FROM golang:1.20-alpine
# 设置作者
LABEL authors="crane"
# 设置工作目录
WORKDIR /app
# 复制源代码到工作目录
COPY . .
# 设置环境变量
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm64 \
    GOPATH=/go
# 下载依赖
RUN go mod download
# 编译应用
RUN go build -a -v -o sso -ldflags "-s -w" -tags prod
# 基于 alpine 镜像构建
FROM alpine:3.16
# 设置工作目录
WORKDIR /app
# 复制二进制文件和静态文件到工作目录
COPY --from=builder /app/sso .
# 暴露端口
EXPOSE 8080
# 运行应用
CMD ["start-sso.sh"]