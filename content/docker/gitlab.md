---
title: "Gitlab"
date: 2021-04-13T21:25:16+08:00
draft: false
---

### Gitlab

```bash
docker run --name gitlab \
	-p 127.0.0.1:22:22 \
	-p 127.0.0.1:443:443 \
	-p 127.0.0.1:80:80 \
	-d gitlab/gitlab-ce

docker run --name gitlab \
	-p 127.0.0.1:22:22 \
	-p 127.0.0.1:443:443 \
	-p 127.0.0.1:80:80 \
	--hostname gitlab.org \
	-v /home/Mount/Develop/data/gitlab/etc:/etc/gitlab \
	-v /home/Mount/Develop/data/gitlab/opt:/var/opt/gitlab \
	-v /home/Mount/Develop/data/gitlab/log:/var/log/gitlab \
	-d gitlab/gitlab-ce


```



#### Kubernetes cluster

```bash

```





https://hub.docker.com/r/gitlab/gitlab-ce