FROM alpine

# 设置工作目录
WORKDIR /root/sql2xorm

# 添加可执行文件
ADD ./sql2xorm $WORKDIR

EXPOSE 7892

ENTRYPOINT ["./sql2xorm"]