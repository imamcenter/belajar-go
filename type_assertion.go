package main

import "fmt"

func random() any {
	return []string{"imam", "ahmad"}
}
func main() {
	message := random()
	switch message.(type) {
	case string:
		fmt.Println(message, "is string")
	case int:
		fmt.Println(message, "is int")
	default:
		fmt.Println("unknown")
	}
}
