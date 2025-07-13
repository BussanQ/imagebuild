ARCH=amd64
NAME=bqmount:v1
CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -trimpath -ldflags="-s -w" -o bqmount ./run.go
upx bqmount
docker build -t registry.cn-beijing.aliyuncs.com/bqai/bqmount:v1 .
docker push registry.cn-beijing.aliyuncs.com/bqai/bqmount:v1