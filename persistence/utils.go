package persistence

import (
	"log"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type bindingVariables map[string]interface{}

var arangoClient driver.Client
var arangoDatabase driver.Database
var arangoCollections = make(map[string]driver.Collection)
var arangoGraph driver.Graph
var arangoEdgeCollections = make(map[string]driver.Collection)

func getArangoClient() (driver.Client, error) {
	if arangoClient == nil {
		log.Println("Initialize ArangoDB client")
		endpoint := "http://localhost:8529"

		connectionConfig := http.ConnectionConfig{
			Endpoints: []string{endpoint},
		}

		connection, err := http.NewConnection(connectionConfig)
		if err != nil {
			return nil, err
		}

		clientConfig := driver.ClientConfig{
			Connection:     connection,
			Authentication: driver.BasicAuthentication("root", "root"),
		}

		arangoClient, err = driver.NewClient(clientConfig)
		if err != nil {
			return nil, err
		}
	}

	return arangoClient, nil
}

func getArangoDatabase() (driver.Database, error) {
	if arangoDatabase == nil {
		log.Println("Initialize ArangoDB database")
		client, err := getArangoClient()
		if err != nil {
			return nil, err
		}

		database, err := client.Database(nil, "my-cv")
		if err != nil {
			return nil, err
		}

		arangoDatabase = database
	}

	return arangoDatabase, nil
}

func getArangoCollection(name string) (driver.Collection, error) {
	if arangoCollections[name] == nil {
		log.Println("Initialize ArangoDB collection:", name)
		database, err := getArangoDatabase()
		if err != nil {
			return nil, err
		}

		collection, err := database.Collection(nil, name)
		if err != nil {
			return nil, err
		}
		arangoCollections[name] = collection
	}

	return arangoCollections[name], nil
}

func getArangoGraph() (driver.Graph, error) {
	if arangoGraph == nil {
		log.Println("Initialize ArangoDB graph")
		database, err := getArangoDatabase()
		if err != nil {
			return nil, err
		}

		graph, err := database.Graph(nil, "my-cv")
		if err != nil {
			return nil, err
		}

		arangoGraph = graph
	}

	return arangoGraph, nil
}

func getArangoEdgeCollection(name string) (driver.Collection, error) {
	if arangoEdgeCollections[name] == nil {
		log.Println("Initialize ArangoDB edge collection:", name)
		graph, err := getArangoGraph()
		if err != nil {
			return nil, err
		}

		edgeCollection, _, err := graph.EdgeCollection(nil, name)
		arangoEdgeCollections[name] = edgeCollection
	}

	return arangoEdgeCollections[name], nil
}
