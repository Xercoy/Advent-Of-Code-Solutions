package main

import (
	"log"
	"strconv"
	"strings"
)

type dimensions struct {
	Length float64
	Width  float64
	Height float64
}

// Returns -1 on erroneous parsing.
func ParseDimension(dStr string) (float64, float64, float64) {
	xIndex1 := strings.Index(dStr, "x")
	xIndex2 := strings.LastIndex(dStr, "x")

	if xIndex1 == -1 || xIndex2 == -1 {
		return -1, -1, -1
	}

	length, err := strconv.Atoi(dStr[0:xIndex1])
	if err != nil {
		log.Fatalln(err.Error())
	}

	width, err := strconv.Atoi(dStr[xIndex1+1 : xIndex2])
	if err != nil {
		log.Fatalln(err.Error())
	}

	height, err := strconv.Atoi(dStr[xIndex2+1:])
	if err != nil {
		log.Fatalln(err.Error())
	}

	return float64(length), float64(width), float64(height)
}
