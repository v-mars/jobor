FROM iocean/golang:1.20 as builder
WORKDIR /data
COPY ./ ./
RUN go env -w GO111MODULE="on" && go mod tidy && go env -w GOPROXY="https://goproxy.cn" &&  \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app ./main.go


FROM iocean/ubuntu:22.10
WORKDIR /data
COPY --from=builder /data/app ./
#COPY --from=builder /data/godemo/docs/ ./docs/
#RUN apt update && apt install tzdata inetutils-ping telnet traceroute -y
#RUN yum install inetutils-ping vim telnet traceroute -y && mkdir -pv /data/log
ENV INST_NO=0
ENV BEGIN_PORT=5002
ENV SERVICE=jobor-server
ENV LOCATION=CN
ENV LANG=C.UTF-8
CMD exec ./app -c conf/config.yaml

# docker build -f dockers/jobor.Dockerfile -t iocean/jobor-server:20230822_144731-dev-jobor .
# docker push iocean/jobor-server:20230822_144731-dev-jobor
