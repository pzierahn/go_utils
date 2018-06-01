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

	if str, ok := json2Str(mymap); ok {
		fmt.Println("jsonString: " + str)
	}

	primes := []int{2, 3, 5, 7, 11, 13}
	strs := []string{"Hallo", "Blaasdf", "dsfasd", "Bla"}
	fmt.Println(primes)
	fmt.Println(containsInt(primes, 3))
	fmt.Println(containsString(strs, "Bla"))
}
