package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string) *Person {
	return &Person{
		Name: name,
		Age:  0,
	}
}

// ИЗМЕНЯЕТ возраст
func (p *Person) Birthday() {
	p.Age++
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func NewBook(title, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

type Address struct {
	City   string
	Street string
	Home   int
}

type User struct {
	Name    string
	Address Address
}

// получения полного адреса
func FullAddress(u User) string {
	return fmt.Sprintf("Город: %s, Ул. %s, Дом: %d", u.Address.City, u.Address.Street, u.Address.Home)
}

func main() {

	p1 := Person{
		Name: "Иван",
		Age:  25,
	}
	fmt.Println("До: ", p1)
	p1.Birthday() // + 1 год
	fmt.Println("+ 1 год ", p1)

	p2 := NewPerson("Мария")
	fmt.Println("\nДо:", *p2)
	p2.Birthday()
	fmt.Println("+1 год ", *p2)

	fmt.Println("\n=== === ===")

	book1 := NewBook("Война и мир", "Толстой", 1225)
	fmt.Println("Гнига1: ", *book1)

	book2 := NewBook("Мастер и Маргарита", "Булгаков", 480)
	fmt.Println("Гнига2: ", *book2)

	book3 := new(Book)
	book3.Title = "Преступление и наказание"
	book3.Author = "Достоевский"
	book3.Pages = 671
	fmt.Println("Гнига3: ", *book3)

	fmt.Println("\n=== === ===")

	user1 := User{
		Name: "Иван Иванов",
		Address: Address{
			City:   "Алматы",
			Street: "Абая",
			Home:   10,
		},
	}

	fmt.Println("Пользователь: ", user1.Name)
	fmt.Println("Только адрес: ", user1.Address)
	fmt.Printf("Полный адрес: %s\n", FullAddress(user1))
}
