package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"fmt"
	"math/rand"
	"time"
	"github.com/mongodb/mongo-go-driver/bson"
)


const MONGO_SERVER = "mongodb://localhost:27017"
const MONGO_Database = "examples"
const MONGO_Collection = "test"

type Person struct {
	Name      string
	Phone     string
	createdAt time.Time
	created   time.Time
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

func createIndex(coll *mongo.Collection) {
	//
	// Create Index
	//

	// expireAt := bson.NewDocument(bson.EC.Int32("expireAt", 1))
	// expireAfterSeconds := bson.NewDocument(bson.EC.Int64("expireAfterSeconds", 0))

	expireAt := bson.NewDocument(bson.EC.Int32("createdAt", 1))
	expireAfterSeconds := bson.NewDocument(bson.EC.Int32("expireAfterSeconds", 10))
	// expireAfterSeconds := bson.NewDocument(bson.EC.Int64("expireAfterSeconds", 6 * 7 * 24 * 60 * 60))

	indexModel := mongo.IndexModel{
		Keys:    expireAt,
		Options: expireAfterSeconds,
	}

	indexResult, err := coll.Indexes().CreateOne(context.Background(), indexModel)

	if err != nil {
		panic(err)
	}

	fmt.Println("indexResult=" + indexResult)
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

	db.RunCommand(
		context.Background(),
		bson.NewDocument(bson.EC.Int32("dropDatabase", 1)),
	)

	coll := db.Collection(MONGO_Collection)

	createIndex(coll)

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

	// str := "2018-06-05T16:01:26.371Z"
	// removeTime, err := time.Parse(time.RFC3339, str)
	//
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// pups := Person{RandStringRunes(10), "iPhone", removeTime}

	// bson.EC.DateTime()
	// pups := Person{RandStringRunes(10), "iPhone", time.Now(), time.Now()}
	// pups := Person{RandStringRunes(10), "iPhone", removeTime, removeTime}

	pups := bson.NewDocument(
		bson.EC.String("Name", RandStringRunes(10)),
		bson.EC.String("Item", "canvas"),
		bson.EC.Time("created", time.Now().UTC()),
		bson.EC.Time("createdAt", time.Now().UTC()),
		// bson.EC.DateTime("created", time.Now().Unix()),
		// bson.EC.DateTime("createdAt", time.Now().Unix()),
	)

	fmt.Println("pups")
	fmt.Println(pups)

	result, err := coll.InsertOne(context.Background(), pups)

	if err != nil {
		panic(err)
	}

	fmt.Println("result")
	fmt.Println(result.InsertedID)

	client.Disconnect(context.Background())
}
