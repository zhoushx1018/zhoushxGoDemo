// 数组

package main

import "fmt"

func main() {

	//自动初始化的数组
	var a [5]int
	fmt.Println("emp:", a)

	//数组元素赋值；  数组元素打印
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//数组长度
	fmt.Println("len:", len(a))

	// 定义并初始化数组的语法
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 二维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
