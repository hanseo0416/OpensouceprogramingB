package main

import (
	"fmt"
	"os"
	"reflect"
)

func test(i int, strs ...string) {
	//func test(strs ...string,i int) {  // error

	fmt.Println(i, strs, reflect.TypeOf(strs))
}
func main() {
	// fmt.Println(os.Args, len(os.Args), reflect.TypeOf(os.Args))
	slices := os.Args[1:]
	fmt.Println(slices)
	test(10, "123")
	test(20, "123, abc")
	test(55)
	test(30, "123, abc, 123a")
}
