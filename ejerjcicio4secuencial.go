package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"time"
)

func load(filePath string) (grid [][]color.Color) {
	// open the file and decode the contents into an image
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}
	// create and return a grid of pixels
	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		grid = append(grid, y)
	}
	return
}

func save(filePath string, grid [][]color.Color) {
	// create an image and set the pixels using the grid
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}
	// create a file and encode the image into it
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()
	png.Encode(file, img.SubImage(img.Rect))
}

func flip(grid [][]color.Color) {
	for x := 0; x < len(grid); x++ {
		col := grid[x]
		for y := 0; y < len(col)/2; y++ {
			z := len(col) - y - 1
			col[y], col[z] = col[z], col[y]
		}
	}
}
func main() {
	start := time.Now()
	grid := load("img1.png")
	flip(grid)
	save("img1secuencial", grid)
	elapsed := time.Since(start)
	log.Printf("tiempo de ejcucion: %s", elapsed)
}
