package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomMessagem() string {
	messages := []string{
		"Hey %v, how are you doing!",
		"Haaaooowwww %v, whats up!",
		"Good to see you %v!",
	}

	return messages[rand.Intn(len(messages))]
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomMessagem(), name)
	return message, nil
}
