package main

import "fmt"

func getKey(nums []int, target int) int {
	var val = 0
	
	if len(nums) == 1 {
		for m := range nums {
			if target > nums[m] {
				return 1
			} else {
				return m
			}
		}
	}

	for i := range nums {
		if target < nums[0] {
			return 0
		}
		if target == nums[i] {
			val = i
			break
		}
		
		if val == 0 {
			for j := range nums {
				rag := len(nums)
				if rag == j+1 {
					brg := target - nums[j]
					for l := 1; l < brg; l++ {
						val = j+1
					}

					if val == 0 {
						val = j+1
					}
					break
				}

				intv := nums[j+1] - nums[j]				
				for k := 1; k < intv; k++ {
					if target == nums[j]+k {
						val = j+1
					}
				}
				
				if val == 0 {
					if target == nums[j]+1 {
						val = j+1
					}
				}
			}
		}
	}

	return val
}

func main() {
	var arr = []int{-1,3,5,6}
	fmt.Println(getKey(arr, 0))
	// getKey(arr, 5)
}