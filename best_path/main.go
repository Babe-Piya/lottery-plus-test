package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	inputArray := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	inputFormJson := readInputArray2D()

	fmt.Println("result form array: ", bestPath(inputArray))
	fmt.Println("result form json: ", bestPath(inputFormJson))
}

func bestPath(inputArray [][]int) int {
	result := 0
	index := 0
	for _, input := range inputArray {
		moreNum := 0
		path := input
		if len(input) > 2 {
			path = input[index : index+2]
		}
		for i, num := range path {
			if num > moreNum {
				moreNum = num
				index = i
			}
		}
		result += moreNum
	}

	return result
}

func readInputArray2D() [][]int {
	data, err := os.ReadFile("best_path/hard.json")
	if err != nil {
		fmt.Printf("Error reading file: %s", err)
	}

	var result [][]int

	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("Error unmarshalling data: %s", err)
	}

	return result
}
