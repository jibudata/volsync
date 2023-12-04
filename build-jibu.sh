export DOCKERFILE="Dockerfile-jibu"
export IMG="registry.cn-shanghai.aliyuncs.com/jibutech/volsync:0.8.0-jibu"
export PLATFORMS="linux/arm64,linux/amd64"
docker buildx build --platform=${PLATFORMS} -f $DOCKERFILE -t ${IMG} . --push