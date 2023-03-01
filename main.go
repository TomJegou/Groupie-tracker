package main

import (
	"absolut-music/src"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	src.StartServer(&wg)
	wg.Wait()
}
