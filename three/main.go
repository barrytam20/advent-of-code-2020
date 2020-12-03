package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	a, b, c, d, e := countTrees(1, 1), countTrees(3, 1), countTrees(5, 1), countTrees(7, 1), countTrees(1, 2)
	fmt.Println("right 1, down 1:", a)
	fmt.Println("right 3, down 1:", b)
	fmt.Println("right 5, down 1:", c)
	fmt.Println("right 7, down 1:", d)
	fmt.Println("right 1, down 2:", e)

	fmt.Println("multiplied together", a*b*c*e*d)

	slopes := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	trees := countTreesFaster(slopes)
	for i := range trees {
		fmt.Printf("right %d, down %d: %d\n", slopes[i][1], slopes[i][0], trees[i])
	}
}

func countTrees(right, down int) int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	x := 0
	row := 0
	numTrees := 0
	for scanner.Scan() {
		line := scanner.Text()

		if row%down == 0 {
			if x >= len(line) {
				x = x % len(line)
			}
			if line[x] == '#' {
				numTrees++
			}
			x += right
		}

		row++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return numTrees
}

// slopes is a array of []int{down, right}
func countTreesFaster(slopes [][]int) []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	points := make([][]int, len(slopes)) // points is a slice of []int{row, col}
	for i := range points {
		points[i] = []int{0, 0}
	}
	numTrees := make([]int, len(slopes))
	for scanner.Scan() {
		line := scanner.Text()

		for i := range points {
			if points[i][0]%slopes[i][0] == 0 {
				if points[i][1] >= len(line) {
					points[i][1] = points[i][1] % len(line)
				}
				if line[points[i][1]] == '#' {
					numTrees[i]++
				}

				points[i][1] += slopes[i][1]
			}
			points[i][0]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return numTrees
}
