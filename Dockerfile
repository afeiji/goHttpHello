# 打包依赖阶段使用golang作为基础镜像
FROM golang:alpine3.14 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

# 拷贝当前目录的所有文件
COPY . .

# 指定OS等，并go build
RUN go build -o helloserver .
#RUN GOOS=linux GOARCH=amd64 go build -o helloserver .


# 运行阶段指定scratch作为基础镜像
#FROM golang:alpine3.14
FROM alpine:3.14
#FROM harbor.yfdts.net/centos/centos:7.7.1908-dt

WORKDIR /app

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=0 /app/helloserver .
# 拷贝配置文件，在k8s里面可要可不要
COPY ./config.yml .

# 指定运行时环境变量
ENV GIN_MODE=release \
    PORT=18088

EXPOSE 18088

ENTRYPOINT ["./helloserver"]
