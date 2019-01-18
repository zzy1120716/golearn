// 生成随机利萨如图形的GIF动画
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.White, color.Black, color.RGBA{0xff, 0xff, 0, 255}}

const (
	whiteIndex  = 0 // first color in palette
	blackIndex  = 1 // next color in palette
	yellowIndex = 2
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//cycles := 5
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			//params := r.URL.Query()
			//cycles, err := strconv.Atoi(params["cycles"][0])
			//if err != nil {
			//	fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			//}
			//lissajous(w, float64(cycles))
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	//lissajous(os.Stdout, float64(cycles))
	lissajous(os.Stdout)
}

//func lissajous(out io.Writer, cycles float64) {
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
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
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5), // 最后插入的逗号不会导致编译错误，这是Go编译器的一个特性
				yellowIndex,
			) // 小括弧另起一行缩进，和大括弧的风格保持一致
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
