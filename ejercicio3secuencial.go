package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	txt, err := os.Create("redcolorsec.txt")
	if err != nil {
		fmt.Printf("error", err)
		return
	}
	txt1, err := os.Create("greencolorsec.txt")
	if err != nil {
		fmt.Printf("error", err)
		return
	}
	txt2, err := os.Create("bluecolorsec.txt")
	if err != nil {
		fmt.Printf("error", err)
		return
	}
	imgPath := "img1.jpg"
	f, err := os.Open(imgPath)
	check(err)

	img, _, err := image.Decode(f)

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	start := time.Now()
	// loop though all the x

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			pixel := img.At(x, y)

			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			// Offset colors a little, adjust it to your taste
			r := float64(originalColor.R)
			g := float64(originalColor.G)
			b := float64(originalColor.B)

			_, err = txt.WriteString(fmt.Sprintf("%d\n", uint8(r)))
			if err != nil {
				fmt.Printf("error", err)
			}
			_, err = txt1.WriteString(fmt.Sprintf("%d\n", uint8(g)))
			if err != nil {
				fmt.Printf("error", err)
			}
			_, err = txt2.WriteString(fmt.Sprintf("%d\n", uint8(b)))
			if err != nil {
				fmt.Printf("error", err)
			}
		}
	}

	elapsed := time.Since(start)
	log.Printf("tiempo de ejcucion: %s", elapsed)

	ext := filepath.Ext(imgPath)
	name := strings.TrimSuffix(filepath.Base(imgPath), ext)
	newImagePath := fmt.Sprintf("%s/%s_ejercicio3sec%s", filepath.Dir(imgPath), name, ext)
	fg, err := os.Create(newImagePath)
	defer fg.Close()
	check(err)
	err = jpeg.Encode(fg, wImg, nil)
	check(err)
}
