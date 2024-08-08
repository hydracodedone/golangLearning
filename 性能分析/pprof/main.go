package main

import (
	"net/http"
	_ "net/http/pprof"
	"pprof_demo/groutine"
	"runtime"
)

func main() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()
	groutine.GroutineDemo()

}
