package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type metric struct {
	Help  string
	Type  string
	Value int
}

func main() {
	http.HandleFunc("/metrics", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the build id based on the name of the pool
	adoPoolsJson := getAdoPools()
	pools := parseAdoPools(adoPoolsJson)
	poolName := os.Getenv("POOLNAME")
	poolId, err := getPoolByName(pools, poolName)

	if err != nil {
		log.Print(err)
	}

	// Get the builds
	buildJson := getBuilds(poolId)

	// // Parse the json
	adoBuilds := parseAdoBuilds(buildJson)

	// // Filter on build in the queue
	queuedJobs := filterOnJobsInQueue(adoBuilds)

	gauge := metric{
		Help:  "# HELP our_company_blob_storage_ops_queued Number of blob storage operations waiting to be processed.",
		Type:  "# TYPE our_company_blob_storage_ops_queued gauge",
		Value: queuedJobs,
	}

	fmt.Fprintf(w, "%s\r\n%s\r\n%d", gauge.Help, gauge.Type, gauge.Value)
}

func filterOnJobsInQueue(builds AdoBuilds) int {
	var queuedJobsCount int
	for _, job := range builds.Fps.DataProviders.Data.Ms_vss_build_web_agent_jobs_data_provider.Jobs {
		if job.FinishTime == "" && job.ReceiveTime == "" {
			queuedJobsCount++
		}
	}

	return queuedJobsCount
}

func parseAdoBuilds(buildJson []byte) AdoBuilds {
	var adoBuilds AdoBuilds
	err := json.Unmarshal(buildJson, &adoBuilds)

	if err != nil {
		log.Print(err)
	}

	return adoBuilds
}

func parseAdoPools(adoPoolJson []byte) AdoPools {
	var adoPools AdoPools
	err := json.Unmarshal(adoPoolJson, &adoPools)

	if err != nil {
		log.Print(err)
	}

	return adoPools
}

func getBuilds(poolId int64) []byte {
	resp, error := callAdoUrl(fmt.Sprintf("https://dev.azure.com/Marcoippel/_settings/agentpools?poolId=%d&__rt=fps&__ver=2", poolId))

	if error != nil {
		log.Print(error)
	}

	if resp.StatusCode != 200 {
		log.Print("There was an error", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print("There was an error", err)
	}
	return bodyBytes
}

func getAdoPools() []byte {
	resp, error := callAdoUrl(fmt.Sprintf("https://dev.azure.com/%s/_apis/distributedtask/pools?api-version=6.0", os.Getenv("ADOORGANISATION")))

	if error != nil {
		log.Print(error)
	}

	if resp.StatusCode != 200 {
		log.Print("There was an error", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print("There was an error", err)
	}
	return bodyBytes
}

func callAdoUrl(url string) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("", os.Getenv("ADOPAT"))
	return client.Do(req)
}

func getPoolByName(pools AdoPools, poolName string) (poolId int64, err error) {
	for _, pool := range pools.Value {
		if pool.Name == poolName {
			return pool.ID, nil
		}
	}

	return 0, fmt.Errorf("no pool was found with name: %s", poolName)
}
