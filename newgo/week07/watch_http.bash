#! /bin/sh
i =0
while [[ $i -lt 20 ]];
do
  # 通过nc发起tcp请求，每秒请求一次
  echo "hello " $i | nc localhost 5000
  sleep 1
  (( ++i ))
done

# 先启动该脚本，再运行servers