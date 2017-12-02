package main

import (
	//"bufio"
	"flag"
	"fmt"
	//"io"
	//"os"
	//"strconv"
	"goSrc/test01/sorter/algorithm"
	"goSrc/test01/sorter/goio"
)

var infile = flag.String("i", "infile", "File constains values for sorting")
var outfile = flag.String("o", "outfile", "File to receive sorted values")
var algorith = flag.String("a", "qsort", "Sort alforithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorith =", *algorith)
	}

	values, err := goio.ReadValues(*infile)
	if err == nil {
		if *algorith == "bsort" {
			algorithm.Bsort(values)
		}
		fmt.Println("Read values:", values)
		goio.WriteValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
