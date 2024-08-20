package main

import (
	"flag"
	"fmt"
)

func main() {
	// * FLTER FLAG
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "all files including hide files")
	flagNumberRecords := flag.Int("n", 0, "number of record")

	// * ORDER FLAGS
	hasOrderByTime := flag.Bool("t", false, "sort by time, oldest first")
	hasOrderBySize := flag.Bool("s", false, "sort by size, smallest first")
	hasorderReverse := flag.Bool("r", false, "sort reverse with sorting")

	flag.Parse()

	fmt.Println(*flagPattern)
	fmt.Println(*flagAll)
	fmt.Println(*flagNumberRecords)
	fmt.Println(*hasOrderByTime)
	fmt.Println(*hasOrderBySize)
	fmt.Println(*hasorderReverse)
}
