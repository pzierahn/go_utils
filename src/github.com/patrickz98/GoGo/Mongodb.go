package main

import (
	"fmt"
	"time"
	"context"
	"math/rand"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

const MONGO_SERVER = "mongodb://localhost:27017"
const MONGO_Database = "examples"
const MONGO_Collection = "test"

type Person struct {
	_id     objectid.ObjectID
	Name    string
	Bla     string
	Created time.Time
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

	expireAt := bson.NewDocument(bson.EC.Int32("Created", 1))
	expireAfterSeconds := bson.NewDocument(bson.EC.Int32("expireAfterSeconds", 3600))
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

	// db.RunCommand(
	// 	context.Background(),
	// 	bson.NewDocument(bson.EC.Int32("dropDatabase", 1)),
	// )

	coll := db.Collection(MONGO_Collection)

	createIndex(coll)

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

	id := objectid.New()
	fmt.Println("id=" + id.String())

	// bla := Person{RandStringRunes(10), "iPhone", time.Now()}
	pups := Person{
		_id:     id,
		Name:    RandStringRunes(10),
		Bla:     "patrick",
		Created: time.Now(),
	}

	fmt.Println("pups=" + simpleJson2String(pups))

	// pups := bson.NewDocument(
	// 	bson.EC.String("Name", RandStringRunes(8)),
	// 	// bson.EC.String("UUID", uuid()),
	// 	bson.EC.String("bla", "patrick"),
	// 	bson.EC.Time("created", time.Now()),
	// 	bson.EC.ObjectID("_id", id),
	// )

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

func MongoSearch() {
	client, err := mongo.Connect(context.Background(), MONGO_SERVER, nil)

	if err != nil {
		panic(err)
		return
	}

	db := client.Database(MONGO_Database)

	coll := db.Collection(MONGO_Collection)

	id, _ := objectid.FromHex("5b16bd57f9a71e5b6084541b")
	searchQuery := bson.NewDocument(bson.EC.ObjectID("_id", id))

	doc := bson.NewDocument()
	// per := Person{}
	// var per Person
	result := coll.FindOne(context.Background(), searchQuery)
	result.Decode(doc)
	// result.Decode(per)

	// var per Person

	// convert m to s
	// bsonBytes, _ := doc.MarshalBSON()
	// bson.Unmarshal(bsonBytes, &per)

	// fmt.Println("adsfasd --> " + doc.Lookup("_id").ObjectID().String())


	// fmt.Println(per)
	// fmt.Println("pretty=" + simpleJson2String(per))
	fmt.Println("doc=" + doc.ToExtJSON(true))
	fmt.Println("doc=" + doc.ToExtJSON(false))

	iter := doc.Iterator()
	for iter.Next() {
		elem := iter.Element()
		fmt.Println("Key=" + elem.Key())

		if elem.Key() == "_id" {
			fmt.Println("Value=" + elem.Value().ObjectID().String())
		}
	}

	// searchQuery := bson.NewDocument(bson.EC.String("bla", "patrick"))
	// cursur, err := coll.Find(context.Background(), searchQuery)
	//
	// if err != nil {
	// 	panic(err)
	// 	return
	// }
	//
	// doc := bson.NewDocument()
	// for cursur.Next(context.Background()) {
	// 	doc.Reset()
	// 	cursur.Decode(doc)
	//
	// 	fmt.Println("Fund: ")
	// 	fmt.Println(doc)
	// }
}
