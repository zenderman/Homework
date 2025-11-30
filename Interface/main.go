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

	andrey := User{"Andrey"}

	fmt.Println(SayHello(andrey))

	fmt.Println(User{"Кирилл"}.Greet())
	fmt.Println(Robot{"T-800"}.Greet())

}
