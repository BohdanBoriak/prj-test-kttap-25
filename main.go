package main

import (
	"encoding/json"
	"errors"
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
	pointsPerQuestion = 50
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у MathGame!")

	for {
		menu()

		choice := ""
		fmt.Scan(&choice)

		switch choice {
		case "1":
			user := play()
			users := getUsers()
			users = append(users, user)
			sortAndSave(users)
		case "2":
			users := getUsers()
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
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("sortAndSave(file.Close): %s", err)
		}
	}()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		log.Printf("sortAndSave(encoder.Encode): %s", err)
	}
}

func getUsers() []domain.User {
	var users []domain.User

	file, err := os.Open("users.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_, err = os.Create("users.json")
			if err != nil {
				log.Printf("getUsers(os.Create): %s", err)
			}
			return nil
		} else {
			log.Printf("getUsers(os.Open): %s", err)
			return nil
		}
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("getUsers(file.Close): %s", err)
		}
	}()

	fInfo, err := os.Stat("users.json")
	if err != nil {
		log.Printf("getUsers(os.Stat): %s", err)
	}

	if fInfo.Size() != 0 {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&users)
		if err != nil {
			log.Printf("getUsers(decoder.Decode): %s", err)
		}
	}

	return users
}
