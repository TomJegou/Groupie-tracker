package main

import (
	"absolut-music/src"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go src.StartServer(&wg, 3)
	wg.Wait()
}
