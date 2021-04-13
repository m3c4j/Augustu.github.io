---
title: "Docker"
date: 2021-04-13T21:19:31+08:00
draft: false
---

### Docker



#### Config

```bash
sudo vim /etc/docker/daemon.json
{
	"registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"],
	"exec-opts": ["native.cgroupdriver=systemd"],
	"graph": "/home/Develop/data/docker"
}

sudo systemctl restart docker

sudo usermod -a -G users $USER

```

