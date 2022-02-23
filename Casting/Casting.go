package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	fromIntToString()
	fromStringToInt()
	fromFloatToString()

	fromIntToInt()
	fromFloatToInt()
}

func fromFloatToInt() {

	var x int64 = 5
	fmt.Println(fmt.Sprintf("x(%s) = %d", reflect.TypeOf(x), x))
	y := float64(x)
	fmt.Println(fmt.Sprintf("y(%s) = %.2f", reflect.TypeOf(y), y))
	z := float32(x)
	fmt.Println(fmt.Sprintf("z(%s) = %f", reflect.TypeOf(z), z))

}

func fromIntToInt() {

	s := "2147483647" // biggest int32
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	i := int32(i64)
	fmt.Println(fmt.Sprintf("String -> Int64 -> Int32 = %d", i))

}

func fromIntToString() {
	i := 4711
	var s string
	s = strconv.Itoa(i)
	fmt.Println(s)

	var i64 int64
	i64 = 3306
	s = strconv.FormatInt(i64, 10)
	fmt.Println(s)
}

func fromStringToInt() {
	s := "4711"
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("fromStringToInt -> %d", i))
}

func fromFloatToString() {
	var pi float64 = 3.1415926535
	var s string

	s = strconv.FormatFloat(pi, 'E', -1, 32)
	fmt.Println(s)
}
