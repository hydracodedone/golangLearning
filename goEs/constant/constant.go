package constant

var ES_URL string = "http://localhost"
var ES_PORT int = 9200
var USER_MAPPING_TPL = `{
	"mappings":{
		"properties":{
			"id":{
				"type":"long"
            },
            "name":{
                "type":"object",
                "properties":{
                    "firstName":{
                        "type":"keyword"
                    },
                    "lastName":{
                        "type":"keyword"
                    }
                }
            },
            "email":{
                "type":"keyword",
                "index":false
            },
            "personal_info":{
                "type":"text",
                "analyzer":"ik_smart" 
            },
            "created_at":{
                "type":"date"
            },
            "updated_at":{
                "type":"date"
            },
            "deleted_at":{
                "type":"date"
            }
    	}
    }
}`
