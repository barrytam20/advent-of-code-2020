package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	adaptors, highest := getAdaptors()

	// part 1
	d1, d3 := getJoltDiffs(adaptors)
	fmt.Println(d1, d3, d1*d3)

	// part 2
	fmt.Println(getNumCombinationsDP(adaptors, highest))
}

func getAdaptors() (map[int]bool, int) {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	out := map[int]bool{0: true}
	highest := 0
	for scanner.Scan() {
		line := scanner.Text()
		v, _ := strconv.Atoi(line)
		out[v] = true
		if v > highest {
			highest = v
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	out[highest+3] = true
	return out, highest + 3

}

func getJoltDiffs(adaptors map[int]bool) (int, int) {
	d1, d3 := 0, 0
	curr := 0

	for adaptors[curr] {
		if adaptors[curr+1] {
			d1++
			curr++
		} else if adaptors[curr+2] {
			curr += 2
		} else if adaptors[curr+3] {
			d3++
			curr += 3
		} else {
			break
		}
	}
	return d1, d3
}

func getNumCombinationsDP(adaptors map[int]bool, last int) int {
	dp := make([]int, last+1)
	dp[0] = 1
	for i := range dp {
		if !adaptors[i] {
			continue
		}

		if i > 0 && adaptors[i-1] {
			dp[i] += dp[i-1]
		}
		if i > 1 && adaptors[i-2] {
			dp[i] += dp[i-2]
		}
		if i > 2 && adaptors[i-3] {
			dp[i] += dp[i-3]
		}
	}

	return dp[last]
}
