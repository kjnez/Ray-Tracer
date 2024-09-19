package tuples

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Tuple struct {
	x, y, z, w float64
}

type Color struct {
	red, green, blue float64
}

type Canvas struct {
	width int
	height int
	pixels [][]Color
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func NewColor(red, green, blue float64) Color {
	return Color{red, green, blue}
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

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func (t Tuple) IsPoint() bool {
	return t.w == 1.0
}

func (t Tuple) IsVector() bool {
	return t.w == 0.0
}

const EPSILON = 1e-5

func Equal(a, b float64) bool {
	if math.Abs(a - b) < EPSILON {
		return true
	} else {
		return false
	}
}

func Equals(t1, t2 Tuple) bool {
	return Equal(t1.x, t2.x) &&
	Equal(t1.y, t2.y) &&
	Equal(t1.z, t2.z) &&
	Equal(t1.w, t2.w)
}

func EqualsColor(c1, c2 Color) bool {
	return Equal(c1.red, c2.red) &&
	Equal(c1.green, c2.green) &&
	Equal(c1.blue, c2.blue)
}

func Add(t1, t2 Tuple) Tuple {
	return Tuple{
		t1.x + t2.x,
		t1.y + t2.y,
		t1.z + t2.z,
		t1.w + t2.w,
	}
}

func Subtract(t1, t2 Tuple) Tuple {
	return Tuple{
		t1.x - t2.x,
		t1.y - t2.y,
		t1.z - t2.z,
		t1.w - t2.w,
	}
}

func Negate(t Tuple) Tuple {
	return Tuple{-t.x, -t.y, -t.z, -t.w}
}

func Multiply(tuple Tuple, scalar float64) Tuple {
	return Tuple{
		tuple.x * scalar,
		tuple.y * scalar,
		tuple.z * scalar,
		tuple.w * scalar,
	}
}

func Divide(tuple Tuple, scalar float64) Tuple {
	if scalar == 0 {
		panic("division by zero")
	}
	return Multiply(tuple, 1.0 / scalar)
}

func Magnitude(t Tuple) float64 {
	if t.IsVector() {
		return math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2))
	} else {
		panic("tuple must be a vector")
	}
}

func Normalize(t Tuple) Tuple {
	if t.IsVector() {
		magnitude := Magnitude(t)
		return Divide(t, magnitude)
	} else {
		panic("cannot normalize a point")
	}
}

func DotProduct(t1, t2 Tuple) float64 {
	if !(t1.IsVector() && t2.IsVector()) {
		panic("cannot perform dot product on points")
	} else {
		return t1.x * t2.x + t1.y * t2.y + t1.z * t2.z + t1.w * t2.w
	}
}

func CrossProduct(t1, t2 Tuple) Tuple {
	if !(t1.IsVector() && t2.IsVector()) {
		panic("cannot perform cross product on points")
	} else {
		return Vector(
			t1.y * t2.z - t1.z * t2.y,
			t1.z * t2.x - t1.x * t2.z,
			t1.x * t2.y - t1.y * t2.x,
		)
	}
}

func AddColor(c1, c2 Color) Color {
	return Color{
		c1.red + c2.red,
		c1.green + c2.green,
		c1.blue + c2.blue,
	}
}

func SubtractColor(c1, c2 Color) Color {
	return Color{
		c1.red - c2.red,
		c1.green - c2.green,
		c1.blue - c2.blue,
	}
}

func MultiplyColorByScalar(c Color, s float64) Color {
	return Color{
		c.red * s,
		c.green * s,
		c.blue * s,
	}
}

func HadamardProduct(c1, c2 Color) Color {
	return Color{
		c1.red * c2.red,
		c1.green * c2.green,
		c1.blue * c2.blue,
	}
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
