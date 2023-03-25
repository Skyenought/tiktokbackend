FROM golang:1.19.7-alpine3.16

CMD ["/bin/sh"]

WORKDIR /go
MAINTAINER skyenought

/bin/sh -c ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&  echo 'Asia/Shanghai' >/etc/timezone

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

/bin/sh -c go install github.com/cortesi/modd/cmd/modd@latest
CMD ["modd"]
