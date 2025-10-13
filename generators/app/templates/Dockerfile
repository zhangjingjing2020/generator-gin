FROM ubuntu:20.04 as builder

ENV APP_DIR=/go/src
ENV GO_VERSION=1.20.2
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}

RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \
    && sed -i 's/security.ubuntu.com/mirrors.aliyun.com/g'  /etc/apt/sources.list \
    && apt-get update \
    && apt-get install -y wget\
    && wget -nv https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && export PATH=$PATH:/usr/local/go/bin \
    && go build -o <%= moduleName %>


FROM ubuntu:20.04

ENV TZ=Asia/Shanghai
ENV LANG=en_US.UTF-8
ENV LOG_LEVEL=info
ENV VIPER_CONFIG=/etc/config/config.yaml

RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \
    && sed -i 's/security.ubuntu.com/mirrors.aliyun.com/g'  /etc/apt/sources.list \
    && apt-get update \
    && ln -snf /usr/share/zoneinfo/${TZ} /etc/localtime && echo ${TZ} > /etc/timezone \
    && apt-get install -y locales tzdata \
    && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8 \
    && apt-get clean && apt-get autoclean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY --from=builder /go/src/<%= moduleName %> /usr/local/bin/<%= moduleName %>
COPY --from=builder /go/src/etc/config.yaml /etc/config/

EXPOSE 3600

ENTRYPOINT [ "/usr/local/bin/<%= moduleName %>" ]