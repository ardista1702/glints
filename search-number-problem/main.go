package main


import (
	"fmt"
	"strconv"
)

func searchNumber(nums []int) int {
	left,right := 1,len(nums)-2
	for left < right{
		middle := (left + right) / 2
		expected := nums[0] + middle
		if nums[middle] < nums[expected]{
			left = middle + 1
		}else{
			right = middle
		}
	}
	return nums[left] + 1
}

func main() {
	// nums := []int{1,2,3,4,5,7}
	// res := searchNumber(nums)
	// fmt.Println(res)
	res1 := splitNumbers("101102103105106107")
	fmt.Println(res1)
}
