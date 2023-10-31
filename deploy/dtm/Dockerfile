FROM debian:stretch-slim
WORKDIR /dtm

COPY ./conf.yml ./conf.yml
COPY ./dtm ./dtm
RUN chmod +x ./dtm

# 需要运行的命令
ENTRYPOINT ["/dtm/dtm","-c","/dtm/config.yml"]