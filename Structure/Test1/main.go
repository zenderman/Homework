package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) Len2() int {
	return p.X*p.X + p.Y*p.Y
}

func main() {

	// Тест 1: Point{3,4}
	p1 := Point{3, 4}
	result1 := p1.Len2()
	fmt.Printf("Point{%d, %d}.Len2() = %d\n", p1.X, p1.Y, result1)
	if result1 != 25 {
		panic("want 25")
	}

	// Дополнительные тесты
	tests := []Point{
		{-3, 4}, // 9 + 16 = 25
		{0, 5},  // 0 + 25 = 25
		{6, 0},  // 36 + 0 = 36
	}

	fmt.Println("\nДополнительные тесты:")
	for i, p := range tests {
		result := p.Len2()
		if result != 25 {
			fmt.Printf("Тест feild %d: Point{%d, %d}.Len2() = %d\n", i+1, p.X, p.Y, result)
			panic("want 25")
		}
		fmt.Printf("Тест %d: Point{%d, %d}.Len2() = %d\n", i+1, p.X, p.Y, result)
	}

	fmt.Println("\nВсе тесты пройдены успешно!")

}
