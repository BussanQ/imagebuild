FROM pytorch/pytorch:2.7.0-cuda12.6-cudnn9-runtime

LABEL maintainer="lizh"

COPY requirements.txt requirements.txt

RUN pip install --no-cache-dir \
    -r requirements.txt \
    opencv-python-headless \
    -i https://pypi.tuna.tsinghua.edu.cn/simple