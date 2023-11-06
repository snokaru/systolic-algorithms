package main

import (
	"fmt"
	"os"
	"strconv"
	"image"
	"image/color"
	"image/png"
)


func draw_transitions(transitions [][]uint8, filename string) {
	width := len(transitions[0])
	height := len(transitions)
	fmt.Printf("making a %dx%d image\n", width, height)

	topLeft := image.Point{0, 0}
	bottomRight := image.Point{width, height}
	
	img := image.NewRGBA(image.Rectangle{topLeft, bottomRight})
	
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if transitions[x][y] == 0 {
				img.Set(y, x, color.Black)
			} else if transitions[x][y] == 1 {
				img.Set(y, x, color.White)
			} else {
				fmt.Printf("something bad is happening: %d\n", transitions[x][y])
			}
		}
	}

	f, _ := os.Create(filename)
	png.Encode(f, img)
}


func neighborhood_to_number(left uint8, mid uint8, right uint8) uint8 {
	return left * 4 + mid * 2 + right
}


func apply_transition_rule(cells []uint8, rule uint8) []uint8 {
	transition := make([]uint8, len(cells))

	for i := 0; i < len(cells); i++ {
		var left uint8 = 0
		var mid uint8 = cells[i]
		var right uint8 = 0

		if i - 1 >= 0 {
			left = cells[i-1]
		}
		if i + 1 < len(cells) {
			right = cells[i+1]	
		}
		neigh := neighborhood_to_number(left, mid, right)

		transition[i] = (rule >> neigh) & 0x01
	}
	return transition
}


func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: transition rule required\n", os.Args[0])
		os.Exit(1)
	}

	rule, err := strconv.Atoi(os.Args[1])
	if err != nil || rule < 0 || rule > 255 {
		fmt.Fprintf(os.Stderr, "%s: transition rule must be an integer between 0 and 255\n", os.Args[0])
		os.Exit(2)
	}
	transition_rule := uint8(rule)

	height := 100
	if len(os.Args) == 3 {
		h, err := strconv.Atoi(os.Args[2])
		if err == nil {
			height = h
		}
	}
	transitions := make([][]uint8, height)

	width := 100

	start := make([]uint8, width)
	start[width / 2] = 1
	transitions[0] = start

	for i := 1; i < len(transitions); i++ {
		transition := apply_transition_rule(transitions[i - 1], transition_rule)
		transitions[i] = transition
	}

	draw_transitions(transitions, "result.png")
}
