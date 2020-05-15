package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{255, 0, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	// for x := 0; x < 380; x++ {
	// 	y := 25
	// 	rectImg.Set(x, y, red)
	// }

	// for y := 0; y < 190; y++ {
	// 	x := 100
	// 	rectImg.Set(x, y, red)
	// }

	for i := 10; i < 190; i++ {
		rectImg.Set(i, 20, red)
		rectImg.Set(i, 180, red)
		rectImg.Set(20, i, red)
		rectImg.Set(180, i, red)
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
