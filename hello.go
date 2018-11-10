package main

import "fmt"

var nama string = "Khansa"
var angkapositip uint8 = 27
var decimalNumber float32


func main() {
	fmt.Println("hello world")
	nama = "wnr"
	namaDepan := "elko"
	angkaNegatip := -123
	decimalNumber = 2.712
	fmt.Println("hello ", nama)
	fmt.Println("hello ", namaDepan)
	fmt.Printf("angka positip : %d \n ", angkapositip)
	fmt.Printf("angka negatip : %d \n", angkaNegatip)
	fmt.Printf("angka decimal : %f \n", decimalNumber)
	fmt.Printf("angka decimal : %.3f \n", decimalNumber)

}
