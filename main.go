package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	totalPoints       = 50
	pointsPerQuestion = 10
)

func main() {
	fmt.Println("Вітаємо у MathGame!")

	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	start := time.Now()
	play()
	finish := time.Now()
	timeSpent := finish.Sub(start)
	fmt.Printf("Вітаю! Ти впорався за: %v", timeSpent)
	time.Sleep(10 * time.Second)
}

func play() {
	myPoints := 0
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)

		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("НЕ ПРАВИЛЬНО!")
		} else {
			if x+y == ansInt {
				myPoints += pointsPerQuestion
				fmt.Printf("Ви набрали: %v очок!\n", myPoints)
			} else {
				fmt.Println("НЕ ПРАВИЛЬНО!")
			}
		}
	}
}
