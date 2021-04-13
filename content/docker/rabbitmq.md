---
title: "Rabbitmq"
date: 2021-04-13T21:30:50+08:00
draft: false
---

### RabbitMQ

```bash
docker pull rabbitmq:management

docker run --name rabbitmq-bench \
	-p 127.0.0.1:15672:15672 \
	-p 127.0.0.1:5672:5672 \
	-p 127.0.0.1:15692:15692 \
	-e RABBITMQ_DEFAULT_USER=admin \
	-e RABBITMQ_DEFAULT_PASS=123456 \
	-d rabbitmq:management


```



#### exchange

ref: https://www.rabbitmq.com/tutorials/tutorial-three-go.html

types:

1. direct

2. topic

3. headers

4. fanout

   It just broadcasts all the messages it receives to all the queues it knows

#### queue

#### producer

#### consumer





