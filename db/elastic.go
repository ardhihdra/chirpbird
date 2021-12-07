package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/tidwall/gjson"
)

const (
	IdxEvents          string = "events"
	IdxMessages        string = "messages"
	IdxGroups          string = "groups"
	IdxUsers           string = "users"
	IdxSessions        string = "sessions"
	IdxUserConnections string = "userconnections"
)

var Elastic *elasticsearch.Client

type RawReturn struct {
	ID     string      `json:"_id"`
	Source interface{} `json:"_source"`
}

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

func ExecuteQuery(query map[string]interface{}, index string, b *bytes.Buffer) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding: %s", err)
		return err
	}
	// return u, db.DB.Users.Find(bson.M{"username": strings.ToLower(username)}).One(&u)
	searchRes, err := Elastic.Search(
		Elastic.Search.WithIndex(index),
		Elastic.Search.WithBody(&buf),
		Elastic.Search.WithPretty(),
	)
	if err != nil {
		fmt.Printf("Error searching: %s\n", err)
		// os.Exit(2)
		return err
	}
	defer searchRes.Body.Close()
	if searchRes.IsError() {
		PrintErrorResponse(searchRes)
	}

	// parse with gjson
	b.ReadFrom(searchRes.Body)
	return nil
}

func FindOne(query map[string]interface{}, index string) ([]gjson.Result, error) {
	var b bytes.Buffer
	ExecuteQuery(query, index, &b)
	values := gjson.GetManyBytes(b.Bytes(), "hits.hits.0._id", "hits.hits.0._source")
	return values, nil
}

func FindAll(query map[string]interface{}, index string) ([]gjson.Result, error) {
	var b bytes.Buffer
	ExecuteQuery(query, index, &b)
	values := gjson.GetManyBytes(b.Bytes(), "hits.hits.0._id", "hits.hits")
	return values, nil
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

func MatchCondition(cond map[string]interface{}) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": cond,
		},
	}
	return query
}

func MatchFilterCondition(cond map[string]interface{}, filter map[string]interface{}) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": cond,
				},
				"filter": map[string]interface{}{
					"range": filter,
				},
			},
		},
	}
	return query
}
