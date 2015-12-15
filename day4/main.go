package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Secret key is first argument.
	input := os.Args[1]

	fmt.Printf("\nThe mystery key of input %s with 5 leading 0s is %d\n\n",
		input, MysteryNumber(input, 5))

	fmt.Printf("\nThe mystery key of input %s with 6 leading 0s is is %d\n\n",
		input, MysteryNumber(input, 6))
}

func MysteryNumber(input string, zeroAmt int) int {
	var (
		mysteryNumber int
		secretKey     string
	)

	for true {
		secretKey = input + strconv.Itoa(mysteryNumber)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(secretKey)))
		leadingZeros := strings.Count(hash[0:zeroAmt], "0")

		if (len(hash) > zeroAmt) && (leadingZeros == zeroAmt) {
			break
		}

		mysteryNumber += 1
	}

	return mysteryNumber
}
