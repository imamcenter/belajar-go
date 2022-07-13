package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	//TODO implement me
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	//TODO implement me
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	//TODO implement me
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{Name: "imam", Age: 20},
		{Name: "ahmad", Age: 12},
		{Name: "fahrezi", Age: 50},
		{Name: "shaqil", Age: 30},
	}
	sort.Sort(UserSlice(users))

	fmt.Println(users)

}
