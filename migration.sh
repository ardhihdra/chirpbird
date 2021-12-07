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
        "type": "keyword"
      },
      "username": {
        "type": "keyword"
      },
      "email": {
        "type": "keyword"
      },
      "password": {
        "type": "keyword"
      },
      "profile": {
        "type": "keyword"
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
}
'
# Use Kibana for alternative
PUT /users
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
      "username": {
        "type": "keyword"
      },
      "email": {
        "type": "keyword"
      },
      "password": {
        "type": "keyword"
      },
      "profile": {
        "type": "keyword"
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
}


# SESSIONS
curl -X PUT "localhost:9200/sessions?pretty" -H 'Content-Type: application/json' -d'
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
      "user_id": {
        "type": "keyword"
      },
      "type": {
        "type": "keyword"
      },
      "device_id": {
        "type": "keyword"
      },
      "platform": {
        "type": "keyword"
      },
      "build": {
        "type": "integer"
      },
      "name": {
        "type": "keyword"
      },
      "access_token": {
        "type": "keyword"
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
}
'

# EVENTS
curl -X PUT "localhost:9200/events?pretty" -H 'Content-Type: application/json' -d'
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
}
'

# GROUPS
curl -X PUT "localhost:9200/groups?pretty" -H 'Content-Type: application/json' -d'
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
        "type": "keyword"
      },
      "user_id": {
        "type": "keyword"
      },
      "user_ids": {
        "type": "keyword"
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
}
'

# "composed_of": ["component_template1", "runtime_component_template"], 
#    "aliases": {
#      "mydata": { }
#    }


# curl -X GET "localhost:9200/users/_search?&pretty" -H 'Content-Type: application/json' -d'{ "query": { "match_all": {} } }'
# curl -X DELETE "localhost:9200/users/_doc/1?routing=shard-1&pretty"

# curl -X DELETE "localhost:9200/twitter?pretty"

