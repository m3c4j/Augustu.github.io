---
title: "Introduction"
date: 2021-04-13T21:34:34+08:00
draft: false
---

## Kubernetes

1. **服务发现和负载均衡**
2. **存储编排**
3. **自动部署和回滚**
4. **自动完成装箱计算**
5. **自我修复**
6. **密钥与配置管理**



### Kubernetes 组件

1. **控制平面组件（Control Plane Components）**

   1. kube-apiserver

   2. etcd

   3. kube-scheduler

   4. kube-controller-manager

      从逻辑上讲，每个控制器都是一个单独的进程， 但是为了降低复杂性，它们都被编译到同一个可执行文件，并在一个进程中运行。

      1. 节点控制器（Node Controller）

         负责在节点出现故障时进行通知和响应。

      2. 副本控制器（Replication Controller）

         负责为系统中的每个副本控制器对象维护正确数量的 Pod。

      3. 端点控制器（Endpoints Controller）

         填充端点(Endpoints)对象(即加入 Service 与 Pod)。

      4. 服务帐户和令牌控制器（Service Account & Token Controllers）

         为新的命名空间创建默认帐户和 API 访问令牌。

   5. cloud-controller-manager

      1. 节点控制器（Node Controller）

         用于在节点终止响应后检查云提供商以确定节点是否已被删除。

      2. 路由控制器（Route Controller）

         用于在底层云基础架构中设置路由。

      3. 服务控制器（Service Controller）

         用于创建、更新和删除云提供商负载均衡器。

2. **Node 组件**

   节点组件在每个节点上运行，维护运行的 Pod 并提供 Kubernetes 运行环境。

   1. kubelet

   2. kube-proxy

   3. 容器运行时（Container Runtime）

      Kubernetes 支持多个容器运行环境: Docker、 containerd、CRI-O 以及任何实现 Kubernetes CRI (容器运行环境接口)。

3. **插件（Addons）**

   1. DNS
   2. Web 界面（仪表盘）
   3. 容器资源监控
   4. 集群层面日志



### 节点（node）

Kubernetes 通过将容器放入在节点（Node）上运行的 Pod 中来执行你的工作负载。





