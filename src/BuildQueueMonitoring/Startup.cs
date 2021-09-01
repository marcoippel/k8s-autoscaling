using System;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text;
using BuildQueueMonitoring.Middleware;
using BuildQueueMonitoring.Services;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Prometheus;

namespace BuildQueueMonitoring
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            var pat = Environment.GetEnvironmentVariable("ADO_PAT");
            var organization = Environment.GetEnvironmentVariable("ADO_ORGANIZATION");
            string token = Convert.ToBase64String(Encoding.ASCII.GetBytes(string.Format($"{string.Empty}:{pat}")));

            services.AddControllers();

            
            services.AddHttpClient("Ado", c =>
                {
                    c.BaseAddress = new Uri($"{Environment.GetEnvironmentVariable("ADO_URL")}");
                    c.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Basic", token);
                });

            services.AddSingleton<IAdoService>(s => new AdoService(s.GetRequiredService<IHttpClientFactory>(), organization));
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseHttpsRedirection();

            app.UseRouting();
            app.UseHttpMetrics();

            app.UseAuthorization();

            app.UseEndpoints(endpoints =>
            {
                app.UsePrometheusMiddleware();
                endpoints.MapMetrics();
                endpoints.MapControllers();
            });

        }
    }
}
