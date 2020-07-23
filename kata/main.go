package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func summation(number int) int {
	return number * (number + 1) / 2
}

// * my solution

func printerError(s string) string {
	var errors int
	lens := len(s)
	for _, charactor := range s {
		matched, _ := regexp.MatchString("[^a-m]", string(charactor))
		if matched {
			errors++
		}
	}
	return strconv.Itoa(errors) + "/" + strconv.Itoa(lens)
}

/**
 * * best solution
  func PrinterError(s string) string {
	 e := 0
   for _, c := range s {
     if c >= 'a' && c <= 'm' {
       continue
     } else {
       e++
     }
   }

   return fmt.Sprintf("%d/%d", e, len(s))
}
*/

func main() {
	result := printerError("aavbbeqwer")
	fmt.Println("errors:", result)
}
