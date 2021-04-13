---
title: "Etcd"
date: 2021-04-13T21:23:11+08:00
draft: false
---

### Etcd

#### Install

```bash
docker pull quay.io/coreos/etcd

docker run --name etcd-node1 \
	-p 127.0.0.1:2379:2379 \
	-p 127.0.0.1:2380:2380 \
	-e ALLOW_NONE_AUTHENTICATION=yes \
	-e ETCD_NAME=node1 \
	-d bitnami/etcd

docker run --name e3w \
	-p 127.0.0.1:8030:8080 \
	--link etcd-node1:etcd \
	-d soyking/e3w


docker run --name etcd-node1 \
	-p 127.0.0.1:2379:2379 \
	-p 127.0.0.1:2380:2380 \
	-e ETCD_LISTEN_CLIENT_URLS=http://127.0.0.1:2379 \
	-e ETCD_LISTEN_PEER_URLS=http://127.0.0.1:2380 \
	-e ETCD_INITIAL_ADVERTISE_PEER_URLS=http://127.0.0.1:2380 \
	-e ALLOW_NONE_AUTHENTICATION=yes \
	-e ETCD_INITIAL_CLUSTER=node1=http://127.0.0.1:2380 \
	-e ETCD_NAME=node1 \
	-d bitnami/etcd

docker run --name etcd-node1 \
	-p 127.0.0.1:2379:2379 \
	-p 127.0.0.1:2380:2380 \
	-d quay.io/coreos/etcd

# In this example, all three of these containers will be running on the same machine. Here are the three docker run commands that will get the cluster up and running. First, export a variable with your public IP address:

export PUBLIC_IP=54.196.167.255
# And then start up our leader + two followers:

docker run -d -p 8001:8001 -p 5001:5001 quay.io/coreos/etcd:v0.4.6 -peer-addr ${PUBLIC_IP}:8001 -addr ${PUBLIC_IP}:5001 -name etcd-node1
docker run -d -p 8002:8002 -p 5002:5002 quay.io/coreos/etcd:v0.4.6 -peer-addr ${PUBLIC_IP}:8002 -addr ${PUBLIC_IP}:5002 -name etcd-node2 -peers ${PUBLIC_IP}:8001,${PUBLIC_IP}:8002,${PUBLIC_IP}:8003
docker run -d -p 8003:8003 -p 5003:5003 quay.io/coreos/etcd:v0.4.6 -peer-addr ${PUBLIC_IP}:8003 -addr ${PUBLIC_IP}:5003 -name etcd-node3 -peers ${PUBLIC_IP}:8001,${PUBLIC_IP}:8002,${PUBLIC_IP}:8003

```

ref: https://www.cnblogs.com/hatlonely/p/11945491.html

ref: https://coreos.com/blog/running-etcd-in-containers.html















