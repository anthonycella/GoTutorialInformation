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

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
