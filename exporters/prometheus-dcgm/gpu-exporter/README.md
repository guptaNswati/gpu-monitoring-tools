# gpu-exporter

A simple go http server serving per pod GPU metrics(dcgm-pod.prom) at localhost:9400/gpu/metrics

```sh
# Add a gpu metrics endpoint to prometheus 
$ kubectl create -f prometheus-configmap.yaml 

$ kubectl create -f gpu-exporter-daemonset.yaml

$ curl localhost:9090

# To open in browser: ssh -L 9090:localhost:9090 node@ip_addr
```