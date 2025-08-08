package main

import (
	"strings"
)

func reverse(s string)string {
	sSplit := strings.Split(s,"")
	left,right := 0,len(s) - 1
	for left < right{
		sSplit[left],sSplit[right] = sSplit[right],sSplit[left]
		left++
		right--
	}
	return strings.Join(sSplit,"")

}

func solve(s string) string {
	sSplit := strings.Split(s, " ")
	for i := range sSplit{
		sSplit[i] = reverse(sSplit[i])
	}
	return strings.Join(sSplit," ")
}

