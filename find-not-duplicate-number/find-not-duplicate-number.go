package main

import "fmt"

func main() {
	arr := []int{7, 3, 5, 4, 5, 3, 4, 7, 8}
	fmt.Println(solution1(arr))
	fmt.Println(solution2(arr))
	fmt.Println(solution3(arr))
	fmt.Println(solution4(arr))
}

func solution4(arr []int) int {
	var sumSingle, sumAll int
	mapCheck := make(map[int]bool)
	for _, item := range arr {
		if _, ok := mapCheck[item]; !ok {
			sumSingle += item
			mapCheck[item] = true
		}

		sumAll += item
	}

	return sumSingle*2 - sumAll
}

func solution3(arr []int) int {
	var result int
	for _, item := range arr {
		result = result ^ item
	}

	return result
}

func solution2(arr []int) int {
	mapCheck := make(map[int]int)

	for _, item := range arr {
		mapCheck[item]++
	}

	for number, value := range mapCheck {
		if value == 1 {
			return number
		}
	}

	return -1
}

func solution1(arr []int) int {
mainLoop:
	for i := 0; i < len(arr); i++ {
		pivot := arr[i]

		for j := 0; j < i; j++ {
			if arr[j] == pivot {
				continue mainLoop
			}
		}

		for j := i + 1; j < len(arr); j++ {
			if arr[j] == pivot {
				continue mainLoop
			}
		}

		return pivot
	}

	return -1
}
