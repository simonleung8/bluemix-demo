package db

import (
	"encoding/json"
	"log"
	"os"

	couchdb "github.com/fjl/go-couchdb"
)

type vcapServices struct {
	DB []connectionStruct `json:"cloudantNoSQLDB"`
}

type connectionStruct struct {
	Credentials credentialsStruct `json:"credentials"`
}

type credentialsStruct struct {
	Url string `json:"url"`
}

func NewClient() *couchdb.Client {
	var url string
	vcap_services := parseVcapServices()

	if len(vcap_services.DB) != 0 {
		url = vcap_services.DB[0].Credentials.Url
	} else {
		log.Fatal("DB Credential not found in env.")
	}

	cloudant, err := couchdb.NewClient(url, nil)
	if err != nil {
		log.Fatal("Error connecting to Cloudant: " + err.Error())
	}

	return cloudant
}

func parseVcapServices() vcapServices {
	envVar := os.Getenv("VCAP_SERVICES")
	if envVar == "" {
		return vcapServices{}
	}

	vs := vcapServices{}
	err := json.Unmarshal([]byte(envVar), &vs)
	if err != nil {
		log.Print("Error parsing env var from VCAP_SERVICES: " + err.Error())
	}

	return vs
}
