package main

import "fmt"

func Sum(numbers []int) int {
	var suma int
	for _, sum := range numbers {
		suma += sum
	}

	return suma
}

func FilterEven(num []int) []int {

	res := make([]int, 0, len(num))

	for _, v := range num {
		if v%2 == 0 {
			res = append(res, v)
		}
	}

	return res
}

func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func DeleteElement(nums []int, num int) []int {
	nums = append(nums[:num], nums[num+1:]...)

	return nums
}

func InsertElement(nums []int, index int, val int) []int {
	nums = append(nums[:index+1], val)

	return nums
}

func main() {

	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(FilterEven([]int{1, 2, 3, 33, 44}))
	n1 := []int{1, 2, 3, 4, 5, 6}
	Reverse(n1)
	fmt.Println(n1)
	fmt.Println(DeleteElement(n1, 2))
	fmt.Println(n1)
	fmt.Println(InsertElement(n1, 2, 55))
}
