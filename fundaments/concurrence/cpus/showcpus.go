package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Total CPUS in your computer: ", runtime.NumCPU())
}
