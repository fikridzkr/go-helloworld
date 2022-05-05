package main

import "fmt"

func main() {
	const TITLE = "PROFILE"
	var a, b, c, d int = 1, 3, 5, 7
	var age int32
	firstName := "Fikri"

	age = 18

	fmt.Println(a, b, c, d)
	fmt.Println(TITLE)
	fmt.Println(age)
	fmt.Println((string(firstName[0])), "F")
	fmt.Println("Hello, World!")

	array1 := [...]int{4, 5, 6, 7, 8}
	array2 := [2]int{2, 3}

	fmt.Println(array2)
	fmt.Println(cap(array1))

}
