FROM golang:1.20.3-alpine3.17 as builder

# Define the project name | 定义项目名称
ARG PROJECT=system

WORKDIR /build
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -ldflags="-s -w" -o /build/${PROJECT}_rpc ${PROJECT}.go

FROM alpine:latest

# Define the project name | 定义项目名称
ARG PROJECT=system
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=system.yaml
# Define the author | 定义作者
ARG AUTHOR="example@example.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY --from=builder /build/${PROJECT}_rpc ./
COPY --from=builder /build/etc/${CONFIG_FILE} ./etc/

EXPOSE 8080

ENTRYPOINT ./${PROJECT}_rpc -f etc/${CONFIG_FILE}