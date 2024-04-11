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

	wg := &sync.WaitGroup{}
	c := canvas.NewCanvas()
	go c.RenderLoop()

	for i, l := range text {
		go func() {
			wg.Add(1)
			chars := strings.Split(l, "")
			for j := 0; j < len(chars); j++ {
				time.Sleep(time.Millisecond * 100)
				j2 := len(chars) - 1 - j
				s1, s2 := chars[j], chars[j2]
				c.ApplyPaint([2]int{i, j}, s1)
				c.ApplyPaint([2]int{i, j2}, s2)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	c.Stop = true
	<-c.Done
}
