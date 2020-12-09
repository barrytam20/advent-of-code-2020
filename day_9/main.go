package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	input := getInput()

	invalidNum := findInvalidNum(input, 25)

	sub := subsetSum(invalidNum, input)
	fmt.Println(findHighLowDiff(sub))
}

func getInput() []int {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	out := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		v, _ := strconv.Atoi(line)
		out = append(out, v)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return out

}

func findInvalidNum(nums []int, k int) int {
	for i := k; i < len(nums); i++ {
		if !isValid(nums[i-k:i], nums[i]) {
			return nums[i]
		}
	}
	return -1
}

func isValid(nums []int, target int) bool {
	pre := map[int]bool{}
	for _, n := range nums {
		if pre[n] {
			return true
		}
		pre[target-n] = true
	}
	return false
}

func subsetSum(target int, nums []int) []int {
	for i := range nums {
		remaining := target
		j := i
		for j < len(nums) && remaining > 0 {
			remaining -= nums[j]
			j++
		}
		if remaining == 0 {
			return nums[i:j]
		}
	}

	return []int{}
}

func findHighLowDiff(nums []int) int {
	high, low := math.MinInt64, math.MaxInt64
	sum := 0
	for _, n := range nums {
		sum += n
		if n > high {
			high = n
		}
		if n < low {
			low = n
		}
	}
	return high + low
}
