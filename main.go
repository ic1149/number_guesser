package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func main() {
	var min int = 1
	var max int = 100
	var target int = rand.IntN(max) + min
	var guess int = 0
	var input string
	var err error = nil
	var tries uint = 0
	for target != guess {
		fmt.Printf("guess (%v<=x<=%v): ", min, max)
		fmt.Scanln(&input)
		guess, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			continue

		} else if guess > target {
			if guess < max {
				max = guess - 1
			}
			fmt.Println("too big")

		} else if guess < target {
			if guess > min {
				min = guess + 1
			}
			fmt.Println("too small")
		}
		tries++

	}
	fmt.Printf("correct, you got it in %v tries\n", tries)

}
