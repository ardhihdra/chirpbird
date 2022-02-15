package db

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/joho/godotenv"
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

var (
	listIndex = []string{IdxEvents, IdxMessages, IdxGroups, IdxUsers, IdxSessions, IdxUserConnections}
	ES_HOSTS  string
	ES_PORT   string
	esAddress string
	Elastic   *elasticsearch.Client
)

type RawReturn struct {
	ID     string      `json:"_id"`
	Source interface{} `json:"_source"`
}

func loadEnv() {
	godotenv.Load()
	ES_HOSTS = os.Getenv("ES_HOSTS")
	// ES_PORT = os.Getenv("ES_PORT")
	if ES_HOSTS == "" {
		ES_HOSTS = "http://localhost:9200"
	}
	// if ES_PORT == "" {
	// 	ES_PORT = "9200"
	// }
}

func Migrate() string {
	query := InitQuery(strings.Split(esAddress, ",")[0])
	cmd, err := exec.Command("/bin/sh", "-c", query).Output()
	if err != nil {
		fmt.Println("error migration", err, cmd)
	}
	output := string(cmd)
	fmt.Println(output)
	return output
}

func initIndex() {
	fmt.Println("Initiating Index")
	for i := range listIndex {
		res, err := Elastic.Cat.Indices(
			Elastic.Cat.Indices.WithIndex(listIndex[i]),
		)
		if err != nil || res.IsError() {
			res, err := Elastic.Index(
				listIndex[i],
				strings.NewReader("{}"),
			)
			if err != nil {
				log.Fatalf("ERROR: %s", err)
			}
			defer res.Body.Close()
			if res.IsError() {
				PrintErrorResponse(res)
			}
			fmt.Println("Index created: ", listIndex[i])
		} else {
			fmt.Println("Index already exist: ", listIndex[i])
		}
	}
}

func init() {
	loadEnv()
	esClient, err := createElasticsearchClient()
	if err != nil {
		fmt.Println("failed to create es client", err)
	}
	Elastic = esClient

	var r map[string]interface{}
	fmt.Println("ES Client created", Elastic, ES_HOSTS, ES_PORT)
	// 1. Get cluster info
	clusterInfo, err := Elastic.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println("cluster info", clusterInfo)
	defer clusterInfo.Body.Close()
	if err := json.NewDecoder(clusterInfo.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	// Migrate()
	initIndex()
}

func createElasticsearchClient() (*elasticsearch.Client, error) {
	adresses := fmt.Sprintf("%v", ES_HOSTS)
	flag.StringVar(&esAddress, "es-addresses", adresses, "elasticsearch addresses")
	cfg := elasticsearch.Config{
		Addresses: strings.Split(esAddress, ","),
		Username:  "",
		Password:  "",
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
	searchRes, err := Elastic.Search(
		Elastic.Search.WithIndex(index),
		Elastic.Search.WithBody(&buf),
		Elastic.Search.WithPretty(),
	)
	if err != nil {
		fmt.Printf("Error searching: %s\n", err)
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

func MoreLikeCondition(cond map[string]interface{}) map[string]interface{} {
	query := map[string]interface{}{
		"from": 0, "size": 1000,
		"query": map[string]interface{}{
			"more_like_this": cond,
		},
	}
	return query
}

func MatchCondition(cond map[string]interface{}) map[string]interface{} {
	query := map[string]interface{}{
		"from": 0, "size": 1000,
		"query": map[string]interface{}{
			"match": cond,
		},
	}
	return query
}

func MustMatch(cond []map[string]interface{}) map[string]interface{} {
	query := map[string]interface{}{
		"from": 0, "size": 1000,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": cond,
			},
		},
	}
	return query
}

func QueryString(cond map[string]interface{}) map[string]interface{} {
	query := map[string]interface{}{
		"from": 0, "size": 1000,
		"query": map[string]interface{}{
			"query_string": cond,
		},
	}
	return query
}

func MatchFilterCondition(cond map[string]interface{}, filter map[string]interface{}) map[string]interface{} {
	query := map[string]interface{}{
		"from": 0, "size": 1000,
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
