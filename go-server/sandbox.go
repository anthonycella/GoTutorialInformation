package main

import (
	"fmt"
	"runtime"
	"time"
)

func factorial(inputNumber int) int {
	if inputNumber < 2 {
		return 1
	}

	factorialResult := 1

	for inputNumber > 0 {
		factorialResult *= inputNumber
		inputNumber -= 1
	}

	return factorialResult
}

func factorialRecursive(inputNumber int) int {
	if inputNumber < 2 {
		return 1
	}

	return inputNumber * factorialRecursive(inputNumber-1)
}

func yodaSwap(firstString, secondString string) (string, string) {
	return secondString, firstString
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func goOperatingSystem() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	default:
		return "Windows"
	}
}

func printTimeOfDay() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	case t.Hour() < 22:
		fmt.Println("Good evening!")
	default:
		fmt.Println("Good night...zzzzzzz")
	}
}

func whenIsFriday() string {
	today := time.Now().Weekday()
	// fmt.Println("time.Now().Weekday() results to", today)

	switch time.Friday {
	case (today + 0) % 7:
		return "Thank God it's today!"
	case (today + 1) % 7:
		return "Tomorrow"
	case (today + 2) % 7:
		return "We're getting there"
	case (today + 3) % 7:
		return "Almost getting close"
	default:
		return "Never gonna happen"
	}
}

func usePointers() {
	firstNumber := 90
	pointerNumber := &firstNumber

	fmt.Println("First number is", firstNumber)
	fmt.Println("Pointer number is", *pointerNumber)

	*pointerNumber = 232023

	fmt.Println("First number is", firstNumber)
	fmt.Println("Pointer number is", *pointerNumber)
}

func playWithDucks() {
	var ducks [10]RubberDuck

	for i := 0; i < len(ducks); i++ {
		ducky := &ducks[i]
		ducky.fillDefaults()
	}

	for j := 0; j < len(ducks); j++ {
		printDuckInfo(ducks[j])
	}
	// ducky := RubberDuck{}
	// ducky.fillDefaults()
	// printDuckInfo(ducky)
	// fmt.Println()

	// ducky.makeJavaScript()
	// printDuckInfo(ducky)
	// fmt.Println()

	// deadDucky := &ducky

	// deadDucky.kill()
	// printDuckInfo(ducky)
	// fmt.Println()

	// otherDucky := &RubberDuck{name: "unknown"}
	// fmt.Println(otherDucky)
}

func whyDidTheSeniorEngineerQuit() {
	punchline := []string{"because", "he", "didn't", "get", "arrays"}

	// slices do not store any data, they just point to the array of origin
	punch := punchline[:2:4]

	for i := 0; i < len(punch); i++ {
		fmt.Print(punch[i] + " ")
	}

	fmt.Println()

	for i := 0; i < cap(punch); i++ {
		fmt.Print(punchline[i] + " ")
	}

	fmt.Println()
}

// var c, python, java bool
// var c, python, java = true, false, "no!"

func main() {
	// fmt.Println("I am Batman")
	// fmt.Println("I am going to kill you", time.Now())
	// fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Printf("Now you have %g problems. \n", math.Sqrt(4))
	// fmt.Println(math.Pi)
	// fmt.Println(yodaSwap("I am", "Batman"))
	// var i int
	// fmt.Println(i, c, python, java)

	// fmt.Println("Go runs on", goOperatingSystem())
	// fmt.Println(whenIsFriday())
	// defer fmt.Println("Go runs on", goOperatingSystem())
	// printTimeOfDay()
	whyDidTheSeniorEngineerQuit()
}
