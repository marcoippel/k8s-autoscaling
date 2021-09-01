using System;

namespace BuildQueueMonitoring.Models
{
    public class PoolValue
    {
        public DateTime createdOn { get; set; }
        public bool autoProvision { get; set; }
        public bool autoUpdate { get; set; }
        public bool autoSize { get; set; }
        public int? targetSize { get; set; }
        public int? agentCloudId { get; set; }
        public Createdby createdBy { get; set; }
        public Owner owner { get; set; }
        public int id { get; set; }
        public string scope { get; set; }
        public string name { get; set; }
        public bool isHosted { get; set; }
        public string poolType { get; set; }
        public int size { get; set; }
        public bool isLegacy { get; set; }
        public string options { get; set; }
    }
}