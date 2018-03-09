package persistence

import (
	"log"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type bindingVariables map[string]interface{}

var arangoClient driver.Client
var arangoDatabase driver.Database
var arangoCollections map[string]driver.Collection

func getArangoClient() *driver.Client {
	if arangoClient == nil {
		log.Println("Initializing ArangoDB client")
		endpoint := "http://localhost:8529"

		connectionConfig := http.ConnectionConfig{
			Endpoints: []string{endpoint},
		}

		connection, err := http.NewConnection(connectionConfig)
		if err != nil {
			log.Fatalln("ERROR: can't create connection:", err)
		}

		clientConfig := driver.ClientConfig{
			Connection:     connection,
			Authentication: driver.BasicAuthentication("root", "root"),
		}

		arangoClient, err = driver.NewClient(clientConfig)
		if err != nil {
			log.Fatalln("ERROR: can't create ArangoDB client:", err)
		}
	}

	return &arangoClient
}

func getArangoDatabase() *driver.Database {
	if arangoDatabase == nil {
		log.Println("Initializing ArangoDB database")
		client := getArangoClient()
		var err error
		arangoDatabase, err = (*client).Database(nil, "my-cv")
		if err != nil {
			log.Fatalln("ERROR: can't get database:", err)
		}
	}

	return &arangoDatabase
}

func getArangoCollection(name string) *driver.Collection {
	if arangoCollections[name] == nil {
		log.Println("Initialize ArangoDB collection:", name)
		database := getArangoDatabase()
		collection, err := (*database).Collection(nil, name)
		if err != nil {
			log.Fatalln("ERROR: can't get collection", name, err)
		}
		arangoCollections[name] = collection
	}

	col := arangoCollections[name]
	return &col
}
