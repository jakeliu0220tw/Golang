package main

import (
	"errors"
	"fmt"
)

var ErrHourlyRate = errors.New("invalid hourly rate")
var ErrHoursWorked = errors.New("invalid worked hours")

func payDay(hoursWorked int, hoursRate int) (int, error) {
	if hoursRate < 10 || hoursRate > 75 {
		// return 0, ErrHourlyRate
		return 0, fmt.Errorf("invalid HourlyRate! err=%w", ErrHourlyRate)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, ErrHoursWorked
	}

	if hoursWorked > 40 {
		wage := 40*hoursRate + (hoursWorked-40)*hoursRate*2
		return wage, nil
	}

	wage := hoursWorked * hoursRate
	return wage, nil
}

func main() {
	pay, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	pay, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pay)

	pay, err = payDay(8, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	pay, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)
}
