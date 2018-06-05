package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"fmt"
	"math/rand"
	"time"
)

const MONGO_SERVER = "mongodb://localhost:27017"
const MONGO_Database = "examples"
const MONGO_Collection = "test"

type Person struct {
	Name string
	Phone string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func mongoSample() {
	// pups := Person{"Bob", "iPhone"}
	//
	// data, err := bson.Marshal(pups)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// fmt.Println("data")
	// fmt.Println(string(data))

	// os.Exit(0)

	client, err := mongo.Connect(context.Background(), MONGO_SERVER, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	db := client.Database(MONGO_Database)
	coll := db.Collection(MONGO_Collection)

	// json := bson.NewDocument(
	// 	bson.EC.String("item", "canvas"),
	// 	bson.EC.Int32("qty", 100),
	// 	bson.EC.ArrayFromElements("tags",
	// 		bson.VC.String("cotton"),
	// 	),
	// 	bson.EC.SubDocumentFromElements("size",
	// 		bson.EC.Int32("h", 28),
	// 		bson.EC.Double("w", 35.5),
	// 		bson.EC.String("uom", "cm"),
	// 	),
	// )

	// result, err := coll.InsertOne(context.Background(), json)

	pups := Person{RandStringRunes(10), "iPhone"}
	result, err := coll.InsertOne(context.Background(), pups)

	fmt.Println("result")
	fmt.Println(result)
}
