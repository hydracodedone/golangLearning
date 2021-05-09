package action

import (
	"fmt"
	"log"
	"mongodb_demo/connection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var colletionName string = "basic_query"

func queryProcedure(queryBson *bson.D) {
	db, ctx := connection.GetDB()
	collection := db.Collection(colletionName, options.Collection())
	cursor, err := collection.Find(ctx, queryBson, &options.FindOptions{})
	if err != nil {
		log.Fatalln(err)
	} else {
		var results []bson.D
		if err = cursor.All(ctx, &results); err != nil {
			log.Fatal(err)
		}
		for _, result := range results {
			fmt.Println(result)
		}
	}
}

func queryProcedureWithQueryOption(queryBson *bson.D, option *options.FindOptions) {
	db, ctx := connection.GetDB()
	collection := db.Collection(colletionName, options.Collection())
	cursor, err := collection.Find(ctx, queryBson, option)
	if err != nil {
		log.Fatalln(err)
	} else {
		var results []bson.D
		if err = cursor.All(ctx, &results); err != nil {
			log.Fatal(err)
		}
		for _, result := range results {
			fmt.Println(result)
		}
	}
}
func SearchInsert() {
	docs := []interface{}{
		bson.D{
			bson.E{Key: "nilValue", Value: nil},
			bson.E{Key: "textValue", Value: "this is a test"},
			bson.E{Key: "item", Value: "journal"},
			bson.E{Key: "qty", Value: 25},
			bson.E{Key: "size", Value: bson.D{
				bson.E{Key: "h", Value: 14},
				bson.E{Key: "w", Value: 21},
				bson.E{Key: "uom", Value: "cm"},
			}},
			bson.E{Key: "status", Value: "A"},
			bson.E{Key: "status", Value: "A"},
			bson.E{Key: "dim_cm", Value: bson.A{1, 2, 4, 5, 6}},
			bson.E{Key: "instock", Value: bson.A{
				bson.D{
					bson.E{Key: "warehouse", Value: "A"},
					bson.E{Key: "qty", Value: 40},
				},
				bson.D{
					bson.E{Key: "warehouse", Value: "B"},
					bson.E{Key: "qty", Value: 5},
				},
			}},
		},
		bson.D{
			bson.E{Key: "textValue", Value: "this is a test too"},
			bson.E{Key: "item", Value: "notebook"},
			bson.E{Key: "qty", Value: 50},
			bson.E{Key: "size", Value: bson.D{
				bson.E{Key: "h", Value: 8.5},
				bson.E{Key: "w", Value: 11},
				bson.E{Key: "uom", Value: "in"},
			}},
			bson.E{Key: "status", Value: "A"},
			bson.E{Key: "dim_cm", Value: bson.A{1, 3, 4, 5}},
			bson.E{Key: "instock", Value: bson.A{
				bson.D{
					bson.E{Key: "warehouse", Value: "B"},
					bson.E{Key: "qty", Value: 15},
				},
				bson.D{
					bson.E{Key: "warehouse", Value: "C"},
					bson.E{Key: "qty", Value: 35},
				},
			}},
		},
	}
	db, ctx := connection.GetDB()
	collection := db.Collection(colletionName, options.Collection())
	_, err := collection.InsertMany(ctx, docs)
	if err != nil {
		log.Fatalln(err)
	}
	index := mongo.IndexModel{
		Keys: bson.D{
			bson.E{
				Key:   "textValue",
				Value: "text",
			},
		},
	}
	_, err = collection.Indexes().CreateOne(ctx, index, &options.CreateIndexesOptions{})
	if err != nil {
		log.Fatalln(err)
	}
}
func SearchDrop() {
	db, ctx := connection.GetDB()
	collection := db.Collection(colletionName, options.Collection())
	err := collection.Drop(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func Search1() {
	//查询全部
	queryBson := bson.D{}
	queryProcedure(&queryBson)
}
func Search2() {
	//精确匹配
	queryBson := bson.D{
		bson.E{Key: "status", Value: "A"},
	}
	queryProcedure(&queryBson)
}
func Search3() {
	//in 查询
	queryBson := bson.D{
		bson.E{
			Key: "status",
			Value: bson.D{
				bson.E{
					Key:   "$in",
					Value: bson.A{"A", "D"},
				}},
		},
	}
	queryProcedure(&queryBson)
}

func Search4() {
	//and 查询
	//gt 查询
	queryBson := bson.D{
		bson.E{Key: "status", Value: "A"},
		bson.E{Key: "qty", Value: bson.D{{Key: "$gt", Value: 10}}},
	}
	queryProcedure(&queryBson)
}
func Search5() {
	//or 查询
	queryBson := bson.D{
		bson.E{
			Key: "$or",
			Value: bson.A{
				bson.D{bson.E{Key: "status", Value: "A"}},
				bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$lt", Value: 10}}}},
			},
		},
	}
	queryProcedure(&queryBson)
}

func Search6() {
	//and or 混合查询
	queryBson := bson.D{
		bson.E{Key: "status", Value: "A"},
		{
			Key: "$or",
			Value: bson.A{
				bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$gt", Value: 30}}}},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search7() {
	//嵌入式文档字段查询,通过.方式指定字段
	queryBson := bson.D{
		{
			Key: "size.h",
			Value: bson.D{
				bson.E{
					Key:   "$gt",
					Value: 10,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search8() {
	//嵌入式文档精确查询,要求元素数量和顺序的精确匹配,顺序不一致或者字段缺失会匹配失败
	//不推荐
	queryBson := bson.D{
		{
			Key: "size",
			Value: bson.D{
				bson.E{
					Key:   "h",
					Value: 14,
				},
				bson.E{
					Key:   "w",
					Value: 21,
				},
				bson.E{
					Key:   "uom",
					Value: "cm",
				},
			},
		},
	}
	queryProcedure(&queryBson)
}

func Search9() {
	//列表字段的精确匹配,要求元素数量和顺序的精确匹配,顺序不一致或者字段缺失会匹配失败
	queryBson := bson.D{
		bson.E{
			Key:   "dim_cm",
			Value: bson.A{1, 3, 4},
		},
	}
	queryProcedure(&queryBson)
}
func Search10() {
	//all查询,要求只包含全部的指定元素,顺序不要求
	queryBson := bson.D{
		bson.E{
			Key: "dim_cm",
			Value: bson.D{
				bson.E{
					Key:   "$all",
					Value: bson.A{4, 1},
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search11() {
	//单个元素查询,要求包含指定的元素即可
	queryBson := bson.D{
		bson.E{
			Key:   "dim_cm",
			Value: 4,
		},
	}
	//等价于
	// queryBson = bson.D{
	// 	bson.E{
	// 		Key: "dim_cm",
	// 		Value: bson.D{
	// 			bson.E{
	// 				Key:   "$all",
	// 				Value: bson.A{4},
	// 			},
	// 		},
	// 	},
	// }
	queryProcedure(&queryBson)
}
func Search12() {
	//要求至少存在一个元素满足要求
	queryBson := bson.D{
		bson.E{
			Key: "dim_cm",
			Value: bson.D{
				bson.E{
					Key:   "$gt",
					Value: 3,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search13() {
	//存在满足所有要求的不同元素(同一元素)
	queryBson := bson.D{
		bson.E{
			Key: "dim_cm",
			Value: bson.D{
				bson.E{
					Key:   "$gt",
					Value: 5,
				},
				bson.E{
					Key:   "$gte",
					Value: 6,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search14() {
	//至少存在同一个元素满足所有要求
	queryBson := bson.D{
		bson.E{
			Key: "dim_cm",
			Value: bson.D{
				bson.E{
					Key: "$elemMatch",
					Value: bson.D{
						bson.E{
							Key:   "$gt",
							Value: 5,
						},
						bson.E{
							Key:   "$lte",
							Value: 6,
						},
					},
				},
			},
		},
	}
	queryProcedure(&queryBson)
}

func Search15() {
	//根据索引进行查询
	//index beigin with 0
	queryBson := bson.D{
		bson.E{
			Key:   "dim_cm.0",
			Value: 1,
		},
	}
	queryProcedure(&queryBson)
}
func Search16() {
	//数组长度查询
	queryBson := bson.D{
		bson.E{
			Key: "dim_cm",
			Value: bson.D{
				bson.E{
					Key:   "$size",
					Value: 5,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}

func Search17() {
	//要求列表存在一个元素的精确匹配,要求元素数量和顺序的精确匹配,顺序不一致或者字段缺失会匹配失败
	queryBson := bson.D{
		bson.E{
			Key: "instock",
			Value: bson.D{
				bson.E{
					Key:   "warehouse",
					Value: "A",
				},
				bson.E{
					Key:   "qty",
					Value: 40,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search17_1() {
	//要求存在整个元素的精确匹配,要求元素数量和顺序的精确匹配,顺序不一致或者字段缺失会匹配失败
	queryBson := bson.D{
		bson.E{
			Key: "instock",
			Value: bson.A{
				bson.D{
					bson.E{Key: "warehouse", Value: "B"},
					bson.E{Key: "qty", Value: 15},
				},
				bson.D{
					bson.E{Key: "warehouse", Value: "C"},
					bson.E{Key: "qty", Value: 35},
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search18() {
	// 至少有一个嵌入文档的元素满足要求
	queryBson := bson.D{
		bson.E{
			Key: "instock.qty",
			Value: bson.D{
				bson.E{
					Key:   "$gt",
					Value: 35,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search19() {
	// 某个具体的嵌入文档的元素满足要求
	queryBson := bson.D{
		bson.E{
			Key: "instock.0.qty",
			Value: bson.D{
				bson.E{
					Key:   "$gt",
					Value: 35,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search20() {
	//至少有一个嵌入文档同时满足要求,没有顺序要求
	queryBson := bson.D{
		bson.E{
			Key: "instock",
			Value: bson.D{
				bson.E{
					Key: "$elemMatch",
					Value: bson.D{
						bson.E{Key: "qty", Value: 40},
						bson.E{Key: "warehouse", Value: "A"},
					},
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search21() {
	//存在满足所有要求的不同元素(同一元素)
	queryBson := bson.D{
		bson.E{
			Key: "instock.qty",
			Value: bson.D{
				bson.E{
					Key:   "$lte",
					Value: 35,
				},
				bson.E{
					Key:   "$gte",
					Value: 15,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}

func Search22() {
	queryBson := bson.D{}
	option := options.Find()
	option.SetProjection(
		bson.D{
			bson.E{Key: "_id", Value: 0},
			bson.E{Key: "size.uom", Value: 1},
			bson.E{Key: "dim_cm", Value: 1},
			bson.E{Key: "instock.qty", Value: 1},
		},
	)
	queryProcedureWithQueryOption(&queryBson, option)
}
func Search23() {
	queryBson := bson.D{}
	option := options.Find()
	option.SetProjection(
		bson.D{
			bson.E{Key: "_id", Value: 0},
			//exclude mode
			bson.E{Key: "size.uom", Value: 0},
		},
	)
	queryProcedureWithQueryOption(&queryBson, option)
}
func Search24() {
	// Specify a positive number n to return the first n elements.(left include right not include)
	//Specify a negative number n to return the last n elements.
	queryBson := bson.D{}
	option := options.Find()
	option.SetProjection(
		bson.D{
			bson.E{
				Key: "instock",
				Value: bson.D{
					bson.E{
						Key:   "$slice",
						Value: 1,
					},
				}},
		},
	)
	queryProcedureWithQueryOption(&queryBson, option)
}

func Search26() {
	//查询不包含的字段或者字段对应的值为nil
	queryBson := bson.D{
		bson.E{Key: "nilValue", Value: nil},
	}
	queryProcedure(&queryBson)
}
func Search27() {
	//查询包含的字段且字段对应的值不为nil
	queryBson := bson.D{
		bson.E{
			Key: "nilValue",
			Value: bson.D{
				bson.E{
					Key:   "$ne",
					Value: nil,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search28() {
	//查询包含的字段且字段对应的值为nil
	queryBson := bson.D{
		bson.E{
			Key: "nilValue",
			Value: bson.D{
				bson.E{
					Key:   "$type",
					Value: 10,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}

func Search29() {
	//查询不包含的字段
	queryBson := bson.D{
		bson.E{
			Key: "nilValue",
			Value: bson.D{
				bson.E{
					Key:   "$exists",
					Value: false,
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func Search30() {
	//文本搜索
	//要执行文本搜索查询，您必须在集合上有一个文本索引。一个集合只能有一个文本搜索索引，但是该索引可以覆盖多个字段。
	queryBson := bson.D{
		bson.E{
			Key: "$text",
			Value: bson.D{
				bson.E{
					Key:   "$search",
					Value: "too",
				},
			},
		},
	}
	queryProcedure(&queryBson)
}
func SearchDemo() {
	SearchDrop()
	SearchInsert()
}
