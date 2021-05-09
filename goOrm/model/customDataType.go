package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CustomInfo struct {
	Name string
	Age  string
}

type CustomDataTable struct {
	gorm.Model
	CustomInfo CustomInfo
}

func (cd *CustomDataTable) TableName() string {
	return "custom_data_table"
}

func (ci *CustomInfo) Scan(value interface{}) error {

	if bytesValue, ok := value.([]byte); ok {
		err := json.Unmarshal(bytesValue, ci)
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return fmt.Errorf("info is not []byte")
	}
}

func (ci CustomInfo) Value() (driver.Value, error) {
	fmt.Println(456)
	bytesData, err := json.Marshal(ci)
	if err != nil {
		return "", err
	}
	return string(bytesData), nil
}

func (ci CustomInfo) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	marshalData, err := ci.Value()
	if err != nil {
		marshalData = "{}"
	}
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{marshalData},
	}
}
