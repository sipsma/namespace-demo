package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	host, _ := os.Hostname()
	fmt.Printf("hello from namespace demo (%s/%s, built with %s) on %s\n", runtime.GOOS, runtime.GOARCH, runtime.Version(), host)
}
