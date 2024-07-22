package main

import "fmt"

func main() {
	var scores [3]int
	scores[0] = 118
	scores[1] = 135
	scores[2] = 126

	var sum int
	var avg int

	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	avg = sum / len(scores)

	fmt.Printf("总分=%d 平均分=%d\n", sum, avg)
	fmt.Printf("&scores = %p\n", &scores)
	fmt.Printf("&scores[0] = %p\n", &scores[0])

	var arr1 [3]int = [3]int{3, 6, 9}
	var arr2 = [3]int{1, 4, 7}
	var arr3 = [...]int{2, 5, 8}
	var arr4 = [...]int{1: 1, 5: 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

}
