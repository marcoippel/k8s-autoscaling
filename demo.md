### Install kind
https://kind.sigs.k8s.io/

### Create KIND Cluster
```bash
cat <<EOF | kind create cluster --name=autoscaling --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
EOF
```

### Create the namespaces
```bash
kubectl create ns monitoring && \
kubectl create ns buildagentpool-team-a
```

### Install the prometheus components 
```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts && \
helm repo update && \

helm upgrade --install prometheusadapter prometheus-community/prometheus-adapter -f ./resources/prometheusAdapter-values.yaml -n monitoring && \
helm upgrade --install prometheus prometheus-community/prometheus -f ./resources/prometheus-values.yaml -n monitoring
```

### Install the custom components
```bash
kubectl apply -k ./resources/buildQueueMonitor/ && \
kubectl apply -k ./resources/buildagent/
```

### Get the api result for the custom metric
```bash
kubectl get --raw /apis/external.metrics.k8s.io/v1beta1/namespaces/buildagentpool-team-a/jobs_in_queue_total | jq
```

### Query that is executed against Prometheus
```bash
ceil(sum(avg_over_time(jobs_in_queue_total{kubernetes_namespace="buildagentpool-team-a"}[2m])))
```

### Enable the query logging in Prometheus
```bash
query_log_file: /prometheus/query.log
```