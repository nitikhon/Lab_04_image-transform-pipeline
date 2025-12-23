package main

import (
	"fmt"
	"lab04/utils"
	"log"
	"time"
)

func main() {
	start := time.Now()

	imageNames := []string{
		"photo1.jpg", "photo2.jpg", "photo3.jpg", "photo4.jpg", "photo5.jpg",
		"photo6.jpg", "photo7.jpg", "photo8.jpg", "photo9.jpg", "photo10.jpg",
	}

	for _, name := range imageNames {
		// 1. Load Image
		img, err := utils.LoadImage(name)
		if err != nil {
			log.Printf("Failed to load %s: %v", name, err)
			continue
		}

		// 2. Pipeline steps
		utils.Resize(img)
		utils.Watermark(img)

		// 3. Save Image
		if err := utils.SaveImage(img); err != nil {
			log.Printf("Failed to save %s: %v", name, err)
		}
	}

	fmt.Printf("\nSequential Pipeline took: %v\n", time.Since(start))
}
