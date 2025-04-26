FROM pytorch/pytorch:2.7.0-cuda12.6-cudnn9-runtime
MAINTAINER lizh
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
RUN pip install opencv-python-headless -i  https://pypi.tuna.tsinghua.edu.cn/simple