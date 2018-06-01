package main

import (
	"fmt"
)

func main() {
	var mymap = make(map[ string ] string)
	mymap[ "Test1" ] = "asdf"
	mymap[ "Test2" ] = "asdasff"
	mymap[ "Test3" ] = "agasfhdfht"

	fmt.Println(mymap)

	if val, ok := mymap[ "foo" ]; ok {
		fmt.Println(val)
	}


	str, err := JsonStr(mymap)

	if err != nil {
		fmt.Println("jsonString: " + str)
	}
}
