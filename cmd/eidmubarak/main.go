package main

import (
	"strings"
	"sync"
	"time"

	"github.com/sifatulrabbi/eidmubarak/internals/canvas"
)

func main() {
	text := []string{
		"##############   #########   ##########",
		"##############   #########   ###########",
		"###                 ###      ###      ###",
		"###                 ###      ###       ###",
		"###                 ###      ###       ###",
		"##########          ###      ###        ###",
		"##########          ###      ###        ###",
		"###                 ###      ###       ###",
		"###                 ###      ###       ###",
		"###                 ###      ###     ###",
		"##############   #########   ##########",
		"##############   #########   #########",
		"",
		"        M B A R A K    2 0 2 4",
	}
	wg := sync.WaitGroup{}

	c := canvas.NewCanvas()
	go c.RenderLoop()

	for i, l := range text {
		go func() {
			wg.Add(1)
			for j, s := range strings.Split(l, "") {
				time.Sleep(time.Millisecond * 100)
				c.ApplyPaint([2]int{i, j}, s)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	c.Stop = true
	<-c.Done
}
