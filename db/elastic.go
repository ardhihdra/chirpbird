package db

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

var DB *elasticsearch

func createElasticsearchClient() (*elasticsearch.Client, error) {
	var r map[string]interface{}
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "",
		Password: "",
	}
	DB, err := elasticsearch.NewClient(cfg)

	fmt.Println("ES Client created")
	// 1. Get cluster info
	clusterInfo, err := DB.Info()
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
	return DB, err
}

func printErrorResponse(res *esapi.Response) {
	fmt.Printf("[%s] ", res.Status())
	var e map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	} else {
		// Print the response status and error information.
		log.Fatalf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)
	}
}
