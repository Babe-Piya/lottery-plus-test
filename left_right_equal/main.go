package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter text: ")
	var input string

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("result form Left-Right: ", leftRightEqual(input))
}

func leftRightEqual(s string) string {
	var nums []int
	sLower := strings.ToLower(s)
	for i, r := range sLower {
		fmt.Println("rune ", string(r))
		if i == 0 {
			nums = append(nums, 0)
		}
		if r == 'l' {
			nums = append(nums, nums[i])
			for j := 0; j < i+1; j++ {
				nums[j] = nums[j] + 1
			}
		} else if r == 'r' {
			nums = append(nums, nums[i]+1)
		} else if r == '=' {
			nums = append(nums, nums[i])
		} else {
			return "Error text input"
		}
	}

	var result string
	for _, num := range nums {
		result += fmt.Sprintf("%d", num)
	}

	return result
}
