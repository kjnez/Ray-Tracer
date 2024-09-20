package features

import (
	"testing"
	"os"
	"strings"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	if !(c.width == 10 && c.height == 20) {
		t.Errorf("Expected height %v, got %v", 20, c.height)
		t.Errorf("Expected width %v, got %v", 10, c.width)
	}
	black := NewColor(0, 0, 0)
	if !EqualsColor(c.pixels[3][4], black) {
		t.Errorf("Expected color %v, got %v", black, c.pixels[3][4])
	}
}

func TestWritePixel(t *testing.T) {
	c := *NewCanvas(10, 20)
	red := NewColor(1, 0, 0)
	WritePixel(c, 2, 3, red)
	if !EqualsColor(red, c.pixels[3][2]) {
		t.Errorf("Expected color %v, got %v", red, c.pixels[3][2])
	}
}

func TestPixelAt(t *testing.T) {
	c := *NewCanvas(10, 20)
	red := NewColor(1, 0, 0)
	WritePixel(c, 2, 3, red)
	if PixelAt(c, 2, 3) != red {
		t.Errorf("Expected color %v, got %v", red, PixelAt(c, 2, 3))
	}
}

func TestCanvasToPPM(t *testing.T) {
	c := *NewCanvas(5, 3)
	c1 := NewColor(1.5, 0, 0)
	c2 := NewColor(0, 0.5, 0)
	c3 := NewColor(-0.5, 0, 1)
	WritePixel(c, 0, 0, c1)
	WritePixel(c, 2, 1, c2)
	WritePixel(c, 4, 2, c3)

	filename := "test_output.ppm"

	err := CanvasToPPM(c, filename)
	if err != nil {
		t.Fatalf("CanvasToPPM returned an error: %v", err)
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	expectedContent := `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`
	if string(content) != expectedContent {
		t.Errorf("File content does not match expected. \nGot:\n%s\nWant:\n%s", string(content), expectedContent)
	}
	err = os.Remove(filename)
	if err != nil {
		t.Fatalf("Failed to remove test output file: %v", err)
	}
	
}

func TestSplittingLongLines(t *testing.T) {
	canvas := *NewCanvas(10, 2)
	color := NewColor(1, 0.8, 0.6)
	for y := 0; y < canvas.height; y++ {
		for x := 0; x < canvas.width; x++ {
			WritePixel(canvas, x, y, color)
		}
	}

	filename := "test_output.ppm"

	err := CanvasToPPM(canvas, filename)
	if err != nil {
		t.Fatalf("CanvasToPPM returned an error: %v", err)
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) < 8 {
		t.Fatalf("Output file does not contain expected number of lines.")
	}
	extractedLines := strings.Join(lines[3:7], "\n")
	expectedLinesFourToSeven := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153`
	if extractedLines != expectedLinesFourToSeven {
		t.Errorf("File content does not match expected. \nGot:\n%s\nWant:\n%s", extractedLines, expectedLinesFourToSeven)
	}
	err = os.Remove(filename)
	if err != nil {
		t.Fatalf("Failed to remove test output file: %v", err)
	}
}

func TestEndingWithNewline(t *testing.T) {
	canvas := *NewCanvas(5, 3)
	filename := "test_output.ppm"
	err := CanvasToPPM(canvas, filename)
	if err != nil {
		t.Fatalf("CanvasToPPM returned an error: %v", err)
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	lastChar := content[len(content) - 1]
	if lastChar != '\n' {
		t.Errorf("Expecting the last character to be %v, got %v", '\n', lastChar)
	}
	err = os.Remove(filename)
	if err != nil {
		t.Fatalf("Failed to remove test output file: %v", err)
	}
}
