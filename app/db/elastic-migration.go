package db

import (
	"bytes"
	"fmt"
	"text/template"
)

func InitQuery(address string) string {
	data := map[string]interface{}{
		"ES_HOST": address,
		"ES_PORT": ES_PORT,
	}

	var ES_MIGRATION_QUERY = `
	curl -X PUT "{{.ES_HOST}}/users?pretty" -H 'Content-Type: application/json' -d'
	{
	  "settings": {
		"number_of_shards": 1
	  },
	  "mappings": {
		"_source": {
		  "enabled": true
		},
		"properties": {
		  "id": {
			"type": "text"
		  },
		  "username": {
			"type": "text"
		  },
		  "email": {
			"type": "text",
			"analyzer": "whitespace"
		  },
		  "password": {
			"type": "text"
		  },
		  "profile": {
			"type": "text",
			"analyzer": "whitespace"
		  },
		  "interests": {
			"type": "keyword"
		  },
		  "created_at": {
			"type": "date"
		  },
		  "updated_at": {
			"type": "date"
		  }
		}
	  }
	}' && curl -X PUT "{{.ES_HOST}}/sessions?pretty" -H 'Content-Type: application/json' -d'
	{
	  "settings": {
		"number_of_shards": 1
	  },
	  "mappings": {
		"_source": {
		  "enabled": true
		},
		"properties": {
		  "id": {
			"type": "text"
		  },
		  "user_id": {
			"type": "text"
		  },
		  "type": {
			"type": "text"
		  },
		  "device_id": {
			"type": "text"
		  },
		  "platform": {
			"type": "keyword"
		  },
		  "build": {
			"type": "integer"
		  },
		  "name": {
			"type": "text"
		  },
		  "access_token": {
			"type": "text"
		  },
		  "online": {
			"type": "boolean"
		  },
		  "created_at": {
			"type": "date"
		  },
		  "updated_at": {
			"type": "date"
		  },
		  "online_at": {
			"type": "date"
		  },
		  "offline_at": {
			"type": "date"
		  }
		}
	  }
	}' && curl -X PUT "{{.ES_HOST}}/events?pretty" -H 'Content-Type: application/json' -d'
	{
	  "settings": {
		"number_of_shards": 1
	  },
	  "mappings": {
		"_source": {
		  "enabled": true
		},
		"properties": {
		  "id": {
			"type": "keyword"
		  },
		  "type": {
			"type": "integer"
		  },
		  "object_id": {
			"type": "keyword"
		  },
		  "user_ids": {
			"type": "keyword"
		  },
		  "timestamp": {
			"type": "date"
		  }
		}
	  }
	}' && curl -X PUT "{{.ES_HOST}}/groups?pretty" -H 'Content-Type: application/json' -d'
	{
	  "settings": {
		"number_of_shards": 1
	  },
	  "mappings": {
		"_source": {
		  "enabled": true
		},
		"properties": {
		  "id": {
			"type": "keyword"
		  },
		  "name": {
			"type": "text"
		  },
		  "user_id": {
			"type": "text"
		  },
		  "user_ids": {
			"type": "text"
		  },
		  "deleted": {
			"type": "boolean"
		  },
		  "created_at": {
			"type": "date"
		  },
		  "updated_at": {
			"type": "date"
		  }
		}
	  }
	}' && curl -X PUT "{{.ES_HOST}}/messages?pretty" -H 'Content-Type: application/json' -d'
	{
	  "settings": {
		"number_of_shards": 3,
		"number_of_replicas": 2
	  },
	  "mappings": {
		"_source": {
		  "enabled": true
		},
		"properties": {
		  "id": {
			"type": "keyword"
		  },
		  "user_id": {
			"type": "text"
		  },
		  "group_id": {
			"type": "text"
		  },
		  "body": {
			"properties": {
			  "data": { "type": "text" }
			}
		  },
		  "created_at": {
			"type": "date"
		  },
		  "updated_at": {
			"type": "date"
		  }
		}
	  }
	}
	'
	`

	tmpl, err := template.New("ES_MIGRATION_QUERY").Parse(ES_MIGRATION_QUERY)
	buf := &bytes.Buffer{}
	if err = tmpl.Execute(buf, data); err != nil {
		fmt.Println(err)
	}
	s := buf.String()
	fmt.Println("result query", s)
	return s
}
