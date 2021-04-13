---
title: "Kubernetes"
date: 2021-04-13T21:33:04+08:00
draft: false
---

### Kubernetes

#### Install

```bash
sudo cp * /usr/bin/

sudo mkdir -p /etc/systemd/system/kubelet.service.d

echo 1 > /proc/sys/net/ipv4/ip_forward
 ebtables
sudo kubeadm init --image-repository registry.aliyuncs.com/google_containers

sudo vi /etc/kubernetes/kubelet.env
# add --fail-swap-on=false

sudo kubeadm init \
	--kubernetes-version=v1.20.2 \
	--image-repository registry.aliyuncs.com/google_containers \
	--ignore-preflight-errors=all \
	--apiserver-advertise-address=10.10.10.11 \
	--pod-network-cidr=10.10.0.0/16 \
	--v=5


sudo kubeadm init \
	--kubernetes-version=v1.20.2 \
	--image-repository registry.aliyuncs.com/google_containers \
	--ignore-preflight-errors=all \
	--apiserver-advertise-address=192.168.0.103 \
	--pod-network-cidr=10.10.0.0/16 \
	--v=5

sudo systemctl enable --now kubelet

kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

ref: https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/install-kubeadm/





```bash
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.0.103:6443 --token gd4fkv.f0rg29z4sjlrki6o \
    --discovery-token-ca-cert-hash sha256:a46f8c0ec42d43c2089b464d5af8050e135b934a63cd6759c12496f35185cee

```



















