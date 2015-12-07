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

	fmt.Printf("\nSanta will end up on floor #%d.\n\n",
		getFloorLevel(string(instructions)))

	return
}

// if ( go up, if ) go down.
func getFloorLevel(instructions string) (finalFloor int) {
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
