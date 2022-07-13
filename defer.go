package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func Divide(val int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / val
	println(result)
}

func main() {
	Divide(0)
}
