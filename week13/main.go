package main

import (
	"fmt"
)

func main() {
	var emptyslice []bool
	emptyslice = make([]bool, 5)
	fmt.Printf("%#v\n", emptyslice)

	gpa := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23}
	gpa_slice := gpa[1:4]
	//gpa_slice[1] = 2.71
	gpa_slice[1] = 2.71
	gpa_slice = append(gpa_slice, 4.3, 5.55)
	fmt.Println(len(gpa_slice), gpa_slice, gpa)

}
