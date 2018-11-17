package main

import "fmt"

func main() {

	operator()
	seleksi()
	perulangan()
}

func operator() {
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)

}

func seleksi() {
	var nilai = 8

	if nilai > 8 {
		fmt.Print("luar biasa \n")
	} else if nilai > 6 {
		fmt.Print("lulus \n")
	} else {
		fmt.Print("tidak lulus \n")
	}

	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
	var skor = 6

	switch skor {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

}

func perulangan() {
	for i := 0; i < 5; i++ {
		fmt.Println("angaka : ", i)
	}

	i := 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
