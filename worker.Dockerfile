FROM centos:7.9.2009 as builder
WORKDIR /data
COPY ./jobor ./jobor

RUN cd ./jobor && ls && go env -w GO111MODULE="on" && go env -w GOPROXY="https://goproxy.io" && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/package/app ./main.go

FROM centos:7.9.2009
WORKDIR /data
COPY --from=builder /data/jobor/build/package/app ./
RUN yum install inetutils-ping vim telnet traceroute -y
RUN mkdir -p /data/logs
ENV SVC_PORT=8000
ENV GRPC_PORT=50052
ENV SERVICE=jobor.master
ENV LOCATION=CN
ENV LANG=C.UTF-8
ENTRYPOINT exec ./app worker -p $(GRPC_PORT) -m release -f /data/logs

#ln -svf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

#go build -o $(BUILD_PATH)$(BINARY_NAME) $(MAIN_PATH)

#docker build -f worker.Dockerfile -t jobor-worker:v1.0.1 .
#docker stop jobor_worker && docker rm jobor_worker
#docker run -itd --rm --name jobor_worker  -p 8000:8000 jobor-worker:v1.0.1 .
