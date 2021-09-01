namespace BuildQueueMonitoring.Models
{
    public class JobRequestValue
    {
        public int requestId { get; set; }
        public string queueTime { get; set; }
        public string assignTime { get; set; }
        public string receiveTime { get; set; }
        public string finishTime { get; set; }
        public string result { get; set; }
        public string serviceOwner { get; set; }
        public string hostId { get; set; }
        public string scopeId { get; set; }
        public string planType { get; set; }
        public string planId { get; set; }
        public string jobId { get; set; }
        public string[] demands { get; set; }
        public Reservedagent reservedAgent { get; set; }
        public Definition definition { get; set; }
        public Owner owner { get; set; }
        public Data data { get; set; }
        public int poolId { get; set; }
        public Agentspecification agentSpecification { get; set; }
        public string orchestrationId { get; set; }
        public bool matchesAllAgentsInPool { get; set; }
        public int priority { get; set; }
    }
}