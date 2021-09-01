using System;
using System.Threading.Tasks;
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

        public PrometheusMiddleware(
            RequestDelegate next, 
            ILogger<PrometheusMiddleware> logger)
        {
            
            _next = next;
            _logger = logger;
        }

        public async Task Invoke(HttpContext httpContext)
        {
            if (httpContext.Request.Path.HasValue && httpContext.Request.Path.Value == "/metrics")
            {
                _logger.LogTrace($"{DateTime.Now:o} - Start scraping");
                
                var status = 1;
                try
                {
                    //call the azuredevops api to get the builds from the queue
                }
                catch (Exception e)
                {
                    status = 0;
                    _logger.LogError($"{DateTime.Now:o} - {e.Message}");
                }
                finally
                {
                    var gauge = Metrics.CreateGauge("jobs_in_queue_total", "The amount of jobs in the azuredevops queue");
                    gauge.Set(status);

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
