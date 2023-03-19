package main

import "fmt"

func first_add(x int, y int) int {
	return x + y
}

func second_add(x int, y int, z int) int {
	return x + y + z
}

func main() {
	fmt.Println(first_add(42, 13))
	fmt.Println(second_add(42, 13, 33))
}
