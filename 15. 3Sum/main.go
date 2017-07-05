package main

import "fmt"

func main() {

	var nums  = []int {-1, 0, 1, 2, -1, -4}

	threeSum(nums)
}

func threeSum(nums []int) [][]int {
	var res  [][] int
	var values = make(map[int] int)

	for tk,tv := range nums {
		if _, ok := values[tv]; ok {
			 nums = append(nums[:tk],nums[tk+1:]...)
		} else {
			values[tv] = 1
		}
	}
	for k,v := range nums {
		for kk, vv := range nums[k+1:] {
			for _, vvv := range nums[kk+1:] {
				if v + vv + vvv == 0 {
					res = append(res, []int{v,vv,vvv})
					fmt.Printf("%v %v %v\n", v,vv,vvv)
				}
			}
		}
	}

	return res
}