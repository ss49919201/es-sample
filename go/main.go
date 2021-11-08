package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type data struct {
	Body      string    `json:"body"`
	Number    int       `json:"number"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	rand.Seed(time.Now().Unix())
	charSet := "abcdedfghijklmnopqrst"

	for i := 1; i <= 10; i++ {
		rs := string(charSet[rand.Intn(len(charSet))])

		d := data{
			Body:      rs,
			Number:    rand.Intn(100),
			CreatedAt: time.Now(),
		}

		b, err := json.Marshal(&d)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	}
}
