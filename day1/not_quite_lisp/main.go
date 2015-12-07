package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	instructions, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\nSanta will end up on floor #%d.\n",
		FinalFloorLevel(string(instructions)))

	fmt.Printf("\nHe will enter the basement level -1 at position #%d\n\n",
		BasementLevelIndex(string(instructions)))

	return
}

func BasementLevelIndex(instructions string) int {
	var position int
	var floorLevel int

	for index, character := range instructions {
		switch character {
		case '(':
			floorLevel += 1

		case ')':
			floorLevel -= 1
		}

		if floorLevel == -1 {
			position = index + 1
			break
		}
	}

	return position
}

// if ( go up, if ) go down.
func FinalFloorLevel(instructions string) int {
	var floorLevel int

	for _, character := range instructions {
		switch character {
		case '(':
			floorLevel += 1

		case ')':
			floorLevel -= 1
		}
	}

	return floorLevel
}
