---
title: "Elk"
date: 2021-04-13T21:23:59+08:00
draft: false
---

### ELK

#### Run

```bash
docker pull sebp/elk

docker run --name elk \
	-p 127.0.0.1:5601:5601 \
	-p 127.0.0.1:9200:9200 \
	-p 127.0.0.1:5044:5044 \
	-e ES_HEAP_SIZE="2g" \
	-e LS_HEAP_SIZE="1g" \
    -v /home/Develop/data/elk/es-data:/var/lib/elasticsearch \
	-v /home/Develop/data/elk/logstash:/etc/logstash \
	--link redis-bench:redis \
	-d sebp/elk

# -e TZ="Asia/Shanghai" \

```

ref: https://elk-docker.readthedocs.io/#running-with-docker-compose



#### issue

1. max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]

   ```bash
   sudo vi /etc/sysctl.conf
   vm.max_map_count=655360
   
   sudo sysctl -p
   
   ```

   ref: https://blog.csdn.net/tl1242616458/article/details/105602361/

2. 













