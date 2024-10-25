package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	r := fmt.Sprintf("%d\n", rand.Intn(100)+1)
	fmt.Println(reflect.TypeOf(r))
	fmt.Printf("%T\n", r)

	i := 1
	for i <= 100 {
		fmt.Printf("%3d점\n", i)
		i = i + 1
	}
}
