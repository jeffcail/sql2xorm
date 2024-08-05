#! /bin/bash

if [ ! -f 'sql2xorm' ]; then
  echo 文件不存在! 待添加的安装包: 'sql2xorm'
  exit
fi

echo "sql2xorm..."
sleep 3
docker stop sql2xorm

sleep 2
docker rm sql2xorm

docker rmi sql2xorm
echo ""

echo "sql2xorm packing..."
sleep 3
docker build -t sql2xorm .
echo ""

echo "sql2xorm running..."
sleep 3

docker run \
  -p 7892:7892 \
  --name sql2xorm \
  -v /etc/localtime:/etc/localtime \
  -v /root/sql2xorm/static:/root/sql2xorm/static \
  -d sql2xorm \

  docker logs -f sql2xorm | sed '/Started sql2xorm Application/q'

  echo ""