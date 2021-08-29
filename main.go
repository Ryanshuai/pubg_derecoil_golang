package main

import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"golang_derecoil/src"
)


func main() {
	press := src.Press{}
	robotgo.EventHook(hook.MouseHold, []string{}, func(e hook.Event) {
		press.Stop = false
		distances := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		for i := 0; i < len(distances); i++ {
			distances[i] *= 1
		}
		press.Distances = distances
		go press.Act()
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)

}
