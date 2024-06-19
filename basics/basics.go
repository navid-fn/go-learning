package main

import "fmt"

func main() {
	fmt.Println("This is Print!")

	// if you dont assign anything to variable, default will be set (which is zero value for each type)
	var NoValue int 
	fmt.Println(NoValue)

	// multiple variable with same type
	var X1, X2 int = 10, 20
	fmt.Println(X1, X2)

	// decalration list
	var (
		myVar          int
		myVar2, myVar3     = "a", "b"
		myVar4             = 3
		myVar5         int = 4
		myVar6             = "string"
	)

	fmt.Println(myVar, myVar2, myVar3, myVar4, myVar5, myVar6)

	// comparison for having value or is true
	// you there is no such a thing as `if value:`
	var Var1 int = 20 // or use var1 := 20
	var Var2 string = "abc" // or use var2 := "abc"

	var IsBiger bool = Var1 > 10  // true
	var IsEmpty bool = Var2 == "" // false
	fmt.Println(IsBiger, IsEmpty)

	// dont use `x := ` outside of function


	// constant variable, immutable thing in go
	const a = 1
	const b = "string"

	// all will be failed:
	// a = a + 1
	// a += 1
	// b = "bye"


	// List
	var myList [3]int
	fmt.Println(myList)
	
}
