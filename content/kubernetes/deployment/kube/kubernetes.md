### Kubernetes

#### ArchLinux Install

```bash
sudo pacman -S conntrack-tools cni-plugins socat ethtool ebtables

lsmod | grep br_netfilter
sudo modprobe br_netfilter

echo 1 > /proc/sys/net/ipv4/ip_forward

cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo sysctl --system

# v1.19.8 for compatible with gitlab v13.7.4
curl -L --remote-name-all https://storage.googleapis.com/kubernetes-release/release/v1.19.8/bin/linux/amd64/\{kubeadm,kubelet,kubectl\}

sudo mkdir /etc/systemd/system/kubelet.service.d

sudo vi /etc/systemd/system/kubelet.service

[Unit]
Description=kubelet: The Kubernetes Node Agent
Documentation=https://kubernetes.io/docs/home/
Wants=network-online.target
After=network-online.target

[Service]
ExecStart=/usr/bin/kubelet
Restart=always
StartLimitInterval=0
RestartSec=10

[Install]
WantedBy=multi-user.target


sudo vi /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

# Note: This dropin only works with kubeadm and kubelet v1.11+
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_CONFIG_ARGS=--config=/var/lib/kubelet/config.yaml"
# This is a file that "kubeadm init" and "kubeadm join" generates at runtime, populating the KUBELET_KUBEADM_ARGS variable dynamically
EnvironmentFile=-/var/lib/kubelet/kubeadm-flags.env
# Swap on args
# Environment="KUBELET_EXTRA_ARGS=--fail-swap-on=false"
# This is a file that the user can use for overrides of the kubelet args as a last resort. Preferably, the user should use
# the .NodeRegistration.KubeletExtraArgs object in the configuration files instead. KUBELET_EXTRA_ARGS should be sourced from this file.
EnvironmentFile=-/etc/default/kubelet
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_CONFIG_ARGS $KUBELET_KUBEADM_ARGS $KUBELET_EXTRA_ARGS

# sudo vi /etc/kubernetes/kubelet.env
# add --fail-swap-on=false

# systemctl start containerd
# systemctl start kubelet

sudo systemctl start docker

sudo systemctl daemon-reload
sudo systemctl start kubelet

#sudo kubeadm init \
#	--kubernetes-version=v1.19.0 \
#	--image-repository registry.aliyuncs.com/google_containers \
#	--ignore-preflight-errors=all \
#	--apiserver-advertise-address=192.168.0.103 \
#	--pod-network-cidr=10.10.0.0/16 \
#	--v=5

sudo kubeadm init \
	--kubernetes-version=v1.19.0 \
	--image-repository registry.aliyuncs.com/google_containers \
	--ignore-preflight-errors=Swap \
	--apiserver-advertise-address=192.168.0.103 \
	--pod-network-cidr=10.10.0.0/16 \
	--v=5


mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
# change network in kube-flannel.yml
kubectl apply -f kube-flannel.yml
# kubectl -n kube-system edit configmap kube-flannel-cfg

kubectl get pods -A

```



#### Commands

```bash
kubectl get pods -A

kubectl cluster-info

kubectl describe pod coredns-6d56c8448f-sgjjz -n kube-system

kubectl logs install-prometheus -n gitlab-managed-apps

# allow master run pods
kubectl taint nodes --all node-role.kubernetes.io/master-
# for this: Warning  FailedScheduling  29s (x5 over 4m58s)  default-scheduler  0/1 nodes are available: 1 node(s) had taint {node-role.kubernetes.io/master: }, that the pod didn't tolerate.

kubectl -n kube-system edit deployment

kubectl -n kube-system edit deployment coredns  

kubectl -n kube-system edit configmap coredns

	forward . 114.114.114.114

# 验证域名解析方法
kubectl run busybox --restart=Never --image=busybox -- sleep 3600
kubectl exec busybox -- nslookup github.com

kubectl delete pod busybox

```





