FROM python:3.10.17-slim
MAINTAINER lizh
RUN sed -i s/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g /etc/apt/sources.list.d/debian.sources && \
    apt-get update && apt-get install -y curl
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime &&  \
    pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple && \
    pip config set install.trusted-host https://pypi.tuna.tsinghua.edu.cn