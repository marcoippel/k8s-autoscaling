using Microsoft.AspNetCore.Mvc;
using Prometheus;

namespace BuildQueueMonitoring.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class FakeQueueController : ControllerBase
    {
        private static readonly Gauge NumberOfJobsInQueue = Metrics.CreateGauge("jobs_in_queue_total", "The total of jobs in the build queue");
        
        [HttpGet]
        public Gauge Get()
        {
            double numberOfJobs = 0;
            double.TryParse(Request.Query["number"], out numberOfJobs);
            NumberOfJobsInQueue.Set(numberOfJobs);

            return NumberOfJobsInQueue;
        }
    }
}
