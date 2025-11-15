package main

import "fmt"

type Point struct{ X, Y int }

// Конструктор возвращает указатель на Point
func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}

// Move изменяет состояние структуры через указатель
func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Вспомогательный метод для красивого вывода
func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
	// Тест из задания
	fmt.Println("=== Основной тест ===")
	p := NewPoint(1, 1)
	fmt.Printf("Создан: %s\n", p)

	p.Move(2, -1)
	fmt.Printf("После Move(2, -1): %s\n", p)

	if p.X != 3 || p.Y != 0 {
		panic("bad move")
	}
	fmt.Println("✓ Основной тест пройден")

	// Дополнительные тесты
	fmt.Println("\n=== Дополнительные тесты ===")

	// Тест 1: Несколько последовательных перемещений
	p1 := NewPoint(0, 0)
	fmt.Printf("Начало: %s\n", p1)
	p1.Move(5, 3)
	fmt.Printf("После Move(5, 3): %s\n", p1)
	p1.Move(-2, 1)
	fmt.Printf("После Move(-2, 1): %s\n", p1)

	// Тест 2: Отрицательные координаты
	p2 := NewPoint(10, 10)
	fmt.Printf("\nНачало: %s\n", p2)
	p2.Move(-15, -5)
	fmt.Printf("После Move(-15, -5): %s\n", p2)

	// Тест 3: Нулевые перемещения
	p3 := NewPoint(7, 8)
	fmt.Printf("\nНачало: %s\n", p3)
	p3.Move(0, 0)
	fmt.Printf("После Move(0, 0): %s\n", p3)

	// Тест 4: Проверка, что разные экземпляры независимы
	fmt.Println("\n=== Тест независимости экземпляров ===")
	pointA := NewPoint(1, 2)
	pointB := NewPoint(1, 2)

	pointA.Move(3, 4)
	pointB.Move(5, 6)

	fmt.Printf("Point A после Move(3, 4): %s\n", pointA)
	fmt.Printf("Point B после Move(5, 6): %s\n", pointB)

	fmt.Println("\n✓ Все тесты пройдены успешно!")
}
