package main

import "fmt"

func main() {
	nilaiA := 20
	nilaiB := 10
	nilaiC := nilaiA + nilaiB
	nilaiD := nilaiA - nilaiB
	nilaiE := nilaiA / nilaiB
	nilaiF := nilaiA * nilaiB
	nilaiG := nilaiA % nilaiB

	nilaiB++
	fmt.Println(nilaiB)
	fmt.Println("===================")
	fmt.Println(nilaiC)
	fmt.Println(nilaiD)
	fmt.Println(nilaiE)
	fmt.Println(nilaiF)
	fmt.Println(nilaiG)
}
