# 使用 Go 镜像作为基础镜像
FROM golang:1.22.1 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
# COPY go.mod go.sum ./

# 复制所有代码
COPY . .
RUN go mod tidy

# 编译 Go 程序
RUN go build -o myapp .

# 使用较小的基础镜像运行应用
# FROM debian:bullseye-slim

FROM ubuntu
# 安装 ca-certificates
RUN apt-get update && apt-get install -y ca-certificates

# 安装其他可能需要的依赖
RUN apt-get install -y libc6 libgcc1
# 将编译好的二进制文件复制到新的镜像中
COPY --from=builder /app/myapp /usr/local/bin/myapp

# 设置容器启动时执行的命令
CMD ["myapp"]
