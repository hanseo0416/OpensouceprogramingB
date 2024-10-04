package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	f := 12.9
	c1 := 'Z'
	c2 := '김'

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	fmt.Println(c1, c2)    // 10진수
	fmt.Printf("%x\n", c2) // 유니코드 '김'에 대한 16진수 출력,home.unicode.org
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i), reflect.TypeOf(c1), reflect.TypeOf(c2))
}
