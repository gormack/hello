package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(getGreeting())
}

func getGreeting() string {
	return getGreetingVerb() + ", Entire Studio!"
}

func getGreetingVerb() string {
	rand.Seed(time.Now().Unix())

	verbs := []string{
		"Hello",
		"Howdy",
		"Hi",
	}

	return verbs[rand.Intn(len(verbs))]
}
