package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Receive instructions from stdin.
	instructions, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("Read Error:", err)
	}

	santa := NewSanta(string(instructions))
	santa.Deliver()

	fmt.Printf("\n\nSanta delivered presents to %d houses.",
		santa.HousesVisited())

	/* End Part 1, Starting Part two with RoboSanta */

	var sInstructions, rsInstructions string

	// Divide instructions in two
	for index, move := range instructions {
		if index%2 == 0 {
			sInstructions += string(move)
		} else {
			rsInstructions += string(move)
		}
	}

	// Give original Santa and RoboSanta instructions:
	santa.Instructions = sInstructions
	roboSanta := NewSanta(rsInstructions)

	santa.Deliver()
	roboSanta.Deliver()

	fmt.Printf("\n\nNext year, with RoboSanta, Santa delivered presents to %d house(s).",
		santa.HousesVisited())

	fmt.Printf("\n\nRoboSanta delivered presents to %d house(s).",
		roboSanta.HousesVisited())

	fmt.Printf("\n\nSanta and RoboSanta visited %d house(s) in total.\n",
		CombinedHousesVisited(santa, roboSanta))

	fmt.Printf("\n\n")

	return
}

func CombinedHousesVisited(s1 Santa, s2 Santa) int {
	masterList := s1.DeliveryList

	for index, amount := range s2.DeliveryList {
		masterList[index] += amount
	}

	return len(masterList)
}
