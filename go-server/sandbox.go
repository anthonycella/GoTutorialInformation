package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// func factorial(inputNumber int) int {
// 	if inputNumber < 2 {
// 		return 1
// 	}

// 	factorialResult := 1

// 	for inputNumber > 0 {
// 		factorialResult *= inputNumber
// 		inputNumber -= 1
// 	}

// 	return factorialResult
// }

// func factorialRecursive(inputNumber int) int {
// 	if inputNumber < 2 {
// 		return 1
// 	}

// 	return inputNumber * factorialRecursive(inputNumber-1)
// }

// func yodaSwap(firstString, secondString string) (string, string) {
// 	return secondString, firstString
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

// func goOperatingSystem() string {
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		return "OS X"
// 	case "linux":
// 		return "Linux"
// 	default:
// 		return "Windows"
// 	}
// }

// func printTimeOfDay() {
// 	t := time.Now()

// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good afternoon.")
// 	case t.Hour() < 22:
// 		fmt.Println("Good evening!")
// 	default:
// 		fmt.Println("Good night...zzzzzzz")
// 	}
// }

// func whenIsFriday() string {
// 	today := time.Now().Weekday()
// 	// fmt.Println("time.Now().Weekday() results to", today)

// 	switch time.Friday {
// 	case (today + 0) % 7:
// 		return "Thank God it's today!"
// 	case (today + 1) % 7:
// 		return "Tomorrow"
// 	case (today + 2) % 7:
// 		return "We're getting there"
// 	case (today + 3) % 7:
// 		return "Almost getting close"
// 	default:
// 		return "Never gonna happen"
// 	}
// }

// func usePointers() {
// 	firstNumber := 90
// 	pointerNumber := &firstNumber

// 	fmt.Println("First number is", firstNumber)
// 	fmt.Println("Pointer number is", *pointerNumber)

// 	*pointerNumber = 232023

// 	fmt.Println("First number is", firstNumber)
// 	fmt.Println("Pointer number is", *pointerNumber)
// }

// func playWithDucks() {
// 	var ducks [10]RubberDuck

// 	for i := 0; i < len(ducks); i++ {
// 		ducky := &ducks[i]
// 		ducky.fillDefaults()
// 	}

// 	for j := 0; j < len(ducks); j++ {
// 		fmt.Println(ducks[j])
// 	}
// 	ducky := RubberDuck{}
// 	ducky.fillDefaults()
// 	fmt.Println(ducky)
// 	fmt.Println()

// 	ducky.makeJavaScript()
// 	fmt.Println(ducky)
// 	fmt.Println()

// 	deadDucky := &ducky

// 	deadDucky.kill()
// 	fmt.Println(ducky)
// 	fmt.Println()

// 	ducky.changeName("Fred")
// 	fmt.Println("Now my name is", ducky.name)

// 	otherDucky := &RubberDuck{name: "unknown"}
// 	fmt.Println(*otherDucky)
// }

// func mapTheDucks() {
// 	var ducks [10]RubberDuck

// 	for i := 0; i < len(ducks); i++ {
// 		ducky := &ducks[i]
// 		ducky.fillDefaults()
// 	}

// 	// fmt.Println(ducks)

// 	mapOfDucks := make(map[string]RubberDuck)

// 	for _, duck := range ducks {
// 		mapOfDucks[duck.name] = duck
// 	}

// 	// fmt.Println(mapOfDucks)

// 	mapOfDucks = map[string]RubberDuck{
// 		"JavaScript": {"JavaScript", "yellow", 0, true},
// 		"DuckyMoMo":  {"DuckyMoMo", "yellow", 2, false},
// 		"Wilbur":     {"Wilbur", "pink", 21, false},
// 		"DuckJr":     {"DuckJr", "white", 12, false},
// 	}

// 	fmt.Println(mapOfDucks)

// 	mapOfDucks["DuckyMoMo"] = RubberDuck{"Fred", "green", 4, false}
// 	fmt.Println(mapOfDucks)

// 	fred := mapOfDucks["DuckyMoMo"]

// 	fmt.Println(fred)

// 	elem, ok := mapOfDucks["fred"]

// 	if ok == false {
// 		fmt.Println("Fred is not in the map under Fred")
// 	} else {
// 		fmt.Println("Fred is in the map", elem)
// 	}

// 	elem, ok = mapOfDucks["DuckyMoMo"]

// 	if ok {
// 		fmt.Println("Fred is in the map under Ducky MoMo")
// 	} else {
// 		fmt.Println("Fred is not in the map at all")
// 	}

// 	delete(mapOfDucks, "DuckyMoMo")

// 	fmt.Println(mapOfDucks)

// 	mapOfDucks["Fred"] = fred
// 	fmt.Println(mapOfDucks)
// }

// func printBoard(queensBoard [9][9]int) {
// 	for i := 0; i < len(queensBoard); i++ {
// 		currentLine := queensBoard[i]
// 		fmt.Println(currentLine)
// 	}
// }

// func matrixTime() {
// 	queensBoard := [9][9]int{}
// 	printBoard(queensBoard)

// 	fmt.Print("\n\n\n")

// 	for i := 0; i < len(queensBoard); i++ {
// 		queensBoard[i][i] = 1
// 	}

// 	printBoard(queensBoard)
// }

// func whyDidTheSeniorEngineerQuit() {
// 	punchline := []string{"because", "he", "didn't", "get", "arrays"}

// 	// slices do not store any data, they just point to the array of origin
// 	punch := punchline[:2:5]
// 	line := make([]string, 20, 25)
// 	// third argument is optional: capacity

// 	// fmt.Println(len(punch))
// 	// fmt.Println(cap(punch))
// 	// fmt.Println(cap(punchline))

// 	for i := 0; i < len(punch); i++ {
// 		fmt.Print(punch[i] + " ")
// 	}

// 	fmt.Println()

// 	for i := 0; i < cap(punch); i++ {
// 		fmt.Print(punchline[i] + " ")
// 	}

// 	fmt.Println()

// 	for i := range line {
// 		line[i] = "ha "
// 	}

// 	for i, v := range line {
// 		fmt.Println(i, v)
// 	}

// 	for _, v := range line {
// 		fmt.Print(v + " ")
// 	}

// 	fmt.Println()
// }

// func getGroceryList() []string {
// 	var userInput string
// 	var shoppingList []string

// 	fmt.Println("Add an item to the shopping list or hit x to exit")
// 	fmt.Scanln(&userInput)

// 	for userInput != "x" {
// 		shoppingList = append(shoppingList, userInput)
// 		fmt.Println("Add an item to the shopping list or hit x to exit")
// 		fmt.Scanln(&userInput)
// 	}

// 	fmt.Println("returning to main function")
// 	return shoppingList
// }

// func Pic(dx, dy int) [][]uint8 {
// 	var picture [][]uint8

// 	for y := 0; y < dy; y++ {
// 		var row []uint8
// 		uintX := uint8(y)

// 		for x := 0; x < dx; x++ {
// 			uintY := uint8(x)

// 			result := uintX ^ uintY
// 			row = append(row, result)
// 		}

// 		picture = append(picture, row)
// 	}

// 	return picture
// }

// func typeSwitch(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Printf("Twice %v is %v\n", v, v*2)
// 	case string:
// 		fmt.Printf("The length of \"%v\" is %v\n", v, len(v))
// 	case RubberDuck:
// 		fmt.Println("This interface is in the shape of a duck!")
// 	default:
// 		fmt.Println("Let me get back to you on that")
// 	}
// }

// func funWithTypeSwitches() {
// 	typeSwitch(32)
// 	typeSwitch("IKEA sent me duplicates and I'm not happy")
// 	typeSwitch(RubberDuck{})
// 	typeSwitch(true)
// }

// func migrateDucky(duckNumber int) {
// 	ducky := RubberDuck{}
// 	ducky.fillDefaults()

// 	fmt.Println(ducky.name, "is migrating. This is ducky number", duckNumber)
// 	wg.Done()
// }

// func migrateTheDuckies() {
// 	fmt.Println("Number of CPUs running:", runtime.NumCPU())
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go migrateDucky(i + 1)
// 		fmt.Println("Number of go routines running:", runtime.NumGoroutine())
// 	}

// 	wg.Wait()
// }

func dataRacer() {
	const raceCars = 100
	counter := 0

	wg.Add(raceCars)

	for i := 0; i < raceCars; i++ {
		go func() {
			raceCar := counter
			runtime.Gosched()

			raceCar++
			counter = raceCar
			fmt.Println("Number of go routines:", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Number of go routines:", runtime.NumGoroutine())
	fmt.Println("Counter is", counter)
}

func mutexVsAliens() {
	var mutex sync.Mutex

	const raceCars = 100
	counter := 0

	wg.Add(raceCars)

	for i := 0; i < raceCars; i++ {
		go func() {
			mutex.Lock()
			raceCar := counter
			runtime.Gosched()

			raceCar++
			counter = raceCar
			fmt.Println("Number of go routines:", runtime.NumGoroutine())
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Number of go routines:", runtime.NumGoroutine())
	fmt.Println("Counter is", counter)
}

func atomicPackageUser() {

	const raceCars = 100
	var counter int64

	wg.Add(raceCars)

	for i := 0; i < raceCars; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter is:", atomic.LoadInt64(&counter))
			// runtime.Gosched()
			fmt.Println("Number of go routines:", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Number of go routines:", runtime.NumGoroutine())
	fmt.Println("Counter is", counter)
}

func useChannels() {
	c := make(chan int)

	go send(c)
	receive(c)

	fmt.Println("About to exit")
}

func send(c chan<- int) {
	for i := 14; i <= 113; i *= 2 {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
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
	// shoppingList := getGroceryList()

	// fmt.Println("Shopping list")
	// fmt.Println("--------------------------")
	// for i := 0; i < len(shoppingList); i++ {
	// 	fmt.Println(shoppingList[i])
	// }
	// whyDidTheSeniorEngineerQuit()
	// mapTheDucks()
	useChannels()
}
