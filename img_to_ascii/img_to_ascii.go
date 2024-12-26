package img_to_ascii

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const asciiChars = "@%#*+=-:. "

func pixelToASCII(pixel color.Color) string {
	r, g, b, _ := pixel.RGBA()
	gray := (r + g + b) / 3 // Average to convert RGB to grayscale
	numChars := len(asciiChars)
	asciiIndex := int((gray * uint32(numChars-1)) / 65535) // Map pixel value to ASCII index
	return string(asciiChars[asciiIndex])
}

func resizeImage(img image.Image, newWidth int) image.Image {
	oldBounds := img.Bounds()
	aspectRatio := float64(oldBounds.Dy()) / float64(oldBounds.Dx())
	newHeight := int(float64(newWidth) * aspectRatio * 0.55) // Adjust height for font aspect ratio

	newImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := x * oldBounds.Dx() / newWidth
			srcY := y * oldBounds.Dy() / newHeight
			newImage.Set(x, y, img.At(srcX, srcY))
		}
	}

	return newImage
}

func ImageToASCII(imgPath string, outputPath string, width int) error {
	// Open image file
	file, err := os.Open(imgPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	// Resize image
	resizedImg := resizeImage(img, width)

	// Convert image to ASCII
	var asciiArt strings.Builder
	bounds := resizedImg.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			asciiArt.WriteString(pixelToASCII(resizedImg.At(x, y)))
		}
		asciiArt.WriteString("\n")
	}

	// Save ASCII art to file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(asciiArt.String())
	if err != nil {
		return fmt.Errorf("failed to write ASCII art to file: %v", err)
	}

	fmt.Printf("ASCII art saved to %s\n", outputPath)
	return nil
}
