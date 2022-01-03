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

const IdxMessaging string = "messaging"
const (
	TypeEvents          string = "events"
	TypeMessages        string = "messages"
	TypeGroups          string = "groups"
	TypeUsers           string = "users"
	TypeSessions        string = "sessions"
	TypeUserConnections string = "userconnections"
)

var (
	ES_HOST   string
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
	ES_HOST = os.Getenv("ES_HOST")
	ES_PORT = os.Getenv("ES_PORT")
	if ES_HOST == "" {
		ES_HOST = "http://localhost"
	}
	if ES_PORT == "" {
		ES_PORT = "9200"
	}
}

func Migrate() string {
	// absPath, _ := filepath.Abs("./")
	// query := fmt.Sprintf("%s/%s", absPath, "migration/migration.sh")
	/** local migration might be wont work */
	query := InitQuery(strings.Split(esAddress, ",")[0])
	cmd, err := exec.Command("/bin/sh", "-c", query).Output()
	if err != nil {
		fmt.Println("error migration", err, cmd)
	}
	output := string(cmd)
	fmt.Println(output)
	return output
}

func init() {
	loadEnv()
	esClient, err := createElasticsearchClient()
	if err != nil {
		fmt.Println("failed to create es client", err)
	}
	Elastic = esClient

	var r map[string]interface{}
	fmt.Println("ES Client created", Elastic, ES_HOST, ES_PORT)
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

	fmt.Println("Initiating Index")
	Migrate()
}

func createElasticsearchClient() (*elasticsearch.Client, error) {
	adresses := fmt.Sprintf("http://elastic:9200,http://%s:%s", ES_HOST, ES_PORT)
	flag.StringVar(&esAddress, "es-addresses", adresses, "elasticsearch addresses")
	cfg := elasticsearch.Config{
		Addresses: strings.Split(esAddress, ","),
		Username:  "",
		Password:  "",
	}
	client, err := elasticsearch.NewClient(cfg)
	return client, err
}

func ExecuteQuery(query map[string]interface{}, dtype string, b *bytes.Buffer) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding: %s", err)
		return err
	}
	searchRes, err := Elastic.Search(
		Elastic.Search.WithIndex(IdxMessaging),
		Elastic.Search.WithSearchType(dtype),
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
