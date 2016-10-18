package db

import (
	"encoding/json"
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

func NewClient() *jet.Db {
	var uri string
	vcap_services := parseVcapServices()

	if len(vcap_services.DB) != 0 {
		uri = vcap_services.DB[0].Credentials.Uri
	} else {
		uri = "postgres://gmbiqhhk:c96VTB4YWKhy6wewZAYT4Vu8jzU9OdxE@jumbo.db.elephantsql.com:5432/gmbiqhhk"
	}

	pgUrl, err := pq.ParseURL(uri)
	if err != nil {
		log.Fatal("Error parsing env var from ELEPHANTSQL_URL: " + err.Error())
	}

	db, err := jet.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal("Error connecting to DB: " + err.Error())
	}

	return db
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
