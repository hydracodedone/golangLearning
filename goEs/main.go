package main

import (
	"fmt"
	"go_es/connection"
	"go_es/constant"
	"go_es/mapping"

	_ "github.com/olivere/elastic/v7"
)

func main() {
	con := connection.GetConnection()
	fmt.Println(con)
	indexName := "user_info"
	fmt.Println(mapping.CreateUserMapping(con, indexName, constant.USER_MAPPING_TPL))
}
