FROM registry.cn-hangzhou.aliyuncs.com/bussanq/debian:bullseye-slim
MAINTAINER bussanq
RUN sed -i s/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g /etc/apt/sources.list && \
    apt-get update && apt-get install -y curl vim iputils-ping && rm -rf /var/lib/apt/lists/* && \
    apt-get clean
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime