package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	nums := getNums()
	for i := 0; i <= len(nums)-2; i++ {
		target := 2020 - nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				fmt.Println(nums[i] * nums[left] * nums[right])
				return
			}
			if sum < target {
				left++
				for left < right && nums[left-1] == nums[left] {
					left++
				}
			} else {
				right--
				for left < right && nums[right+1] == nums[right] {
					right--
				}
			}
		}
	}
}

func getNums() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	nums := []int{}
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(nums)
	return nums
}
