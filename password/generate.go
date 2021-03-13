package password

import (
	"math/rand"
	"time"
)

func GenerateRand(length uint, spec bool) string {
	rand.Seed(time.Now().UnixNano())

	const symbols string = "qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM"
	const symbolsSpec string = "qwertyuiopasdfghjklzxcvbnm1234567890!@#$%&*-_=~?QWERTYUIOPASDFGHJKLZXCVBNM"
	var result string = ""

	for i := 1; i <= int(length); i++ {
		if spec == false {
			rndIndex := rand.Intn(len(symbols))
			result += string(symbols[rndIndex])
		} else {
			rndIndex := rand.Intn(len(symbolsSpec))
			result += string(symbolsSpec[rndIndex])
		}
	}

	return result
}
