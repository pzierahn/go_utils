package main

import (
	"fmt"
)

func main() {
	// 	SocketServer(3333)

	var mymap = make(map[ string ] string)
	mymap[ "Test1" ] = "asdf"
	mymap[ "Test2" ] = "asdasff"
	mymap[ "Test3" ] = "agasfhdfht"

	fmt.Println(mymap)

	if val, ok := mymap[ "foo" ]; ok {
		fmt.Println(val)
	}

	if str, ok := json2Str(mymap); ok {
		fmt.Println("jsonString: " + str)
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
}
