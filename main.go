package main

import (
	"fmt"
	"img_to_ascii/img_to_ascii"
)

func main() {
	imagePath := "gnome.png"
	outputPath := "ascii_art.txt"
	width := 100

	err := img_to_ascii.ImageToASCII(imagePath, outputPath, width)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
