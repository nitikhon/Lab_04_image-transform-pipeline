package main

import (
	"fmt"
	"lab04/utils"
	"log"
	"time"
)

// source generates images and sends them to the output channel
func source(names []string) <-chan *utils.Image {
	out := make(chan *utils.Image)
	go func() {
		for _, name := range names {
			img, err := utils.LoadImage(name)
			if err != nil {
				log.Printf("Source error: %v", err)
				continue
			}

			// Simulate upload/arrival delay
			time.Sleep(50 * time.Millisecond)
			out <- img
		}
		close(out)
	}()
	return out
}

// resize accepts images from the input channel, resizes them, and sends them to the output channel
func resize(in <-chan *utils.Image) <-chan *utils.Image {
	out := make(chan *utils.Image)
	// TODO: Implement the resize stage
	// - Launch a goroutine
	// - Iterate over 'in' and process images using utils.Resize()
	// - Send to 'out'
	// - Ensure 'out' is closed when done
	return out
}

// watermark accepts images, watermarks them, and sends them to the output channel
func watermark(in <-chan *utils.Image) <-chan *utils.Image {
	out := make(chan *utils.Image)
	// TODO: Implement the watermark stage
	// - Launch a goroutine
	// - Iterate over 'in' and process images using utils.Watermark()
	// - Send to 'out'
	// - Ensure 'out' is closed when done
	return out
}

func main() {
	start := time.Now()

	imageNames := []string{
		"photo1.jpg", "photo2.jpg", "photo3.jpg", "photo4.jpg", "photo5.jpg",
		"photo6.jpg", "photo7.jpg", "photo8.jpg", "photo9.jpg", "photo10.jpg",
	}

	// 1. Source
	srcChannel := source(imageNames)

	// 2. Resize
	resizeChannel := resize(srcChannel)

	// 3. Watermark
	watermarkChannel := watermark(resizeChannel)

	// 4. Sink (Upload)
	// TODO: Consume images from watermarkChannel and call utils.SaveImage()
	_ = watermarkChannel // Placeholder to avoid unused variable error

	fmt.Printf("\nConcurrent Pipeline took: %v\n", time.Since(start))
}
