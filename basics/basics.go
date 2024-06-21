package main

import (
	"fmt"
)

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
	var Var1 int = 20       // or use var1 := 20
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

	// List //
	var myList [3]int
	fmt.Println(myList)

	// can compate two list
	var list1 = [3]int{1, 2, 3}
	var list2 = [3]int{1, 2, 3}

	fmt.Println(list1 == list2)

	// simulate multi-dimentionals
	var multiList = [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(multiList)

	// len
	fmt.Println(len(list1))

	// Slices //
	var slice1 = []int{1, 2, 3, 4, 5, 6}
	var nilSlice []int
	fmt.Println(slice1, "\n", nilSlice)

	fmt.Println(nilSlice == nil)
	slice1 = append(slice1, 1)
	fmt.Println(slice1)

	// create slice with known size
	knowsSizeSlice := make([]int, 5)
	fmt.Println(knowsSizeSlice)
	knowsSizeSlice = append(knowsSizeSlice, 10)
	fmt.Println(knowsSizeSlice)

	setCapacity := make([]int, 0, 10)
	setCapacity = append(setCapacity, 1, 2, 3, 4)

	fmt.Println(setCapacity)

	// clear slice
	var clearSlice = []string{"hello", "bye", "good bye"}
	clear(clearSlice)
	fmt.Println(clearSlice, len(clearSlice))

	// slices of slice
	// there is no copy so change one of them, change another one
	sliceOne := slice1[1:4]
	sliceOne[0] = 100

	fmt.Println(sliceOne, slice1)
	sliceOne = append(sliceOne, 1300)

	fmt.Println(sliceOne, slice1)
	x := slice1[:2:2]
	y := slice1[2:4:4]
	fmt.Println(x, y)

	// copy
	newSlic1 := make([]int, 4)
	copy(newSlic1, slice1[:4])
	newSlic1 = append(newSlic1, 3)
	fmt.Println(newSlic1)
	fmt.Println(slice1)
	// check byte convention
	var z string = "Hello, 🌞"
	var zByte = []byte(z)
	var zRune = []rune(z)

	fmt.Println(zByte, "\n", zRune)

	// MAP //
	translation := map[int]string{
		1: "succesful",
		2: "failed",
	}
	totalCheck := map[string]int{
		"succesful": 1,
		"failed":    2,
	}
	fmt.Println(translation)
	fmt.Println("empty value: ", translation[9])
	fmt.Println(totalCheck)
	fmt.Println("empty value: ", totalCheck["check"])

	// check if value exists in map
	v, ok := totalCheck["sad"]
	fmt.Println(v, ok)

	// simulate set in GO with maps
	values := []int{1, 2, 10, 10, 20, 1, 8}
	valsSet := map[int]bool{}

	for _, i := range values {
		valsSet[i] = true
	}
	if valsSet[18] {
		fmt.Println("value exists")
	} else {
		fmt.Println("value not exists")
	}

	// STRUCT //
	type car struct {
		name        string
		price       float64
		yearCreated string
	}

}
