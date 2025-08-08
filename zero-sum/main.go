package main

import (
	"strconv"
	"strings"
)


// func isPossibleZeroSum(nums []int,start int,sum int) bool {
// 	if start == len(nums){
// 		return  sum == 0
// 	}
// 	if isPossibleZeroSum(nums,start + 1,sum + nums[start]){
// 		return true
// 	}
	
// 	if isPossibleZeroSum(nums,start+1,sum-nums[start]){
// 		return true
// 	}
// 	return false
// }
func isPossibleZeroSum(nums []int) bool{
	dp:= make(map[int]bool)
	dp[0] = true

	for _,num := range nums {
		nextDp := make(map[int]bool)
		for sum := range dp {
			nextDp[sum + num] = true
			nextDp[sum - num] = true
		}
		dp = nextDp
	}
	_,found := dp[0]
	return found
}


func extractData(s string) []int {
	splittedData := strings.Split(s," ")
	var res []int
	for _,str := range splittedData{
		convertedStr,_ := strconv.Atoi(str)
		if convertedStr > 0 {
			res = append(res, convertedStr)
		}
	}
	return res
}