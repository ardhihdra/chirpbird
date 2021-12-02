package db

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type indexList struct {
	Events   string
	Messages string
	Groups   string
	Users    string
	Sessions string
}

var IndexList = indexList{
	Events:   "events",
	Messages: "messages",
	Groups:   "groups",
	Users:    "users",
	Sessions: "sessions",
}

var Elastic *elasticsearch.Client

func init() {
	esClient, err := createElasticsearchClient()
	Elastic = esClient

	var r map[string]interface{}
	fmt.Println("ES Client created")
	// 1. Get cluster info
	clusterInfo, err := Elastic.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer clusterInfo.Body.Close()
	// Deserialize the response into a map.
	if err := json.NewDecoder(clusterInfo.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	fmt.Println("Initiating Index")

}

func createElasticsearchClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "",
		Password: "",
	}
	client, err := elasticsearch.NewClient(cfg)
	return client, err
}

func PrintErrorResponse(res *esapi.Response) {
	fmt.Printf("[%s] ", res.Status())
	var e map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		log.Fatalf("ES, Error parsing the response body: %s", err)
	} else {
		// Print the response status and error information.
		log.Fatalf("ES, [%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)
	}
}
