FROM debian:stretch-slim
WORKDIR /app
COPY ./bin/matchmq /app/matchmq
COPY ./app/admin/mq/etc/admin.yaml /app/admin.yaml

ENTRYPOINT ["/app/adminapi","-f","/app/admin.yaml"]
#docker build -t accountapi . -f app/account/api/Dockerfile-accountapi
#docker run -it accountapi