package Ado

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	endpoint = "https://dev.azure.com"
)

type provider struct {
	patToken     string
	organization string
	poolName     string
}

func NewProvider(patToken string, organization string, poolName string) *provider {
	return &provider{
		patToken:     patToken,
		organization: organization,
		poolName:     poolName,
	}
}

func (provider *provider) GetPoolByName(poolName string) (id int64) {
	poolJson := provider.getAdoPools()
	pools := parseAdoPools(poolJson)

	var poolId int64
	for _, pool := range pools.Value {
		if pool.Name == poolName {
			return pool.ID
		}
	}

	return poolId
}

func (provider *provider) getAdoPools() []byte {

	resp, error := provider.callAdoUrl("_apis/distributedtask/pools?api-version=6.0")

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

	parseAdoBuilds(bodyBytes)

	return bodyBytes
}

func parseAdoPools(adoPoolJson []byte) AdoPools {
	var adoPools AdoPools
	err := json.Unmarshal(adoPoolJson, &adoPools)

	if err != nil {
		log.Print(err)
	}

	return adoPools
}

func (provider *provider) GetBuildsInQueue(poolId int64) int {

	buildJson := provider.getLatestBuilds(poolId)
	builds := parseAdoBuilds(buildJson)
	return getJobsInQueue(builds)
}

func (provider *provider) getLatestBuilds(poolId int64) []byte {
	url := fmt.Sprintf("_settings/agentpools?poolId=%d&__rt=fps&__ver=2", poolId)
	resp, error := provider.callAdoUrl(url)

	if error != nil {
		log.Print(error)
	}

	if resp.StatusCode != 200 {
		log.Printf("There was an error calling the %s : %d", url, resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print("There was an error reading the body: ", err)
	}

	return bodyBytes
}

func parseAdoBuilds(buildJson []byte) AdoBuilds {

	var adoBuilds AdoBuilds
	err := json.Unmarshal(buildJson, &adoBuilds)

	if err != nil {
		log.Print(err)
	}

	return adoBuilds
}

func (provider *provider) callAdoUrl(path string) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", endpoint, provider.organization, path), nil)

	req.SetBasicAuth("", provider.patToken)
	return client.Do(req)
}

func getJobsInQueue(builds AdoBuilds) int {
	var queuedJobsCount int
	for _, job := range builds.Fps.DataProviders.Data.Ms_vss_build_web_agent_jobs_data_provider.Jobs {
		if job.FinishTime == "" && job.ReceiveTime == "" {
			queuedJobsCount++
		}
	}

	return queuedJobsCount
}
