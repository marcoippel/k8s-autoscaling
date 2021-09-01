using System.Net.Http;
using System.Threading.Tasks;
using BuildQueueMonitoring.Models;
using Newtonsoft.Json;

namespace BuildQueueMonitoring.Services
{
    public class AdoService : IAdoService
    {
        private readonly string _organization;
        private readonly HttpClient _client;

        public AdoService(IHttpClientFactory clientFactory, string organization)
        {
            _organization = organization;
            _client = clientFactory.CreateClient("Ado");
        }

        public async Task<Pools> GetAgentPools()
        {
            string requestUrl = $"/{_organization}/_apis/distributedtask/pools?api-version=6.0";
            HttpRequestMessage request = new HttpRequestMessage(HttpMethod.Get, requestUrl);

            HttpResponseMessage response = _client.SendAsync(request).Result;
            string responseString = await response.Content.ReadAsStringAsync();

            return JsonConvert.DeserializeObject<Pools>(responseString);
        }

        public async Task<JobRequests> GetJobRequests(int poolId)
        {
            // Call the azuredevops api to get the builds from the queue
            // https://developercommunity.visualstudio.com/t/api-to-get-current-number-of-builds-running-and-qu/969466

            string requestUrl = $"/{_organization}/_apis/distributedtask/pools/{poolId}/jobrequests"; 
            HttpRequestMessage request = new HttpRequestMessage(HttpMethod.Get, requestUrl);

            HttpResponseMessage response = await _client.SendAsync(request);
            string responseString = await response.Content.ReadAsStringAsync();

            return JsonConvert.DeserializeObject<JobRequests>(responseString);
        }
    }
}
