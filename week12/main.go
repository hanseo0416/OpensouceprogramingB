package main

import (
	"fmt"
	"time"
)

func main() {
	// var dates [3]time.Time
	// dates[0] = time.Unix(0, 0)
	// dates[2] = time.Unix(1708012345, 0)
	// fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, 2024-02-16 ...

	// dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(17080123456, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(17080123456, 0)}
	fmt.Println(dates[0], dates[1], dates[2])
	fmt.Println(dates)         // array
	fmt.Printf("%#v\n", dates) // array literal

}
