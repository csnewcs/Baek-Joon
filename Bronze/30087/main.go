package main

import "fmt"

func main() {
	var seminaClasses = map[string]string{
		"Algorithm":              "204",
		"DataAnalysis":           "207",
		"ArtificialIntelligence": "302",
		"CyberSecurity":          "B101",
		"Network":                "303",
		"Startup":                "501",
		"TestStrategy":           "105",
	}
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var semina string
		fmt.Scanln(&semina)
		fmt.Println(seminaClasses[semina])
	}
}
