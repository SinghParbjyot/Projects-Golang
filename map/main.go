package main

import "fmt"

// Differents type to declare a map
func main() {
	//var colors map[string]string
	colors := make(map[string]string)
	/*colors := map[string]string{
		"red":   "#FF0000",
		"green": "#3eff12",
	}
	*/
	colors["white"] = "#fffff"
	colors["blue"] = "#23b3f1"
	colors["green"] = "#23f134"
	colors["purple"] = "#ad24f2"
	printMap(colors)

}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is "+hex)
	}
}
