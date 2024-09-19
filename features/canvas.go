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

func NewCanvas(height, width int) *Canvas {
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

func WritePixel(canvas Canvas, r, c int, color Color) {
	canvas.pixels[r][c] = color
}

func PixelAt(canvas Canvas, r, c int) Color {
	return canvas.pixels[r][c]
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
	for x := 0; x < c.height; x++ {
		var currentLine string
		lineLength := 0
		for y := 0; y < c.width; y++ {
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

