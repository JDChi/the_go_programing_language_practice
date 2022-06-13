package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette1 = []color.Color{color.White, color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}, color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}}

const (
	blueIndex1 = 2
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// 需要调用 ParseForm() 方法，才能解析 Form
		r.ParseForm()
		cycles := 5
		if r.Form.Has("cycles") {
			cyclesParam, err := strconv.Atoi(r.Form.Get("cycles"))
			if err == nil {
				cycles = cyclesParam
			}
		}
		lissajous2(w, cycles)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	// 需要调用 ParseForm() 方法，才能解析 Form
	r.ParseForm()
	cycles := 5
	if r.Form.Has("cycles") {
		cyclesParam, err := strconv.Atoi(r.Form.Get("cycles"))
		if err == nil {
			cycles = cyclesParam
		}
	}
	lissajous2(w, cycles)

}

func lissajous2(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette1)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blueIndex1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
