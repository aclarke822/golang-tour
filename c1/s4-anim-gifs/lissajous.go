//Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5      // number of complete x oscillator revolutions
		res     = 0.0001 // angular resolution
		size    = 500    // image canvas covers [-size..+size]
		nframes = 64     // number of animation frames
		delay   = 8      // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 3*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//Notes
/*
After importing a package whose path has multiple components, we refer to the package with a name that comes from the last component.
Constants declared with const are fixed at compile time and may be declared at package or function level.
The value of a constant must be a number, string or boolean.
[]color.Color{...} and gif.GIF{...} are composite literals. The first one is a slice and the second one is a struct.
Composite literal expressions are a compact notation for instantiating any of Go's composite types from a sequence of element values.
A struct is a group of values called fields, often of different types, that are collected together in a single object that can be treated as a unit.
*/
