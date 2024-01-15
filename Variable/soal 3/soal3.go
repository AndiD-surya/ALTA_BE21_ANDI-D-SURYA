package main

import "fmt"
//T = 20, r = 4, phi = 3.14159265358979323846
// rumus 2 x π x r x (r + T) 
func main() {
	var T float32 = 20
	var r float32 = 4
	const phi float32 = 3.14 //159265358979323846
	var hasil1 = 2 * phi * r
	var hasil2 = r + T
	var hasil = hasil1 * hasil2
	fmt.Println(hasil)
}