package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math/rand"
	"os"
	"path/filepath"
)

func main() {
	inputDir := "input"
	if err := os.MkdirAll(inputDir, 0755); err != nil {
		panic(err)
	}
	// Create output dir as well
	if err := os.MkdirAll("output", 0755); err != nil {
		panic(err)
	}

	for i := 1; i <= 10; i++ {
		filename := fmt.Sprintf("photo%d.jpg", i)
		createRandomImage(filepath.Join(inputDir, filename))
		fmt.Printf("Created %s\n", filename)
	}
}

func createRandomImage(path string) {
	width, height := 800, 600
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Fill background
	bgColor := color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	// Draw some random rects
	for i := 0; i < 5; i++ {
		rWidth := rand.Intn(200) + 50
		rHeight := rand.Intn(200) + 50
		x := rand.Intn(width - rWidth)
		y := rand.Intn(height - rHeight)

		rectColor := color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		}

		draw.Draw(img, image.Rect(x, y, x+rWidth, y+rHeight), &image.Uniform{rectColor}, image.Point{}, draw.Src)
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, img, nil)
}
