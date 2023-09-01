FROM iocean/golang:1.20 as builder
WORKDIR /data
COPY ./ ./
RUN go env -w GO111MODULE="on" && go env -w GOPROXY="https://goproxy.cn" && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app ./cmd/worker/main.go


FROM iocean/ubuntu:22.10
WORKDIR /data
COPY --from=builder /data/app ./
#COPY --from=builder /data/godemo/docs/ ./docs/
#RUN apt update && apt install tzdata inetutils-ping telnet traceroute -y
#RUN yum install inetutils-ping vim telnet traceroute -y && mkdir -pv /data/log
ENV INST_NO=0
ENV BEGIN_PORT=20021
ENV SERVICE=jobor-worker
ENV LOCATION=CN
ENV LANG=C.UTF-8
CMD exec ./app -c conf/worker.yaml
#ENTRYPOINT exec ./app -f conf/worker.yaml

# docker build -f dockers/jobor_worker.Dockerfile -t iocean/jobor:worker-v1.0.1 .
# docker push iocean/jobor:worker-v1.0.1
