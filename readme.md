## Introduction

In this article I'm gona describe how to host a scaleable private agent pool for Azure Devops. Most of the time you will use the hosted agents provided by Microsoft. There are some cases that you want to use a private build agent or agents.  TODO: usecases when you want a private build agent pool. You can install an Azure devops agent on a virtual machine or on an Azure vm scaleset which is supported by default in Azure Devops. But if you want te be in control of the scaling and if you want a fast scaling solution you can host your buildagents on a Kubernetes 
cluster. By default Kubernetes can scale on CPU and memory usage but there are also options to scale based on an external metric. For example on the count of builds in the queue of a Azure DevOps agent pool.

![Overview](https://raw.githubusercontent.com/marcoippel/k8s-autoscaling/main/Images/K8S%20AutoScaling.png)


### Components

#### Buildqueue monitoring
The solution starts with an application called [Buildqueue monitoring](https://github.com/marcoippel/k8s-autoscaling/tree/main/src/BuildQueueMonitoring). This is a web-api written in C# which monitors the build queue of a specific buildagent pool. The count of jobs that are queued are added to a metric that can be scraped by prometheus.

The amount of builds queued in the queue is checked when Prometheus calls the endpoint /metrics I created a middleware extension. 
When prometheus scrapes the metrics the middleware is calling the api of Azure DevOps and gets the amount of build in the queue.
The endpoint of the Azure Devops api is undocumented but it gets all the builds of the specified buildqueue. Builds that have no finishTime and also not receiveTime then it is queued. Builds that have no finishtime but there is a recieve time then the build is running. Builds that have a finish time these builds are finished.

I have installed the [Prometheus-net.AspNetCore](https://www.nuget.org/packages/prometheus-net.AspNetCore/) nuget package which helps you to create Prometheus metrics.

I add the value of the queued builds to a gauge metric which looks like this
//TODO Image add metrics

![Overview](https://raw.githubusercontent.com/marcoippel/k8s-autoscaling/main/Images/middleware-extension.png)

//Todo How to install the buildagent monitoring
