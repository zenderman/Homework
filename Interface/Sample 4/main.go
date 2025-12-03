package main

import "fmt"

type User struct {
	Name string
}

func (u *User) Greet() string {
	return fmt.Sprintf("Hello %s", u.Name)
}

type Dog struct {
	Name string
}

func (d *Dog) Greet() string {
	return fmt.Sprintf("Woof! My name is %s", d.Name)
}

type Admin struct {
	User
	Dog
	banned     bool
	banMessage string
}

func (a *Admin) GetBanMessage() string {
	return a.banMessage
}

func (a *Admin) Ban(message string) {
	a.banned = true
	a.banMessage = fmt.Sprintf("user %s has been banned for %s", a.User.Name, message)
}

func NewAdmin(user User) *Admin {
	return &Admin{user, Dog{"bojik"}, false, ""}
}

func (a *Admin) Greet() string {
	if a.banned {
		return a.banMessage
	}
	return a.User.Greet()
}

type Greeter interface {
	Greet() string
}

func main() {
	user1 := NewAdmin(User{"John"})
	user1.Ban("for spamming")
	fmt.Println(user1.Greet())
}
