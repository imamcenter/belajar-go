package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("error dengan message", message)
	}
	fmt.Println("application selesai")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("APLKASI ERROR")
	}
	fmt.Println("aplikasi berjalan")
}
func main() {
	runApp(false)
	fmt.Println("hello  world")
}
