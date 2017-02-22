package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type1:", reflect.TypeOf(x))

	x1 := 100
	fmt.Println("type2:", reflect.TypeOf(strconv.Itoa(x1)))
}
