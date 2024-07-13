package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000

func main() {
	fmt.Println("Hello Golang!")

	var msg string = "Hello "
	fmt.Println(msg + "World!")

	fmt.Printf("%d %d", numberInt, numberInt2) // printf("%s",variable) in c language

	numberFloat := 25.4
	fmt.Println(numberFloat)

	fmt.Println(numberInt + numberInt2)
	fmt.Println(float64(numberInt) + numberFloat)

	fmt.Println("My money = ", numberInt)
}
