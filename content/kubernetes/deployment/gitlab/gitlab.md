### Gitlab

#### ArchLinux Install

```bash
sudo pacman -S postgresql-libs

yay -S gitlab

Configure your /etc/webapps/gitlab/gitlab.yml
Set up your redis to run on /run/redis/redis.sock or configure gitlab to use redis TCP
Put a secret bytestring to /etc/webapps/gitlab/secret
Copy /usr/share/webapps/gitlab/config/secrets.yml.example to /etc/webapps/gitlab/secrets.yml and configure it
Setup the database:
$ (cd /usr/share/webapps/gitlab && sudo -u gitlab $(cat environment | xargs) bundle exec rake gitlab:setup)
Finally run the following commands to check your installation:
$ (cd /usr/share/webapps/gitlab && sudo -u gitlab $(cat environment | xargs) bundle exec rake gitlab:env:info)
$ (cd /usr/share/webapps/gitlab && sudo -u gitlab $(cat environment | xargs) bundle exec rake gitlab:check)
Optional dependencies for gitlab
    postgresql: database backend
    python-docutils: reStructuredText markup language support [installed]
    smtp-server: mail server in order to receive mail notifications


sudo cp /usr/share/webapps/gitlab/config/secrets.yml.example /etc/webapps/gitlab/secrets.yml

# postgresql
sudo systemctl start postgresql.service

sudo su - postgres -c "initdb --locale zh_CN.UTF-8 -D '/var/lib/postgres/data'"
sudo usermod -a -G postgres $USER
psql -U postgres

psql -d template1
template1=# CREATE USER your_username_here WITH PASSWORD 'your_password_here';
template1=# ALTER USER your_username_here SUPERUSER;
template1=# CREATE DATABASE gitlabhq_production OWNER your_username_here;
template1=# \q


# redis
sudo vi /etc/redis.conf
unixsocket /var/run/redis/redis.sock
unixsocketperm 770
sudo systemctl restart redis

vi /etc/webapps/gitlab/gitlab.yml
host: gitlab.org

hexdump -v -n 64 -e '1/1 "%02x"' /dev/urandom > /etc/webapps/gitlab/secret
chmod 640 /etc/webapps/gitlab/secret

hexdump -v -n 64 -e '1/1 "%02x"' /dev/urandom > /etc/webapps/gitlab-shell/secret
chmod 640 /etc/webapps/gitlab-shell/secret

vi /etc/webapps/gitlab/secrets.yml
production:
  secret_key_base: /etc/webapps/gitlab/secret
  db_key_base: /etc/webapps/gitlab-shell/secret

sudo vi /etc/webapps/gitlab/resque.yml
development:
  url: unix:/var/run/redis/redis.sock
test:
  url: unix:/var/run/redis/redis.sock
production:
  url: unix:/var/run/redis/redis.sock

sudo chown -R gitlab /usr/share/webapps/gitlab/tmp
sudo chmod -R u+rwX /usr/share/webapps/gitlab/tmp

sudo systemctl start gitlab-gitaly.service

cd /usr/share/webapps/gitlab
# sudo -u gitlab $(cat environment | xargs) bundle exec rake gitlab:setup
sudo -u gitlab DISABLE_DATABASE_ENVIRONMENT_CHECK=1 $(cat environment | xargs) bundle exec rake gitlab:setup


sudo -u gitlab $(cat environment | xargs) bundle exec rake gitlab:env:info

sudo systemctl start gitlab-workhorse.service

sudo -u gitlab $(cat environment | xargs) bundle exec rake gitlab:check
# Init script exists? ... no :) this is ok for we using systemd


# sudo certbot certonly -d *.didida.top
sudo certbot certonly -d didida.top

openssl genpkey -algorithm RSA -out server.key -pkeyopt rsa_keygen_bits:2048
openssl req -new -key ca.key -out ca.csr
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt

sudo systemctl start nginx redis postgresql gitlab-gitaly gitlab-workhorse gitlab.target

sudo systemctl status nginx redis postgresql gitlab-gitaly gitlab-workhorse

sudo systemctl stop nginx redis postgresql gitlab-gitaly gitlab-workhorse


```

ref: https://wiki.archlinux.org/index.php?title=GitLab&oldid=647720#Configuration



#### Docker Install

```bash
docker run --name gitlab \
	-p 127.0.0.1:22:22 \
	-p 127.0.0.1:443:443 \
	-p 127.0.0.1:80:80 \
	-d gitlab/gitlab-ce

export DATA=<ABSOLUTE PATH>
docker run --name gitlab \
	-p 127.0.0.1:22:22 \
	-p 127.0.0.1:443:443 \
	-p 127.0.0.1:80:80 \
	--hostname gitlab.org \
	-v $DATA/gitlab/etc:/etc/gitlab \
	-v $DATA/gitlab/opt:/var/opt/gitlab \
	-v $DATA/gitlab/log:/var/log/gitlab \
	-d gitlab/gitlab-ce

# browser http://localhost:80

```



### Add Kubernetes Cluster

login -> admin root -> Admin Area -> Kubernetes -> Integrade with a cluster certificate -> Connect existiing cluseter

```bash
# https://gitlab.org/admin/application_settings/network#js-outbound-settings
# check Allow requests to the local network from web hooks and services

# get kubernetes api url, should be exactly the same
kubectl cluster-info | grep 'Kubernetes master' | awk '/http/ {print $NF}'

# get kubernetes secrets
kubectl get secrets

kubectl get secret <SECRETS NAME> -o jsonpath="{['data']['ca\.crt']}" | base64 --decode

vi gitlab-admin-service-account.yaml

apiVersion: v1
kind: ServiceAccount
metadata:
  name: gitlab
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: gitlab-admin
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: gitlab
    namespace: kube-system


kubectl apply -f gitlab-admin-service-account.yaml

# get kubernetes token
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep gitlab | awk '{print $1}')

# add gitlab prometheus PersistentVolume, no need to claim, just add
kubectl apply -f gitlab-pv-volume.yaml


```

ref: https://docs.gitlab.com/ee/user/project/clusters/add_remove_clusters.html

ref: https://kubernetes.io/zh/docs/tasks/configure-pod-container/configure-persistent-volume-storage/



