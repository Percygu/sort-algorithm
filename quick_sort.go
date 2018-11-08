package main

import (
	"fmt"
)

//一趟排序
func sortOnce(a []int, left int, right int) int {
	i := left
	j := right
	num := a[left] //选取数列的第一个数为基准数
	for i < j {
		for i < j && a[j] >= num {
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}
		for i < j && a[i] <= num {
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	a[i] = num
	return i
}

func quickSort(a []int, left int, right int) {
	if left < right {
		i := sortOnce(a, left, right)
		quickSort(a, left, i-1)
		quickSort(a, i+1, right)
	}
}

func main() {

	numlist := []int{2, 4, 1, 5, 3, 6, 9, 0, 8, 7}
	quickSort(numlist, 0, len(numlist)-1)
	for i := range numlist {
		fmt.Println(i)
	}
}
