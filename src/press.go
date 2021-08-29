package src

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

type Press struct {
	Distances []float32
	Stop      bool
}

func (press Press) Act() {
	// 15ms / move
	tic := time.Now()
	for idx := 0; idx < len(press.Distances) && !press.Stop; idx++ {
		robotgo.MoveRelative(0, int(press.Distances[idx]))
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println(time.Since(tic))
}
