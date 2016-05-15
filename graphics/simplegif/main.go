package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"

	"github.com/klrkdekira/algorithms/helper"
	"github.com/klrkdekira/algorithms/sort/bubble"
	"github.com/klrkdekira/algorithms/sort/insertion"
)

var (
	pad      = 10
	rect     = image.Rect(0, 0, 320, 320)
	palettes = []color.Color{
		color.RGBA{255, 255, 255, 255},
		color.RGBA{0, 0, 0, 255},
		color.RGBA{255, 0, 0, 255},
	}
	sampleSize = 20
)

func main() {
	var a []int

	a = helper.RandIntN(sampleSize)
	bubbleSort := &bubble.Bubble{}
	animateSort(a, "bubble", bubbleSort)

	a = helper.RandIntN(sampleSize)
	insertionSort := &insertion.Insertion{}
	animateSort(a, "insertion", insertionSort)
}

func animateSort(a []int, name string, sorter helper.Sorter) {
	images := []*image.Paletted{}
	delays := []int{}

	xSize := (rect.Max.X - pad*2) / len(a)
	// cheated, max value == max sample size
	ySize := float64((rect.Max.Y - (pad * 2))) / float64(len(a))

	c := make(chan *helper.Progress)
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
			for p, v := range cur.A {
				start := rect.Min.X + pad + (xSize * p)
				end := rect.Min.X + pad + (xSize * (p + 1))
				for x := start; x < end; x++ {
					for y := rect.Max.Y - pad; y > rect.Max.Y-pad-int(ySize*float64(v)); y-- {
						if p == cur.I || p == cur.J {
							img.Set(x, y, palettes[2])
						} else {
							img.Set(x, y, palettes[1])
						}
					}
				}
			}

		}
		close(done)
	}()
	sorter.SortAnimation(a, c)

	<-done

	file, err := os.Create(fmt.Sprintf("%s.gif", name))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err := gif.EncodeAll(file, &gif.GIF{
		Image: images,
		Delay: delays,
	}); err != nil {
		log.Fatal(err)
	}
}

// func animateBubbleSort(a []int) {
// 	images := []*image.Paletted{}
// 	delays := []int{}

// 	xSize := (rect.Max.X - pad*2) / len(a)
// 	// cheated, max value == max sample size
// 	ySize := float64((rect.Max.Y - (pad * 2))) / float64(len(a))

// 	c := make(chan *helper.Progress)
// 	done := make(chan struct{})
// 	go func() {
// 		for {
// 			cur, more := <-c
// 			if !more {
// 				break
// 			}
// 			img := image.NewPaletted(rect, palettes)
// 			images = append(images, img)
// 			delays = append(delays, 0)
// 			for p, v := range cur.A {
// 				start := rect.Min.X + pad + (xSize * p)
// 				end := rect.Min.X + pad + (xSize * (p + 1))
// 				for x := start; x < end; x++ {
// 					for y := rect.Max.Y - pad; y > rect.Max.Y-pad-int(ySize*float64(v)); y-- {
// 						if p == cur.I || p == cur.J {
// 							img.Set(x, y, palettes[2])
// 						} else {
// 							img.Set(x, y, palettes[1])
// 						}
// 					}
// 				}
// 			}

// 		}
// 		close(done)
// 	}()
// 	bubble.SortAnimation(a, c)

// 	<-done

// 	file, err := os.Create("bubble.gif")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	if err := gif.EncodeAll(file, &gif.GIF{
// 		Image: images,
// 		Delay: delays,
// 	}); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func animateInsertionSort(a []int) {
// 	images := []*image.Paletted{}
// 	delays := []int{}

// 	xSize := (rect.Max.X - pad*2) / len(a)
// 	// cheated, max value == max sample size
// 	ySize := float64((rect.Max.Y - (pad * 2))) / float64(len(a))

// 	c := make(chan *helper.Progress)
// 	done := make(chan struct{})
// 	go func() {
// 		for {
// 			cur, more := <-c
// 			if !more {
// 				break
// 			}
// 			img := image.NewPaletted(rect, palettes)
// 			images = append(images, img)
// 			delays = append(delays, 0)
// 			for p, v := range cur.A {
// 				start := rect.Min.X + pad + (xSize * p)
// 				end := rect.Min.X + pad + (xSize * (p + 1))
// 				for x := start; x < end; x++ {
// 					for y := rect.Max.Y - pad; y > rect.Max.Y-pad-int(ySize*float64(v)); y-- {
// 						if p == cur.I || p == cur.J {
// 							img.Set(x, y, palettes[2])
// 						} else {
// 							img.Set(x, y, palettes[1])
// 						}
// 					}
// 				}
// 			}

// 		}
// 		close(done)
// 	}()
// 	insertion.SortAnimation(a, c)

// 	<-done

// 	file, err := os.Create("insertion.gif")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	if err := gif.EncodeAll(file, &gif.GIF{
// 		Image: images,
// 		Delay: delays,
// 	}); err != nil {
// 		log.Fatal(err)
// 	}
// }
