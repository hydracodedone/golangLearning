package mapping

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type UserInfoMapping struct {
	client  *elastic.Client
	index   string
	mapping string
}

func (uM *UserInfoMapping) mappingInit() error {
	ctx := context.Background()
	exists, err := uM.client.IndexExists(uM.index).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		_, err := uM.client.CreateIndex(uM.index).Body(uM.mapping).Do(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
func CreateUserMapping(esClient *elastic.Client, indexName string, mappingString string) error {
	if esClient != nil {
		uMapping := &UserInfoMapping{
			client:  esClient,
			index:   indexName,
			mapping: mappingString,
		}
		return uMapping.mappingInit()
	} else {
		return fmt.Errorf("unable to create mapping,esClient is nil")
	}
}
