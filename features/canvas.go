package features

import (
	"strings"
	"os"
	"fmt"
	"math"
)

type Canvas struct {
	width int
	height int
	pixels [][]Color
}

func NewCanvas(width, height int) *Canvas {
	c := &Canvas{
		width: width,
		height: height,
		pixels: make([][]Color, height),
	}
	for i := range c.pixels {
		c.pixels[i] = make([]Color, width)
	}

	return c
}

func WritePixel(canvas Canvas, x, y int, color Color) {
	canvas.pixels[y][x] = color
}

func PixelAt(canvas Canvas, x, y int) Color {
	return canvas.pixels[y][x]
}

func clamp(v float64) int {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 255
	}
	return int(math.Ceil(v * 255))
}

func CanvasToPPM(c Canvas, filename string) error {
	var builder strings.Builder
	fmt.Fprintf(&builder, "P3\n%d %d\n255\n", c.width, c.height)
	for y := 0; y < c.height; y++ {
		var currentLine string
		lineLength := 0
		for x := 0; x < c.width; x++ {
			color := PixelAt(c, x, y)
			r := clamp(color.red)
			g := clamp(color.green)
			b := clamp(color.blue)
			pixelStr := fmt.Sprintf("%d %d %d", r, g, b)

			if lineLength + len(pixelStr) > 70 {
				if pixelStr1 := fmt.Sprintf("%d %d", r, g); lineLength + len(pixelStr1) <= 70 {
					builder.WriteString(currentLine + pixelStr1 + "\n")
					currentLine = fmt.Sprintf("%d ", b)
					lineLength = len(currentLine)
				} else if pixelStr3 := fmt.Sprintf("%d", r); lineLength + len(pixelStr3) + 1 <= 70 {
					builder.WriteString(currentLine + pixelStr3 + "\n")
					currentLine = fmt.Sprintf("%d %d ", g, b)
					lineLength = len(currentLine)
				} else {
					builder.WriteString(strings.TrimSpace(currentLine) + "\n")
					currentLine += pixelStr + " "
					lineLength = len(pixelStr) + 1
				}
			} else {
				currentLine += pixelStr + " "
				lineLength += len(pixelStr) + 1
			}
		}
		builder.WriteString(strings.TrimSpace(currentLine) + "\n")
	}

	err := os.WriteFile(filename, []byte(builder.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write PPM file: %w", err)
	}
	return nil
}

