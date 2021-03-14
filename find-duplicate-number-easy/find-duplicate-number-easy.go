package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 3}
	n := 4

	fmt.Println(sol1(arr))
	fmt.Println(sol2(arr))
	fmt.Println(sol3(arr, n))
	fmt.Println(sol4(arr))
}

func sol1(arr []int) int {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		for j := i + 1; j < arrLen; j++ {
			if arr[i] == arr[j] {
				return arr[i]
			}
		}
	}

	return -1
}

func sol2(arr []int) int {
	mapCheck := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		mapCheck[arr[i]]++
	}

	for key, value := range mapCheck {
		if value > 1 {
			return key
		}
	}

	return -1
}

func sol3(arr []int, n int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum - n*(n-1)/2
}

func sol4(arr []int) int {
	res := 0

	for i := 0; i < len(arr); i++ {
		res ^= i
		res ^= arr[i]
	}

	return res
}
