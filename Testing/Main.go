package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	fmt.Println("start")
	i := 5
	j := 10
	fmt.Println(i, "+", j, "=", Add(i, j))
	fmt.Println(i, "*", j, "=", Multiply(i, j))
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)

}

func Multiply(i int, i2 int) int {
	return i * i2
}

func Add(i int, i2 int) int {
	return i + i2
}

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
