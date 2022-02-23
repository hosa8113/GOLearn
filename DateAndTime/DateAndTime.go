package main

import (
	"fmt"
	"time"
)

func main() {

	// as Go uses following constants to format date, refer https://golang.org/src/time/format.go

	/*
		const (
			stdLongMonth             = iota + stdNeedDate  // "January"
			stdMonth                                       // "Jan"
			stdNumMonth                                    // "1"
			stdZeroMonth                                   // "01"
			stdLongWeekDay                                 // "Monday"
			stdWeekDay                                     // "Mon"
			stdDay                                         // "2"
			stdUnderDay                                    // "_2"
			stdZeroDay                                     // "02"
			stdUnderYearDay                                // "__2"
			stdZeroYearDay                                 // "002"
			stdHour                  = iota + stdNeedClock // "15"
			stdHour12                                      // "3"
			stdZeroHour12                                  // "03"
			stdMinute                                      // "4"
			stdZeroMinute                                  // "04"
			stdSecond                                      // "5"
			stdZeroSecond                                  // "05"
			stdLongYear              = iota + stdNeedDate  // "2006"
			stdYear                                        // "06"
			stdPM                    = iota + stdNeedClock // "PM"
			stdpm                                          // "pm"
			stdTZ                    = iota                // "MST"
			stdFracSecond0                                 // ".0", ".00", ... , trailing zeros included
			stdFracSecond9                                 // ".9", ".99", ..., trailing zeros omitted

	*/

	fmt.Println(time.Now())
	fmt.Println(time.Now().String())
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().Location())
	fmt.Println(fmt.Sprintf("Second %v", time.Now().Second()))
	fmt.Println(fmt.Sprintf("Minute %v", time.Now().Minute()))
	fmt.Println(fmt.Sprintf("Hour %v", time.Now().Hour()))
	fmt.Println(fmt.Sprintf("Day %v", time.Now().Day()))
	fmt.Println(fmt.Sprintf("Month %v", time.Now().Month()))
	fmt.Println(fmt.Sprintf("Year %v", time.Now().Year()))
	fmt.Println(fmt.Sprintf("YearDay %v", time.Now().YearDay()))

	year, month, day := time.Now().Date()
	fmt.Println(fmt.Sprintf("Date %v. %v %v", day, month, year))

	fmt.Println(fmt.Sprintf("Iso-Date %v", time.Now().Format("2006-01-02")))

	fmt.Println(fmt.Sprintf("Date (yyyy.mm.dd-hh:mm:ss) %v", time.Now().Format("2006.01.02-15:04:05")))

	fmt.Println(fmt.Sprintf("Time (hh:mm:ss) %v", time.Now().Format("15:04:05")))
	fmt.Println(fmt.Sprintf("Time (hh:mm:ss) %v", time.Now().Format("03:04:05 PM")))

	fmt.Println(time.Parse("2006-01-02T15:04:05.999", "2021-10-26T15:01:47.687"))
}
