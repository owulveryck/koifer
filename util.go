package koifer

import (
	"math/rand"
	"time"
)

func generateRandomToken() string {
	output := make([]rune, 32)

	min := 0
	max := 127

	for i := 0; i < 32; i++ {
		// Seed the random number generator with the current time
		rand.Seed(time.Now().UnixNano())

		randomNumber := rand.Intn(max-min+1) + min
		output[i] = rune(randomNumber)
	}
	return string(output)
}
