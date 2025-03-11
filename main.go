package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"prj-test/domain"
	"sort"
	"strconv"
	"time"
)

const (
	totalPoints       = 50
	pointsPerQuestion = 10
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у MathGame!")

	var users []domain.User
	for {
		menu()

		choice := ""
		fmt.Scan(&choice)

		switch choice {
		case "1":
			user := play()
			users = append(users, user)
		case "2":
			for _, u := range users {
				fmt.Printf(
					"Id: %v, Name: %s, Time: %v\n",
					u.Id, u.Name, u.Time,
				)
			}
		case "3":
			return
		default:
		}
	}
}

func menu() {
	fmt.Println("1. Грати")
	fmt.Println("2. Рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	start := time.Now()
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
	finish := time.Now()
	timeSpent := finish.Sub(start)
	fmt.Printf("Вітаю! Ти впорався за: %v", timeSpent)

	fmt.Println("Введіть ім'я: ")
	name := ""
	fmt.Scan(&name)

	// var u domain.User
	// u.Id = id
	// u.Name = name
	// u.Time = timeSpent

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile(
		"users.json",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755,
	)
	if err != nil {
		log.Printf("sortAndSave(os.OpenFile): %s", err)
	}
}
