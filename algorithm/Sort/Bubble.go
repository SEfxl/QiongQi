package main

import "fmt"

func main()  {
	Bubble()
}

//冒泡排序
func Bubble() {
	a := [...]int{1, 3, 5, 7, 9, 0, 2, 4, 8, 6}
	arrLength := len(a)

	for i := 0; i < arrLength; i++ {
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}

