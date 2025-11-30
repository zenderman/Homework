package main

import "fmt"

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Привет, " + u.Name
}

type Robot struct {
	Model string
}

func (r Robot) Greet() string {
	return "Привет, я робот " + r.Model
}

func SayHello(g Greeter) string {
	return g.Greet()
}

func main() {

	var SomeGreeter Greeter

	SomeGreeter = User{"Andrey"}

	fmt.Println(SayHello(SomeGreeter))

	fmt.Println(User{"Кирилл"}.Greet())
	fmt.Println(Robot{"T-800"}.Greet())

}
