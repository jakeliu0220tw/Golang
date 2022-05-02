package main

import "fmt"

func main() {
	itemsSold()
}

func itemsSold() {
	items := make(map[string]int)
	items["jake"] = 30
	items["eva"] = 60
	items["dezember"] = 120

	for k, v := range items {
		rank := ""
		if v < 40 {
			rank = "poor"
		} else if v >= 40 && v < 100 {
			rank = "good"
		} else {
			rank = "excellent"
		}
		fmt.Println(k, ":", v, ", rank", "=>", rank)
	}
}
