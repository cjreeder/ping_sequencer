package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Reading File and Getting ping sequence")
	var filepath string
	flag.StringVar(&filepath, "opt", "", "Usage")
	flag.Parse()
	fmt.Println(filepath)
}
