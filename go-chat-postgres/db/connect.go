package db

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/eaigner/jet"
	"github.com/lib/pq"
)

type vcapServices struct {
	DB []connectionStruct `json:"elephantsql"`
}

type connectionStruct struct {
	Credentials credentialsStruct `json:"credentials"`
}

type credentialsStruct struct {
	Uri string `json:"uri"`
}

// Create a connection to the Postgres database
func NewClient() (*jet.Db, error) {
	var uri string
	var db *jet.Db

	vcap_services := parseVcapServices()
	if len(vcap_services.DB) != 0 {
		uri = vcap_services.DB[0].Credentials.Uri
	} else {
		// Get the Postgress DB from an environment variable. This is usefull for local development
		var err error

		uri, err = parsePostgresEnvironment()
		if err != nil {
			return db, err
		}
	}

	pgUrl, err := pq.ParseURL(uri)
	if err != nil {
		log.Fatal("Error parsing env var from ELEPHANTSQL_URL: " + err.Error())
		return db, err
	}

	db, err = jet.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal("Error connecting to DB: " + err.Error())
		return db, err
	}

	return db, nil
}

// Pasrse the Postgres the url from the VCAP_SERVICES environment variable bound to the application
func parseVcapServices() vcapServices {
	envVar := os.Getenv("VCAP_SERVICES")
	if envVar == "" {
		log.Println("No VCAP_SERVICES environemnt variable has been set")
		return vcapServices{}
	}

	vs := vcapServices{}
	err := json.Unmarshal([]byte(envVar), &vs)
	if err != nil {
		log.Fatal("Error parsing env var from VCAP_SERVICES: " + err.Error())
	}

	return vs
}

// Parse the Postgress url from the environment variable "POSTGRES_URL"
func parsePostgresEnvironment() (string, error) {
	url := os.Getenv("POSTGRES_URL")
	if url == "" {
		return url, errors.New("Error: No POSTGRES_URL environemnt variable has been set.")
	}

	return url, nil
}
