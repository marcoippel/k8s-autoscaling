using System;
using System.Linq;
using System.Threading.Tasks;
using BuildQueueMonitoring.Services;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Logging;
using Prometheus;

namespace BuildQueueMonitoring.Middleware
{
    public class PrometheusMiddleware
    {
        private readonly RequestDelegate _next;
        private readonly ILogger<PrometheusMiddleware> _logger;
        private readonly IAdoService _adoService;
        
        public PrometheusMiddleware(
            RequestDelegate next,
            ILogger<PrometheusMiddleware> logger, 
            IAdoService adoService)
        {
            _next = next;
            _logger = logger;
            _adoService = adoService;
        }

        public async Task Invoke(HttpContext httpContext)
        {
            if (httpContext.Request.Path.HasValue && httpContext.Request.Path.Value == "/metrics")
            {
                _logger.LogTrace($"{DateTime.Now:o} - Start scraping");

                var poolName = Environment.GetEnvironmentVariable("ADO_POOLNAME");
                
                var queuedJobs = 0;
                try
                {
                    var pools = await _adoService.GetAgentPools();
                    var pool = pools.value.Single(x => x.name == poolName);

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
                    var gauge = Metrics.CreateGauge("jobs_in_queue_total", "The amount of jobs in the azuredevops queue");
                    gauge.Set(queuedJobs);

                    _logger.LogTrace($"{DateTime.Now:o} - End scraping");
                }
            }

            await _next.Invoke(httpContext);
        }
    }

    public static class PrometheusMiddlewareExtensions
    {
        public static IApplicationBuilder UsePrometheusMiddleware(this IApplicationBuilder builder)
        {
            return builder.UseMiddleware<PrometheusMiddleware>();
        }
    }
}
