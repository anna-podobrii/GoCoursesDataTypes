package main

import (
	"fmt"
	"math"
)

var oneApplePrice float64 = 5.99
var onePearPrice float64 = 7
var weHaveMoney float64 = 23

func main() {
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	var cost9ApplesAnd8Pears float64 = oneApplePrice * 9 + onePearPrice * 8
	fmt.Println("Щоб купити 9 яблук та 8 груш ми маємо витратити " , cost9ApplesAnd8Pears, "грн.")

	fmt.Println("2. Скільки груш ми можемо купити?")
	var howManyPears float64 = weHaveMoney / onePearPrice
	fmt.Println("Ми можемо купити", math.Round(howManyPears) , "груші.")

	fmt.Println("3. Скільки яблук ми можемо купити?")
	var howManyApples float64 = weHaveMoney / onePearPrice
	fmt.Println("Ми можемо купити", math.Round(howManyApples), "яблука.")

	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")
	var twoApplesCost = oneApplePrice * 2
	var twoPearsCost = onePearPrice * 2
	var allPrice = twoApplesCost + twoPearsCost
	var canBuyTwoPearsAndTwoApples = allPrice <= weHaveMoney
	if canBuyTwoPearsAndTwoApples {
		fmt.Println("Ми можемо купити 2 груші та 2 яблука")
	} else {
		fmt.Println("Ми не можемо купити 2 груші та 2 яблука")
	}
}