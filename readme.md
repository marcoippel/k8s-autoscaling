## Azure DevOps scalable build agent pool on Kubernetes 

In this article I'm gona describe how to host a scaleable private agent pool for Azure Devops. Most of the time you will use the hosted agents provided by Microsoft. There are some cases that you want to use a private build agent or agents.  

TODO: usecases when you want a private build agent pool. You can install an Azure devops agent on a virtual machine or on an Azure vm scaleset which is supported by default in Azure Devops. 
But if you want te be in control of the scaling and if you want a fast scaling solution you can host your buildagents on a Kubernetes 
cluster. 

By default Kubernetes can scale on CPU and memory usage but there are also options to scale based on an external metric. For example on the count of builds in the queue of a Azure DevOps agent pool.

![Overview](https://raw.githubusercontent.com/marcoippel/k8s-autoscaling/main/Images/K8S%20AutoScaling.png)


## Components

### Buildqueue monitoring
The solution starts with an application called [Buildqueue monitoring](https://github.com/marcoippel/k8s-autoscaling/tree/main/src/BuildQueueMonitoring). This is a web-api written in C# which monitors the build queue of a specific buildagent pool. The count of jobs that are queued are added to a metric that can be scraped by prometheus.

The amount of builds queued in the queue is checked when Prometheus calls the endpoint /metrics I created a middleware extension. 
When prometheus scrapes the metrics the middleware is calling the api of Azure DevOps and gets the amount of build in the queue.
The endpoint of the Azure Devops api is undocumented but it gets all the builds of the specified buildqueue. Builds that have no finishTime and also not receiveTime then it is queued. Builds that have no finishtime but there is a recieve time then the build is running. Builds that have a finish time these builds are finished.

I have installed the [Prometheus-net.AspNetCore](https://www.nuget.org/packages/prometheus-net.AspNetCore/) nuget package which helps you to create Prometheus metrics.

The reason why I have chozen for a middleware extension is that it is simple and effective. Now the Azure Devops Api is called when Prometheus calls the /metrics api. You can create a backgroud worker and cache the value returned from the api. But we just call one Api on Azure Devops I have chozen for a middleware extension.

```csharp
public async Task Invoke(HttpContext httpContext)
{
    if (httpContext.Request.Path.HasValue && httpContext.Request.Path.Value == "/metrics")
    {
        _logger.LogTrace($"{DateTime.Now:o} - Start scraping");

        var poolName = Environment.GetEnvironmentVariable("ADO_POOLNAME");
        
        var queuedJobs = 0;
        try
        {
            // Get all the pools from Azure DevOps
            var pools = await _adoService.GetAgentPools();

            // Get the pool by name.
            var pool = pools.value.Single(x => x.name == poolName);

            // Get all the jobs from the specified pool
            var jobRequests = await _adoService.GetJobRequests(pool.id);
            
            // No finishTime, and also not receiveTime, then it is queued.
            // When there is no finishTime, but there is a receiveTime, it is running.
            // If there is a finishTime, it is finished
            queuedJobs = jobRequests.value.Count(x => string.IsNullOrEmpty(x.finishTime) && string.IsNullOrEmpty(x.receiveTime));
        }
        catch (Exception e)
        {
            queuedJobs = 0;
            _logger.LogError($"{DateTime.Now:o} - {e.Message}");
        }
        finally
        {
            // Set the value for the gauge metric
            var gauge = Metrics.CreateGauge("jobs_in_queue_total", "The amount of jobs in the azuredevops queue");
            gauge.Set(queuedJobs);

            _logger.LogTrace($"{DateTime.Now:o} - End scraping");
        }
    }

    await _next.Invoke(httpContext);
}

```

If prometheus scrapes the endpoint the metric looks like the following. This is the standard format prometheus understands.

```bash
# HELP jobs_in_queue_total The amount of jobs in the azuredevops queue
# TYPE jobs_in_queue_total gauge
jobs_in_queue_total 0
```

// TODO add file path
You can install the buildqueue monitor by applying the resourse files:
- BuildQueueMonitoring.yaml
- BuildQueueMonitoring-service.yaml

You need to install the buildqueue monitor in the same namespace as your buildagents will be installed. This is because this way you can scale multiple agent pools for instance for multiple teams.



### Prometheus
Prometheus is a free software application used for event monitoring and alerting. It records real-time metrics in a time series database built using a HTTP pull model, with flexible queries and real-time alerting. "source: [wikipedia](https://en.wikipedia.org/wiki/Prometheus_(software))"

Prometheus is used to scrape the metrics endpoint which is provided by the buildqueue monitoring application. Prometheus scrapes the endpoint and saves the data. 

To install Prometheus you can install the helm chart 
```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install [RELEASE_NAME] prometheus-community/prometheus -f Prometheus-values.yaml
```

An example of the value file used to install the Prometheus helm chart. 
```yaml
alertmanager:
  enabled: false
kubeStateMetrics:
  enabled: false
nodeExporter:
  enabled: false
pushgateway:
  enabled: false
extraScrapeConfigs: |
  - job_name: 'buildqueuemonitoring'
    metrics_path: /metrics
    static_configs:
      - targets: ['buildqueuemonitoring.buildagentpool-team-a.svc:80']
        labels:
          kubernetes_namespace: 'buildagentpool-team-a'
```
The scrape config adds the buildqueuemonitoring as a target to Prometheus. We add also a label called "kubernetes_namespace" with the value of the namespace where you installed the buildqueuemonitoring application.
