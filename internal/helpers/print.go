package helpers

import (
	"fmt"
	"os"
)

// TODO:turn method to enum
func PretyPrint(str, color, method string) {
	Reset := "\033[0m"

	colorMap := map[string]map[string]string{
		// colors made with chatGPT take this to considreation
		"red": {
			"i3":   "#bf616a",
			"term": "\033[31m",
		},
		"green": {
			"i3":   "#a3be8c",
			"term": "\033[32m",
		},
		"yellow": {
			"i3":   "#ebcb8b",
			"term": "\033[33m",
		},
		"blue": {
			"i3":   "#81a1c1",
			"term": "\033[34m",
		},
		"magenta": {
			"i3":   "#b48ead",
			"term": "\033[35m",
		},
		"cyan": {
			"i3":   "#88c0d0",
			"term": "\033[36m",
		},
		"white": {
			"i3":   "#e5e9f0",
			"term": "\033[37m",
		},
		"black": {
			"i3":   "#2e3440",
			"term": "\033[30m",
		},
	}

	selectedColor, exists := colorMap[color]
	if !exists {
		fmt.Println("Invalid color specified")
		os.Exit(1)
	}

	var chosenColor string
	switch method {
	case "i3":
		chosenColor = selectedColor["i3"]
		// i3 format is to print output then empty line then line contain color
		fmt.Println(str + "\n\n" + chosenColor)
		os.Exit(0)
	case "term":
		chosenColor = selectedColor["term"]
		fmt.Println(chosenColor + str + Reset)
		os.Exit(0)
	case "none":
		fmt.Println(str)
		os.Exit(0)
	default:
		fmt.Println("Invalid method specified")
		os.Exit(1)
	}
}
