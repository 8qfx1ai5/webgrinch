package main

import (
	"fmt"
	"regexp"
)

func main() {
	// str := "[5-8s-y][C-Uafg][asdf]"
	str := "[s-yR-Ub-d]"
	// str := "5-8s-y]"
	strPure := str[1 : len(str)-1]
	digitRange := regexp.MustCompile(`\d-\d`)
	//uppCaseRange := regexp.MustCompile(`[A-Z]-[A-Z]`)
	//lowerCaseRange := regexp.MustCompile(`[a-z]-[a-z]`)
	submatchall := digitRange.FindString(strPure)

	fmt.Println(submatchall)
}
