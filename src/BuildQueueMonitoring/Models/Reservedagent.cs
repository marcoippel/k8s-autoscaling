namespace BuildQueueMonitoring.Models
{
    public class Reservedagent
    {
        public _Links _links { get; set; }
        public int id { get; set; }
        public string name { get; set; }
        public string version { get; set; }
        public string osDescription { get; set; }
        public bool enabled { get; set; }
        public string status { get; set; }
        public string provisioningState { get; set; }
        public string accessPoint { get; set; }
    }
}