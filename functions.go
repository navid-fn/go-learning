package main

import "fmt"

// define function
func increment(x int) int {
	x += 1
	return x
}

// function which return multiple value
func getPoint(x_coordinate, y_coordinate int) (x, y int) {
	// it initialize x and y to Zero
	x = x_coordinate
	y = y_coordinate

	// return // explicit return value
	return x, y  // its better to be like this
}

func main() {
	// check increment function
	x := 3
	x = increment(x)
	fmt.Printf("Your x is %d \n", x)

	// check getPoint function
	x1, _ := getPoint(x, x) // use _ to ignore value for return function
	fmt.Printf("Our first coordinate is %d \n", x1)

}
