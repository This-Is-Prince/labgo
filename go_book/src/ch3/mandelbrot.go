package ch3

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
)

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
func Mandelbrot() {
	createMandelbrotServer()
	// createMandelbrot(os.Stdout)
}

func createMandelbrotServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeMandelbrotToResponseWriter(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func writeMandelbrotToResponseWriter(w http.ResponseWriter) {
	createMandelbrot(w)
}

func createMandelbrot(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	err := png.Encode(out, img)
	if err != nil {
		log.Print(err)
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
