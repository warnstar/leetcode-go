package main

import "fmt"

func main() {
	var input = []int{5,2,4}

	fmt.Printf("%v", twoSum(input,6))
}

func twoSum(nums []int, target int) (data []int) {
	for index, value := range nums {
		for indexSub, valueSub := range nums[index+1:] {
			fmt.Printf("index: %v %v | value: %v %v \n",index,index+1+indexSub,value,valueSub)
			if value + valueSub == target {
				return append(data,index,index + 1 + indexSub)
			}
		}
	}

	return data
}