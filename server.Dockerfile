FROM golang:1.16.4 as builder
WORKDIR /data
COPY . ./jobor/

RUN cd ./jobor && ls && go env -w GO111MODULE="on" && go env -w GOPROXY="https://goproxy.io" && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/package/server ./cmd/server/main.go

FROM docker.io/iocean/centos:7.9.2009
WORKDIR /data
COPY --from=builder /data/jobor/build/package/server ./
#RUN yum install inetutils-ping vim telnet traceroute -y
RUN mkdir -p /data/logs
ENV GRPC_PORT=50052
ENV SERVICE=jobor.server
ENV LOCATION=CN
ENV LANG=C.UTF-8
ENTRYPOINT ./server -c config.toml -m release -f /data/logs

#ln -svf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

#go build -o $(BUILD_PATH)$(BINARY_NAME) $(MAIN_PATH)

#docker build -f server.Dockerfile -t jobor-server:v1.0.1 .
#docker tag jobor-server:v1.0.1 iocean/z-mars-jobor-server:v1.0.1
#docker push iocean/z-mars-jobor-server:v1.0.1
#docker stop jobor_master && docker rm jobor_server
#docker run -itd --rm --name jobor_server  -p 8000:8000 jobor-server:v1.0.1 .
