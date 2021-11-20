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

// Returns a random message which by the caller, will interpolate a given string within the message
func randomMessagem() string {
	messages := []string{
		"Hey %v, how are you doing!",
		"Haaaooowwww %v, whats up!",
		"Good to see you %v!",
	}

	return messages[rand.Intn(len(messages))]
}

//Takes a string as argument and interpolates with the a random message.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomMessagem(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names within random messages
	messsages := make(map[string]string)
	//Loop through all the names to generate a new random message
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//Associate the message with the respective name.
		messsages[name] = message
	}
	return messsages, nil
}
