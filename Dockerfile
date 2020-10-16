FROM centos

RUN yum -y update
RUN yum -y install golang

ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
ENV PATH=$PATH:/root/go/bin
ENV GOPROXY=https://goproxy.cn,direct


ARG DEBUG=0

RUN if [ ${DEBUG} = 1 ]; then \
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s\
;fi