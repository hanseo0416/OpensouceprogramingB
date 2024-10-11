package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("이름 입력 : ")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')
	fmt.Println(name)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}
}
