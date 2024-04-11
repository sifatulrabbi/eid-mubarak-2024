package main

import (
	"strings"
	"sync"
	"time"

	"github.com/sifatulrabbi/eidmubarak/internals/canvas"
)

var text = []string{
	"##############   #########   #########",
	"##############   #########   ##########",
	"###                 ###      ###     ###",
	"###                 ###      ###      ###",
	"###                 ###      ###       ###",
	"##########          ###      ###        ###",
	"##########          ###      ###        ###",
	"###                 ###      ###       ###",
	"###                 ###      ###      ###",
	"###                 ###      ###     ###",
	"##############   #########   ##########",
	"##############   #########   #########",
	"",
	"                M B A R A K",
}

func main() {
	wg := sync.WaitGroup{}

	c := canvas.NewCanvas()
	go c.RenderLoop()

	for i, l := range text {
		go func() {
			wg.Add(1)
			for j, s := range strings.Split(l, "") {
				c.ApplyPaint([2]int{i, j}, s)
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	c.Stop = true
	<-c.Done
}
