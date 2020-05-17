package exercises

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xuongtrantrieu/go-learning/utils"
)

const (
	SpecialCharacter  = "!@#$%^&"
	AlphabetCharacter = "abcdefghijklmnoptuv"
	DigitCharacter    = "1234567890"
)

func PasswordGenerator() {
	var length int
	fmt.Printf("Enter password length: ")
	_, err := fmt.Scan(&length)
	utils.CheckErr(err)

	passwordChars := make([]byte, length)

	remainingLength := length

	specialCharacterLength := rand.Intn(remainingLength) + 1
	remainingLength -= specialCharacterLength

	var alphabetLength int
	if remainingLength > 0 {
		alphabetLength = rand.Intn(remainingLength) + 1
		remainingLength -= alphabetLength
	}

	var digitLength int
	if remainingLength > 0 {
		digitLength = length - remainingLength
	}

	var getRandomByteFromString = func(s string) byte {
		pos := rand.Intn(len(s))
		return []byte(s)[pos]
	}

	for i := 0; i < length; i++ {
		randoming := true

		var chars string
		for randoming {
			rand.Seed(time.Now().UnixNano())
			switch x := rand.Intn(3); x {
			case 0:
				chars = SpecialCharacter
				if specialCharacterLength > 0 {
					specialCharacterLength--
					randoming = false
				}
			case 1:
				chars = AlphabetCharacter
				if alphabetLength > 0 {
					alphabetLength--
					randoming = false
				}
			default:
				chars = DigitCharacter
				if digitLength > 0 {
					digitLength--
					randoming = false
				}
			}
		}

		passwordChars[i] = getRandomByteFromString(chars)
	}

	fmt.Printf("Your password: %s\n", string(passwordChars))
}
