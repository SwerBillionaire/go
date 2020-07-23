package main

import (
	"fmt"
)

func countSheep(sheepCount int) string {
	var result string
	for i := 1; i <= sheepCount; i++ {
		result += fmt.Sprintf("%d Sheep...", i)
	}
	return result
}

func main() {
	fmt.Println(countSheep(4))
}
