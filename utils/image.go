package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"
)

// Image represents the data flowing through the pipeline
type Image struct {
	InputPath  string
	OutputPath string
	FileName   string
	Data       image.Image
	Steps      []string
}

// LoadImage reads the image from the input directory
func LoadImage(filename string) (*Image, error) {
	inputPath := filepath.Join("input", filename)
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	imgData, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return &Image{
		InputPath: inputPath,
		FileName:  filename,
		Data:      imgData,
		Steps:     []string{"loaded"},
	}, nil
}

// Resize scales the image down by 50%
func Resize(img *Image) {
	bounds := img.Data.Bounds()
	newWidth := bounds.Max.X / 2
	newHeight := bounds.Max.Y / 2

	newRect := image.Rect(0, 0, newWidth, newHeight)
	newImg := image.NewRGBA(newRect)

	// Simple nearest neighbor resizing
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// Source coordinates (mapped from destination)
			srcX := x * 2
			srcY := y * 2
			newImg.Set(x, y, img.Data.At(srcX, srcY))
		}
	}

	img.Data = newImg
	img.Steps = append(img.Steps, "resized")
}

// Watermark adds a small red overlay to the top-left corner
func Watermark(img *Image) {
	// Create a mutable copy of the image to draw on
	bounds := img.Data.Bounds()
	m := image.NewRGBA(bounds)
	draw.Draw(m, bounds, img.Data, bounds.Min, draw.Src)

	// Draw a red rectangle as watermark
	watermarkRect := image.Rect(10, 10, 60, 60)
	red := color.RGBA{255, 0, 0, 128} // Semi-transparent red
	draw.Draw(m, watermarkRect, &image.Uniform{red}, image.Point{}, draw.Over)

	img.Data = m
	img.Steps = append(img.Steps, "watermarked")
}

// SaveImage writes the processed image to the output directory
func SaveImage(img *Image) error {
	outputPath := filepath.Join("output", img.FileName)
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, img.Data, nil)
	if err != nil {
		return err
	}

	img.OutputPath = outputPath
	img.Steps = append(img.Steps, "saved")
	fmt.Printf("Finished processing: %s, steps: %v\n", img.FileName, img.Steps)
	return nil
}
