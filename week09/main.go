package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(6) + 1
	fmt.Println(target)

	fmt.Println("숫자 입력")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if target == guess {
		fmt.Println("정답입니다.")
	} else if target > guess {
		fmt.Println("입력하신 값은 정답보다 작은 수 입니다. LOW")
	} else {
		fmt.Println("입력하신 값은 정답보다 큰 수 입니다. HIGH")

	}
}
