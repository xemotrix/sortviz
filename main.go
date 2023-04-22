package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/xemotrix/gocgl"

	"github.com/veandco/go-sdl2/sdl"
)

var engine = gocgl.NewEngine(WIDTH, HEIGHT)

const (
	FACTOR = 50
	WIDTH  = 16 * FACTOR
	HEIGHT = 9 * FACTOR

	COLOR_BLK uint32 = 0xff000000
	COLOR_WHT uint32 = 0xffffffff
	COLOR_RED uint32 = 0xffff3030

	TSTEP time.Duration = 0 //1 * time.Microsecond

	SIZE = 16 * FACTOR
)

func handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			os.Exit(0)
		}
	}
}

var ARR = randomArray(SIZE)

func main() {
	// BubbleSort(ARR)
	RenderArray(ARR)

	ARR = randomArray(SIZE)
	QuickSort(ARR)

	RenderArray(ARR)
	for {
		handleEvents()
	}
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot := arr[len(arr)-1]
	i, j := 0, len(arr)-2

	for i <= j {
		handleEvents()
		RenderArray(ARR, arr[j], arr[i])
		time.Sleep(TSTEP)
		if arr[i] > pivot {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		} else {
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]

	QuickSort(arr[:i])
	QuickSort(arr[i:])

}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			handleEvents()
			RenderArray(arr, arr[j+1], arr[j])
			time.Sleep(TSTEP)
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func randomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	rand.Shuffle(size, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func RenderArray(
	arr []int,
	hl ...int,
) {
	engine.Image.FillWithColor(COLOR_BLK)
	pad := float64(2)
	step := WIDTH / float64(len(arr))

	for i, v := range arr {
		rect := gocgl.Rectangle{
			Points: [2]gocgl.Point{
				{X: float64(i) * step, Y: HEIGHT * (1 - float64(v)/SIZE)},
				{X: float64(i)*step + step - pad, Y: HEIGHT},
			},
		}
		color := COLOR_WHT
		for _, h := range hl {
			if v == h {
				color = COLOR_RED
			}
		}
		rect.Render(engine.Image, color)
	}

	engine.Render()
}
