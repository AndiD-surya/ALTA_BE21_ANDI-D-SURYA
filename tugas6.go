package main

import "fmt"

// baju = 370.000, diskon = 15%, hasil = 314.500
func main() {
	var baju float32 = 370000
	var diskon float32 = 15 * baju / 100
	var hasil = baju - diskon
	fmt.Println(hasil)
}