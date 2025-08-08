package main

import (
	"fmt"
	"lift-problem/queue"
	"strconv"
	"strings"
)




func main() {
	s := "2 5 0 0 4"
	filtereRequest := filterRequest(s)
	res := elevator(filtereRequest)
	fmt.Println(res)
}


func elevator(filteredRequest []int) int {
	if len(filteredRequest) == 0 {
		return 1
	}
	q := queue.NewQueue[int]()
	q.PushBack(filteredRequest[0])

	for i := 1; i < len(filteredRequest); i++ {
		if q.Head != nil && q.Head.Val == i {
			q.PopFront()
		}

		// Kalau queue kosong, masukkan request baru dan lanjut ke iterasi berikutnya
		if q.Head == nil {
			q.PushBack(filteredRequest[i])
			continue
		}

		// jika jarak request terbaru dari lantai ini = jarak request saat ini dari lantai saat ini
		if abs(q.Head.Val-i) == abs(filteredRequest[i]-i) {
			// jika saat ini sedang naik, tetap naik; jika turun tetap turun
			q.PushBack(filteredRequest[i])
		} else if abs(q.Head.Val-i) < abs(filteredRequest[i]-i) {
			// ambil yang paling dekat dulu
			q.PushFront(filteredRequest[i])
		} else {
			q.PushBack(filteredRequest[i])
		}
	}

	if q.Tail != nil {
		return q.Tail.Val + 1
	}
	return -1 // tidak ada request yang tersisa
}

func filterRequest(s string) []int{
	splittedData := strings.Split(s, " ")
	var res []int
	for _,str := range splittedData {
		convertedStr,_ := strconv.Atoi(str)
		if convertedStr != 0 {

			res = append(res, convertedStr - 1)
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return  num * -1
	}
	return num
}