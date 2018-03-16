package goio

import (
	//"bufio"
	"fmt"
	//"io"
	"os"
	"strconv"
)

// WriteValues xxx
func WriteValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "sdfsdfsxxxx\n")
	}
	return nil
}
