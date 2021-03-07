package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%v: %v\n", i, simpleSolution(i))
		fmt.Printf("%v: %v\n", i, betterSolution(i))
	}
}

func simpleSolution(num int) bool {
	if num == 0 {
		return false
	}

	if num == 1 {
		return true
	}

	if num < 0 || num%2 != 0 {
		return false
	}

	for num > 1 {
		if num%2 != 0 {
			return false
		}

		num /= 2
	}

	return true
}

func betterSolution(num int) bool {
	if num < 0 {
		return false
	}

	if num&(num-1) == 0 {
		return true
	}

	return false
}
