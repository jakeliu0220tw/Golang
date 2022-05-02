package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrWorkedHours = errors.New("invalid worked hours")
)

func main() {
	pay := payDay(81, 50)
	fmt.Println(pay)
}

func payDay(workedHours, hourlyRate int) int {
	// print out report at the end
	report := func() {
		fmt.Printf("workedHours: %d\nhourlyRate: %d\n", workedHours, hourlyRate)
		if r := recover(); r != nil {
			if r == ErrHourlyRate {
				fmt.Println("hourlyRate should between 0 and 80")
			}
			if r == ErrWorkedHours {
				fmt.Println("worked hours should be between 10 and 80")
			}
		}
	}
	defer report()

	if workedHours < 10 || workedHours > 80 {
		panic(ErrWorkedHours)
	}

	if hourlyRate < 0 || hourlyRate > 80 {
		panic(ErrWorkedHours)
	}

	wage := workedHours * hourlyRate
	return wage
}
