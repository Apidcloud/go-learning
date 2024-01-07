package main

import "fmt"

type HexColors map[string]string

func main() {
	colors := HexColors{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"
	colors["green"] = "replaced green"

	updateKey(colors)
	printMap(colors)
}

// maps pass by reference, just like slices,
// channels, pointers, and functions
func updateKey(c HexColors) {
	c["red"] = "xpto"
}

func printMap(c HexColors) {
	for key, value := range c {
		fmt.Printf("Hex color of %v is %v\n", key, value)
	}
}
