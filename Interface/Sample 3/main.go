package main

import "fmt"

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return fmt.Sprintf("Hello %s", u.Name)
}

func PrintGreeter(g Greeter) {
	u, ok := g.(User)
	if ok {
		fmt.Println(u.Greet())
	}
}

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

	var u User = User{Name: "Ivan"}
	PrintGreeter(u)
}
