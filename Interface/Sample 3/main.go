package main

import "fmt"

func main() {
	var someValue interface{} = "123"

	v, ok := someValue.(int)
	if ok {
		fmt.Println("это int", v)
	}

	s, ok := someValue.(string)
	if ok {
		fmt.Println("это строка", s)
	}
}
