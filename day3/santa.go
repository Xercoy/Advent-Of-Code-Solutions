package main

type Position struct {
	x int
	y int
}

type Santa struct {
	Instructions string
	DeliveryList map[Position]int
}

func NewSanta(i string) Santa {
	return Santa{
		Instructions: i,
		DeliveryList: make(map[Position]int),
	}
}

func (s *Santa) Deliver() {
	var x, y int

	s.DeliveryList = make(map[Position]int)

	// Starting location gets 1 present.
	s.DeliveryList[Position{x, y}] = 1

	for _, move := range s.Instructions {
		switch move {
		case '^':
			y += 1
		case '<':
			x -= 1
		case '>':
			x += 1
		case 'v':
			y -= 1
		}

		s.DeliveryList[Position{x, y}] += 1
	}

	return
}

func (s *Santa) HousesVisited() int {
	return len(s.DeliveryList)
}
