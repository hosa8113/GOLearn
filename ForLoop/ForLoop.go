package main

import "fmt"

func main() {

	fullForLoop()
	forLoopNoInitAndPostStatement()
	doWhile()
	forever()

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
}
