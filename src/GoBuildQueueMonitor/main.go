package main

import (
	"buildqueuemonitor/src/GoBuildQueueMonitor/Ado"
	"fmt"
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

	log.Print("Starting webserver on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	adoOrganisation := getEnvVar("ADOORGANISATION")
	poolName := getEnvVar("POOLNAME")
	adoPat := getEnvVar("ADOPAT")

	if adoOrganisation == "" || poolName == "" || adoPat == "" {
		w.WriteHeader(500)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	provider := Ado.NewProvider(adoPat, adoOrganisation, poolName)
	poolId := provider.GetPoolByName(poolName)

	jobsInQueue := provider.GetBuildsInQueue(poolId)

	gauge := metric{
		Help:  "# HELP our_company_blob_storage_ops_queued Number of blob storage operations waiting to be processed.",
		Type:  "# TYPE our_company_blob_storage_ops_queued gauge",
		Value: jobsInQueue,
	}

	fmt.Fprintf(w, "%s\r\n%s\r\n%d", gauge.Help, gauge.Type, gauge.Value)
}

func getEnvVar(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Printf("Environment variable %s not found", name)
	}
	return value
}
