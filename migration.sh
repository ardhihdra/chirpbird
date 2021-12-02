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
      "type": {
        "type": "int"
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


