package main

import "fmt"

func main() {
	input := []int{18, 8, 0, 5, 4, 1, 20}

	fmt.Println(play(input, 30000000))
}

func play(input []int, k int) int {
	spoken := map[int]int{}

	i := 0
	for _, n := range input {
		i++
		spoken[n] = i
	}

	prev := input[len(input)-1]
	delete(spoken, prev)

	for i < k {
		i++
		last, ok := spoken[prev]
		spoken[prev] = i - 1
		if !ok || last == i-1 {
			prev = 0
		} else {
			prev = i - last - 1
		}
	}

	return prev
}
