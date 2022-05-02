package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr1 := []string{"empid", "employee", "address", "hours worked", "houly rate", "manager"}
	hdr2 := []string{"employee", "empid", "Houly rate", "Hours worked", "Address", "Manager"}

	csvHdrCol(hdr1)
	csvHdrCol(hdr2)
}

func csvHdrCol(headers []string) {
	HdrToColIndex := make(map[int]string)
	for i, v := range headers {
		str := strings.ToLower(strings.TrimSpace(v))
		switch str {
		case "employee":
			HdrToColIndex[i] = str
		case "address":
			HdrToColIndex[i] = str
		case "hours worked":
			HdrToColIndex[i] = str
		case "hourly rate":
			HdrToColIndex[i] = str
		}
	}
	fmt.Println(HdrToColIndex)
}
