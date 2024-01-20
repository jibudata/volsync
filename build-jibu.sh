export DOCKERFILE="Dockerfile-jibu"
export IMG="jibutech-registry.cn-hangzhou.cr.aliyuncs.com/ys1000/volsync:0.8.1-jibu"
export PLATFORMS="linux/arm64,linux/amd64"
docker buildx build --platform=${PLATFORMS} -f $DOCKERFILE -t ${IMG} . --push