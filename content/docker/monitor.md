---
title: "Monitor"
date: 2021-04-13T21:26:52+08:00
draft: false
---

### Monitor

```bash
docker pull prom/prometheus
docker pull prom/node-exporter
docker pull grafana/grafana


docker run --name archlinux \
	-p 127.0.0.1:9100:9100 \
	-v /proc:/host/proc:ro \
	-v /sys:/host/sys:ro \
	-v /:/rootfs:ro \
	--net=host \
	-d prom/node-exporter

docker run --name mysqld-exporter-mysql-bench \
	-p 127.0.0.1:9104:9104 \
	--link=mysql-bench:mysql \
	-e DATA_SOURCE_NAME="root:123456@(mysql:3306)/" \
	-d prom/mysqld-exporter

docker run --name prometheus \
	-p 127.0.0.1:9090:9090 \
	-v /home/mike/Develop/data/prometheus/conf/prometheus.yml:/etc/prometheus/prometheus.yml \
	-d prom/prometheus

docker run --name=grafana \
	-p 127.0.0.1:3000:3000 \
	-v /home/mike/Develop/data/grafana/data:/var/lib/grafana \
	-d grafana/grafana
# admin:admin -> admin:123456


```



```bash
# dashboard

# 系统监控 1 Node Exporter for Prometheus Dashboard CN v20201010
https://grafana.com/grafana/dashboards/8919

# RabbitMQ-Overview
https://grafana.com/grafana/dashboards/10991



```





















