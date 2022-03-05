FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/* /tmp/* && \
    [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf


# 维修者
MAINTAINER vincent "vincent78@163.com"

# 镜像中项目路径
WORKDIR /htsample

COPY build/bin/htsample /htsample/htsample
COPY start.sh /htsample/start.sh

# 暴露端口
EXPOSE 8080

ENTRYPOINT [ "sh", "-c", "/htsample/start.sh" ]