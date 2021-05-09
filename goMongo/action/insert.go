package action

import (
	"fmt"
	"log"
	"mongodb_demo/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Insert1() {
	var name string = "Hydra"
	var age int = 23
	var data model.Info = model.Info{
		Name:  &name,
		Age:   &age,
		Male:  nil,
		Hobby: []string{"Shooting", "Game"},
		Parents: []model.Parent{
			{
				Name: "Father",
				Age:  50,
			},
			{
				Name: "Mother",
				Age:  51,
			},
		},
	}
	collection, ctx := model.InfoGetCollection()
	result, err := collection.InsertOne(ctx, data, &options.InsertOneOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("insert result %v\n", result.InsertedID)
}
func Insert2() {
	var name1 string = "Hydra"
	var age1 int = 23
	var name2 string = "Hydra2"
	var age2 int = 24
	var male2 bool = true
	var datas []interface{} = []interface{}{
		model.Info{
			Name:  &name1,
			Age:   &age1,
			Male:  nil,
			Hobby: []string{"Shooting", "Game"},
			Parents: []model.Parent{
				{
					Name: "Father",
					Age:  50,
				},
				{
					Name: "Mother",
					Age:  51,
				},
			},
		},
		model.Info{
			Name:  &name2,
			Age:   &age2,
			Male:  &male2,
			Hobby: []string{"Shooting", "Game"},
			Parents: []model.Parent{
				{
					Name: "Father",
					Age:  50,
				},
				{
					Name: "Mother",
					Age:  51,
				},
			},
		},
	}
	collection, ctx := model.InfoGetCollection()
	result, err := collection.InsertMany(ctx, datas, &options.InsertManyOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("insert result %v\n", result)
}
func Insert3() {
	/*
		Adds two documents using insertOne.
		Updates a document using updateOne.
		Deletes a document using deleteOne.
		Replaces a document using replaceOne.
	*/
	var name1 string = "Hydra"
	var age1 int = 23
	var name2 string = "Hydra2"
	var age2 int = 24
	var male2 bool = true
	var writeInfo []mongo.WriteModel = []mongo.WriteModel{
		&mongo.InsertOneModel{
			Document: model.Info{
				Name:  &name1,
				Age:   &age1,
				Male:  nil,
				Hobby: []string{"Shooting", "Game"},
				Parents: []model.Parent{
					{
						Name: "Father",
						Age:  50,
					},
					{
						Name: "Mother",
						Age:  51,
					},
				},
			},
		},
		&mongo.InsertOneModel{
			Document: model.Info{
				Name:  &name2,
				Age:   &age2,
				Male:  &male2,
				Hobby: []string{"Shooting", "Game"},
				Parents: []model.Parent{
					{
						Name: "Father",
						Age:  50,
					},
					{
						Name: "Mother",
						Age:  51,
					},
				},
			},
		},
	}
	ordered := false
	collection, ctx := model.InfoGetCollection()
	result, err := collection.BulkWrite(ctx, writeInfo, &options.BulkWriteOptions{Ordered: &ordered})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("insert result %+v\n", result)
}

func InsetDemo() {
	model.InfoGetCollectionDrop()
	model.InfoGetCollectionDrop()
	model.InfoCreateCollection()
	Insert3()
}
