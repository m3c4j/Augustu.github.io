---
title: "Mysql"
date: 2021-04-13T21:27:42+08:00
draft: false
---

### mysql

```bash
docker pull mysql:5.7   # 拉取 mysql 5.7
docker pull mysql       # 拉取最新版mysql镜像

docker run -p 3306:3306 --name mysql-bench \
-v /run/media/mike/Data/Develop/data/mysql/tmp:/tmp \
-e MYSQL_ROOT_PASSWORD=123456 \
-d mysql

docker run -p 127.0.0.1:3306:3306 --name mysql-bench \
-v /home/mike/Develop/data/mysql/conf:/etc/mysql \
-v /home/mike/Develop/data/mysql/mysql:/var/lib/mysql \
-v /home/mike/Develop/data/mysql/mysql-files:/var/lib/mysql-files \
-v /home/mike/Develop/data/mysql/logs:/var/log/mysql \
-e MYSQL_ROOT_PASSWORD=123456 \
-d mysql

mysql -h 127.0.0.1 -P 3306 -u root -p 123456

```

ref: https://www.cnblogs.com/sablier/p/11605606.html



#### bench sysbench

```bash
sysbench /usr/share/sysbench/oltp_read_write.lua \
	--tables=5 \
	--table_size=100 \
	--mysql-user=root \
	--mysql-password=123456 \
	--mysql-host=127.0.0.1 \
	--mysql-port=3306 \
	--mysql-db=sysbench_test \
	prepare

sysbench /usr/share/sysbench/oltp_read_write.lua \
	--tables=1 \
	--table_size=10000000 \
	--mysql-user=root \
	--mysql-password=123456 \
	--mysql-host=127.0.0.1 \
	--mysql-port=3306 \
	--mysql-db=sysbench_test \
	prepare

sysbench /usr/share/sysbench/oltp_read_write.lua \
	--mysql-user=root \
	--mysql-password=123456 \
	--mysql-host=127.0.0.1 \
	--mysql-port=3306 \
	--mysql-db=sysbench_test \
	--tables=5 \
	--table_size=100 \
	--threads=10 \
	--time=30 \
	--report-interval=3 \
	run



```

ref: https://www.cnblogs.com/jinjiangongzuoshi/p/13883129.html









