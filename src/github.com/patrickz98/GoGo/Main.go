package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const LOGTAG_MAIN = "Main"

type JsonStruc struct {
	Test1 string
	Test2 int
	Test3 []int
	Test4 []string
	Test5 struct {
		name string
		age int
	}
}

func main() {
	fmt.Println("Starting....")

	startWebsocket()

	// parts, _ := sssa.Create(3, 9, RandStringRunes(1024 * 4))
	//
	// for inx, part := range parts {
	// 	fmt.Println("inx=" + strconv.Itoa(inx) + " part=" + part)
	// }
	//
	// msg, _ := sssa.Combine(parts[:3])
	//
	// fmt.Println("msg=" + msg)

	//mongoSample()
	//MongoSearch()

	os.Exit(0)

	StartWebServer()
	//curl()

	SocketServer(3333)

	var mymap = make(map[ string ] string)
	mymap[ "Test1" ] = "asdf"
	mymap[ "Test2" ] = "asdasff"
	mymap[ "Test3" ] = "agasfhdfht"

	fmt.Println(mymap)

	if val, ok := mymap[ "foo" ]; ok {
		fmt.Println(val)
	}

	if str, ok := json2Str(mymap); ok {
		fmt.Println(LOGTAG_MAIN, "-->", "jsonString=" + str)
	}

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println(containsInt(primes, 3))

	strs := []string{"Hallo", "Blaasdf", "dsfasd", "Bla"}
	fmt.Println(strs)
	fmt.Println(containsString(strs, "Bla"))

	dynamticArray := []int{1, 2, 3, 4}
	dynamticArray = append(dynamticArray, 5)
	fmt.Println(dynamticArray)

	jsonTest := JsonStruc{}
	jsonTest.Test1 = "asdfa"
	jsonTest.Test2 = 123
	jsonTest.Test3 = dynamticArray
	// jsonTest.Test4 = strs

	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := jsonfiy(jsonTest)
	if _, err := file.Write(bytes); err != nil {
		fmt.Println(err)
		return
	}

	file.Close()

	fmt.Println("Done")

	inputFile, _ := ioutil.ReadFile("output.json")
	fmt.Println(LOGTAG_MAIN, "-->", "read=" + string(inputFile))

	var jsonObj JsonStruc

	if err := json.Unmarshal(inputFile, &jsonObj); err != nil {
		fmt.Println(err)
	}

	fmt.Println("blabla =", jsonObj)
	fmt.Println(jsonObj.Test3)

	if jsonObj.Test4 == nil {
		fmt.Println("jsonObj.Test4=null")
	}
}
