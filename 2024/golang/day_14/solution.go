package day_14

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

func Solution() {
	content, err := os.ReadFile("./day_14/input.txt")
	if err != nil {
		panic("unable to open day 14 input")
	}

	strContent := string(content)
	answer1 := partOneSolution(strContent, 103, 101)
	fmt.Printf("Day 14 - Part 1: %d \n", answer1)
	answer2 := partTwoSolution(strContent)
	fmt.Printf("Day 14 - Part 2: %d \n", answer2)
}

func partOneSolution(input string, height, width int) int64 {
	sanatizedInupt := strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(sanatizedInupt, "\n")
	guards := []*guard{}
	for _, l := range lines {
		var x, y, moveX, moveY int
		_, err := fmt.Sscanf(l, "p=%d,%d v=%d,%d", &x, &y, &moveX, &moveY)
		if err != nil {
			panic(err)
		}
		guards = append(guards, &guard{
			currentX: x,
			currentY: y,
			moveX:    moveX,
			moveY:    moveY,
		})
	}

	//I know this can e optimized by stepping *100 and doing a modulus I just don't want to figure that out right now.
	for i := 0; i < 100; i++ {
		for _, g := range guards {
			g.Step(width, height)
		}
	}

	var topLeft, bottomLeft, topRight, bottomRight []*guard
	midX, midY := (width / 2), (height / 2)
	for _, g := range guards {
		if g.currentX < midX {
			if g.currentY < midY {
				topLeft = append(topLeft, g)
			} else if g.currentY > midY {
				bottomLeft = append(bottomLeft, g)
			}
		} else if g.currentX > midX {
			if g.currentY < midY {
				topRight = append(topRight, g)
			} else if g.currentY > midY {
				bottomRight = append(bottomRight, g)
			}
		}
	}

	return int64(len(topLeft) * len(bottomLeft) * len(topRight) * len(bottomRight))
}

func partTwoSolution(input string) int64 {
	sanatizedInupt := strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(sanatizedInupt, "\n")
	guards := []*guard{}
	for _, l := range lines {
		var x, y, moveX, moveY int
		_, err := fmt.Sscanf(l, "p=%d,%d v=%d,%d", &x, &y, &moveX, &moveY)
		if err != nil {
			panic(err)
		}
		guards = append(guards, &guard{
			currentX: x,
			currentY: y,
			moveX:    moveX,
			moveY:    moveY,
		})
	}

	preIter := 400
	for i := 0; i < preIter; i++ {
		for _, g := range guards {
			g.Step(101, 103)
		}
	}
	for i := 0; i < 10_000; i++ {
		for _, g := range guards {
			g.Step(101, 103)
		}
		imageGuards(guards, 101, 103, preIter+i)
	}

	//Real answer 7584, not computed, looked at images in file explorer
	return 42
}

func imageGuards(guards []*guard, width, height int, iteration int) {
	arr := make([][]uint8, height)
	for i := range arr {
		arr[i] = make([]uint8, width)
	}

	for _, g := range guards {
		arr[g.currentY][g.currentX] += 1
	}

	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	for y, r := range arr {
		for x, c := range r {
			if c != 0 {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	fileName := fmt.Sprintf("image_%d.png", iteration)
	f, _ := os.Create(fileName)
	png.Encode(f, img)
	f.Close()
}

func printGuards(guards []*guard, width, height int, iteration int) {
	arr := make([][]uint8, height)
	for i := range arr {
		arr[i] = make([]uint8, width)
	}

	for _, g := range guards {
		arr[g.currentY][g.currentX] += 1
	}
	buff := bytes.Buffer{}

	fmt.Fprintln(&buff, "------------------")
	fmt.Fprintf(&buff, "PRINTING GUARDS FOR ITERATION %d\n", iteration)
	fmt.Fprintln(&buff, "------------------")
	for _, r := range arr {
		for _, c := range r {
			if c == 0 {
				fmt.Fprint(&buff, "0")
			} else {
				fmt.Fprint(&buff, c)
			}
		}
		fmt.Fprint(&buff, "\n")
	}
	fmt.Fprintln(&buff, "------------------")
	fmt.Fprint(&buff, "\n")

	fo, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fo.Write(buff.Bytes())
	fo.Close()
}

type guard struct {
	currentX, currentY int
	moveX, moveY       int
}

func (g *guard) Step(width, height int) {
	g.currentX += g.moveX
	g.currentY += g.moveY
	if g.currentX < 0 {
		g.currentX = width + g.currentX
	}
	if g.currentX >= width {
		g.currentX = g.currentX - width
	}
	if g.currentY < 0 {
		g.currentY = height + g.currentY
	}
	if g.currentY >= height {
		g.currentY = g.currentY - height
	}
}
