package main

import "fmt"

func main() {

	fmt.Print("struktur data")
	array()
	slice()
	mapp()
}

func array() {
	var fruits = [4]string{
		"apple",
		"banana",
		"grape",
		"melon",
	}

	fmt.Println("jumlay array \t\t", len(fruits))
	fmt.Println("buah-buahan \t", fruits)
}

func slice() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)
	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

}

func mapp() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	delete(chicken, "januari")
	fmt.Println("---- data setelah dihapus----")
	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}
}
