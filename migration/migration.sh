# USERS
curl -X PUT "localhost:9200/users?pretty" -H 'Content-Type: application/json' -d'
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
}' && curl -X PUT "localhost:9200/sessions?pretty" -H 'Content-Type: application/json' -d'
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
}' && curl -X PUT "localhost:9200/events?pretty" -H 'Content-Type: application/json' -d'
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
}' && curl -X PUT "localhost:9200/groups?pretty" -H 'Content-Type: application/json' -d'
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
}' && curl -X PUT "localhost:9200/messages?pretty" -H 'Content-Type: application/json' -d'
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

# "composed_of": ["component_template1", "runtime_component_template"], 
#    "aliases": {
#      "mydata": { }
#    }


# &&
# Use Kibana for alternative
# PUT /users
# {
#   "settings": {
#     "number_of_shards": 1
#   },
#   "mappings": {
#     "_source": {
#       "enabled": true
#     },
#     "properties": {
#       "id": {
#         "type": "keyword"
#       },
#       "username": {
#         "type": "keyword"
#       },
#       "email": {
#         "type": "keyword"
#       },
#       "password": {
#         "type": "keyword"
#       },
#       "profile": {
#         "type": "keyword"
#       },
#       "interests": {
#         "type": "keyword"
#       },
#       "created_at": {
#         "type": "date"
#       },
#       "updated_at": {
#         "type": "date"
#       }
#     }
#   }
# }

# curl -X GET "localhost:9200/users/_search?&pretty" -H 'Content-Type: application/json' -d'{ "query": { "match_all": {} } }'
# curl -X DELETE "localhost:9200/users/_doc/1?routing=shard-1&pretty"

# curl -X DELETE "localhost:9200/twitter?pretty"

