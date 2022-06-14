package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isValidContainer("CSQU3054390"))
	fmt.Println(isValidContainer("CCLU4843216"))
}

var containerCalculatorValues = map[string]int{
	"A": 10,
	"B": 12,
	"C": 13,
	"D": 14,
	"E": 15,
	"F": 16,
	"G": 17,
	"H": 18,
	"I": 19,
	"J": 20,
	"K": 21,
	"L": 23,
	"M": 24,
	"N": 25,
	"O": 26,
	"P": 27,
	"Q": 28,
	"R": 29,
	"S": 30,
	"T": 31,
	"U": 32,
	"V": 34,
	"W": 35,
	"X": 36,
	"Y": 37,
	"Z": 38,
}

func isValidContainer(containerNumber string) bool {
	containerNoLength := len(containerNumber)

	if containerNoLength < 11 {
		return false
	}

	var withoutLastChar string = containerNumber[:containerNoLength-1]

	lastChar := containerNumber[containerNoLength-1:]

	multiplier := 1

	valueSum := 0

	for _, item := range strings.Split(withoutLastChar, "") {
		equivalentFactor := getCalculatorValue(item) * multiplier

		valueSum += equivalentFactor

		multiplier *= 2
	}

	subTotal := int(math.Floor(float64(valueSum)/float64(containerNoLength))) * containerNoLength
	checkDifference := valueSum - subTotal
	checkDigit := checkDifference

	if checkDifference == 10 {
		checkDigit = 0
	}

	return checkDigit == toInt(lastChar)
}

func getCalculatorValue(item string) int {
	if containerCalculatorValues[item] == 0 {
		return toInt(item)
	} else {
		return containerCalculatorValues[item]
	}
}

func toInt(str string) int {
	intVar, _ := strconv.Atoi(str)
	return intVar
}
