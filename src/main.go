package main

import "fmt"

var mymap = make(map[ string ] string)

func main() {
	mymap[ "Test1" ] = "asdf"
	mymap[ "Test2" ] = "asdasff"
	mymap[ "Test3" ] = "agasfhdfht"

	fmt.Println(mymap)
}
