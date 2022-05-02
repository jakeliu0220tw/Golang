package main

import (
	"fmt"
	"time"
)

// func main() {
// 	dayBorn := time.Sunday

// 	switch dayBorn {
// 	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
// 		fmt.Println("Your birthday is in working day.")
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("Your birthday is in holiday.")
// 	default:
// 		fmt.Println("Something wrong.")
// 	}
// }

func main() {
	switch dayBorn := time.Sunday; {
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("Your birthday is in holiday.")
	default:
		fmt.Println("Your birthday is in working day.")
	}
}
