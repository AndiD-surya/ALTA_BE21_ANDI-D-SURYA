package main

import "fmt"

func main () {
	studentScore := 80

	if studentScore >= 80 && studentScore <= 100{
		fmt.Println("A")
	}
	if studentScore >= 65 && studentScore <= 79 {
		fmt.Println("B+")
	}
	if studentScore >= 50 && studentScore <= 64{
		fmt.Println("B")
	}
	if studentScore >= 35 && studentScore <= 49{
		fmt.Println("C")
	}
	if studentScore >= 0 && studentScore <=35 {
		fmt.Println ("D")
	}
}