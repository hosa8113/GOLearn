package main

import "fmt"

func main() {

	fullForLoop()
	forLoopNoInitAndPostStatement()
	doWhile()
	forever()
	forEachExample()

}

func forever() {
	sum := 1
	for {
		sum += sum
		break
	}
	fmt.Printf("Sum=%v\n", sum)
}

func doWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Printf("Sum=%v\n", sum)
}

func forLoopNoInitAndPostStatement() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Printf("Sum=%v\n", sum)
}

func fullForLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("Sum=%v\n", sum)
	for i := 0; i < 5; i++ {
		// Runs 5 times, with values of step 0 through 4.
		fmt.Println("{}", i)
	}
}

func forEachExample() {
	myList := []string{"dog", "cat", "hedgehog"}

	// for {key}, {value} := range {list}
	for index, animal := range myList {
		fmt.Println("My animal is:", animal, index)
	}

	myList1 := map[string]string{
		"dog":      "woof",
		"cat":      "meow",
		"hedgehog": "sniff",
	}

	for animal, noise := range myList1 {
		fmt.Println("The", animal, "went", noise)
	}
}
