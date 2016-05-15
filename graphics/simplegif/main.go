package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"

	"github.com/klrkdekira/algorithms/helper"
	"github.com/klrkdekira/algorithms/sort/bubble"
)

var (
	pad      = 10
	rect     = image.Rect(0, 0, 320, 320)
	palettes = []color.Color{
		color.RGBA{255, 255, 255, 255},
		color.RGBA{0, 0, 0, 255},
	}
)

func main() {
	a := helper.RandIntN(20)

	file, err := os.Create("test.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err := gif.EncodeAll(file, animateSort(a)); err != nil {
		log.Fatal(err)
	}
}

func animateSort(a []int) *gif.GIF {
	images := []*image.Paletted{}
	delays := []int{}

	xSize := (rect.Max.X - pad*2) / len(a)
	// cheated, max value == max sample size
	ySize := float64((rect.Max.Y - (pad * 2))) / float64(len(a))

	c := make(chan []int)
	done := make(chan struct{})
	go func() {
		for {
			cur, more := <-c
			if !more {
				break
			}
			img := image.NewPaletted(rect, palettes)
			images = append(images, img)
			delays = append(delays, 0)
			for p, v := range cur {
				start := rect.Min.X + pad + (xSize * p)
				end := rect.Min.X + pad + (xSize * (p + 1))
				for x := start; x < end; x++ {
					for y := rect.Max.Y - pad; y > rect.Max.Y-pad-int(ySize*float64(v)); y-- {
						img.Set(x, y, palettes[1])
					}
				}
			}

		}
		close(done)
	}()
	bubble.SortAnimation(a, c)

	<-done
	return &gif.GIF{
		Image: images,
		Delay: delays,
	}
}
