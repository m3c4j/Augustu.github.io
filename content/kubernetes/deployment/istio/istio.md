

```bash
# https://istio.io/latest/docs/examples/bookinfo/

istioctl install --set profile=demo -d manifests -y

kubectl label namespace default istio-injection=enabled


kubectl apply -f bookinfo.yaml

kubectl get services
kubectl get pods

kubectl exec "$(kubectl get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}')" -c ratings -- curl -sS productpage:9080/productpage | grep -o "<title>.*</title>"


kubectl apply -f bookinfo-gateway.yaml

kubectl get gateway

kubectl get svc istio-ingressgateway -n istio-system

export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
echo $INGRESS_PORT

export INGRESS_HOST=192.168.0.10

export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT

curl -s "http://${GATEWAY_URL}/productpage" | grep -o "<title>.*</title>"


```

