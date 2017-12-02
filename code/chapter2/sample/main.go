package main

import (
	"log"
	"os"

	// go语言中，有些包没有用到，但是需要这些包中的init()函数，所以这样做
	_ "goSrc/code/chapter2/sample/matchers"
	//这个包在第22行被调用了
	"goSrc/code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
