package main

import (
	"fmt"
	"math"
)

var oneApplePrice float64 = 5.99
var onePearPrice float64 = 7
var weHaveMoney float64 = 23
var appleCount float64
var pearCount float64

func main() {
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	appleCount = 9
	pearCount = 8
	var cost9ApplesAnd8Pears float64 = oneApplePrice * appleCount + onePearPrice * pearCount
	fmt.Println("Щоб купити 9 яблук та 8 груш ми маємо витратити " , cost9ApplesAnd8Pears, "грн.")

	fmt.Println("2. Скільки груш ми можемо купити?")
	var howManyPears float64 = weHaveMoney / onePearPrice
	fmt.Println("Ми можемо купити", math.Round(howManyPears) , "груші.")

	fmt.Println("3. Скільки яблук ми можемо купити?")
	var howManyApples float64 = weHaveMoney / onePearPrice
	fmt.Println("Ми можемо купити", math.Round(howManyApples), "яблука.")

	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")
	appleCount = 2
	pearCount = 2
	var twoApplesCost = oneApplePrice * appleCount
	var twoPearsCost = onePearPrice * pearCount
	var allPrice = twoApplesCost + twoPearsCost
	var canBuyTwoPearsAndTwoApples = allPrice <= weHaveMoney
	if canBuyTwoPearsAndTwoApples {
		fmt.Println("Ми можемо купити 2 груші та 2 яблука")
	} else {
		fmt.Println("Ми не можемо купити 2 груші та 2 яблука")
	}
}
