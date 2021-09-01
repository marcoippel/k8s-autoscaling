using System;

namespace BuildQueueMonitoring.Models
{
    public class JobRequests
    {
        public int count { get; set; }
        public JobRequestValue[] value { get; set; }
    }
}