package main

import "fmt"

func main() {

	fmt.Println("1. Сумма элементов массива")
	mas := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Сумма элементов %v = %v\n", mas, summa(mas))

	fmt.Println("\n2. Умножение всех элементов слайса")
	numbers1 := []int{2, 4, 6, 8, 10}
	fmt.Printf("Исходный слайс: %v\n", numbers1)

	rez := multiply(numbers1, 3)
	fmt.Println("multiply() без копирования ", numbers1, "; резултат = ", rez)

	rez2 := multiplyCopy(numbers1, 3)
	fmt.Println("multiplyCopy() с копированием ", numbers1, "; резултат = ", rez2)

	fmt.Println("\n3. Обрезка и копирование")
	numbers2 := make([]int, 0, 20)
	numbers2 = append(numbers2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Исходные данные: %v\n", numbers2)
	fmt.Printf("Длина: %d, Емкость: %d\n", len(numbers2), cap(numbers2))
	part := numbers2[3:8]
	fmt.Printf("part = numbers2[3:8]: %v\n", part)
	fmt.Printf("Длина part: %d, Емкость part: %d\n", len(part), cap(part))
	copyPart := make([]int, len(part))
	copy(copyPart, part)
	fmt.Printf("part (копия): %v\n", copyPart)
	fmt.Printf("Длина copyPart: %d, Емкость copyPart: %d\n", len(copyPart), cap(copyPart))
	fmt.Printf("Адрес первого элемента part: %p\n", &part[0])
	fmt.Printf("Адрес первого элемента copyPart: %p\n", &copyPart[0])
	fmt.Printf("Адрес первого элемента numbers: %p\n", &numbers2[0])
	fmt.Printf("Изменяем part[0] с %d на 100\n", part[0])
	part[0] = 100
	fmt.Printf("После изменения part: %v\n", part)
	fmt.Printf("copyPart (не изменился): %v\n", copyPart)
	fmt.Printf("numbers2 (изменился!): %v\n", numbers2)
	fmt.Printf("Изменяем copyPart[0] с %d на 200\n", copyPart[0])
	copyPart[0] = 200

	fmt.Printf("part: %v\n", part)
	fmt.Printf("copyPart: %v\n", copyPart)
	fmt.Printf("numbers: %v\n", numbers2)

	fmt.Println("\nЗадание 4. Матрица 3×3 и диагонали\n")
	fmt.Println("Не осилил\n")

	fmt.Println("\nЗадание 5. Удаление дубликатов из слайса\n")
	// Тестовые примеры
	testCases := [][]int{
		{1, 2, 2, 3, 1, 4, 3},
		{1, 1, 1, 1, 1},
		{1, 2, 3, 4, 5},
		{},
		{5, 3, 2, 3, 5, 1},
	}

	for i, nums := range testCases {
		result := unique(nums)
		fmt.Printf("Пример %d: %v -> %v\n", i+1, nums, result)
	}

}

func summa(mas [10]int) int {
	masSum := 0
	for i := 0; i < len(mas); i++ {
		masSum += mas[i]
	}
	return masSum
}

func multiplyCopy(num []int, factor int) []int {

	cop := make([]int, len(num), len(num))
	copy(cop, num)
	num = cop

	for i := 0; i < len(num); i++ {
		num[i] = num[i] * factor
	}
	return num
}

func multiply(num []int, factor int) []int {

	for i := 0; i < len(num); i++ {
		num[i] = num[i] * factor
	}

	return num
}

func unique(nums []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0)

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}
