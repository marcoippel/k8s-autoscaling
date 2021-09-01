using System.Threading.Tasks;
using BuildQueueMonitoring.Models;

namespace BuildQueueMonitoring.Services
{
    public interface IAdoService
    {
        Task<Pools> GetAgentPools();
        Task<JobRequests> GetJobRequests(int poolId);
    }
}