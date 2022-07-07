package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"sync"
	"time"
)

func load(filePath string) (grid [][]color.Color) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}

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

	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()
	png.Encode(file, img.SubImage(img.Rect))
}

func main() {
	start := time.Now()
	grid := load("img1.png")
	wg := new(sync.WaitGroup)
	for x := 0; x < len(grid); x++ {
		wg.Add(1)
		col := grid[x]
		go func() {
			for y := 0; y < len(col)/2; y++ {
				z := len(col) - y - 1
				col[y], col[z] = col[z], col[y]
			}
			defer wg.Done()
		}()
	}

	wg.Wait()
	save("ejer4recurrente.png", grid)
	elapsed := time.Since(start)
	log.Printf("tiempo de ejcucion: %s", elapsed)
}
