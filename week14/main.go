package main

import "fmt"

func main() {
	ages := make(map[string]int)
	var name string
	var age int

	for {
		fmt.Print("Name :")
		fmt.Scanln(&name)
		if name == "q" {
			break
		}
		fmt.Print("age :")
		fmt.Scanln(&age)

		ages[name] = age

	}
	for key, value := range ages {
		fmt.Printf("%s is %d year old\n", key, value)
	}
}
