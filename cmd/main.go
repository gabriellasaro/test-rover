package main

import (
	"bufio"
	"fmt"
	"os"

	testrover "github.com/gabriellasaro/test-rover"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("file path not found in arguments")
		os.Exit(1)
	}

	read, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer read.Close()

	fileScan := bufio.NewScanner(read)
	fileScan.Split(bufio.ScanLines)

	fileScan.Scan()
	plateau, err := testrover.NewPlateauByParse(fileScan.Text())
	if err != nil {
		panic(err)
	}

	for fileScan.Scan() {
		rover, err := plateau.AddRoverByParse(fileScan.Text())
		if err != nil {
			panic(err)
		}

		fileScan.Scan()
		if err := rover.Commands(fileScan.Text()); err != nil {
			panic(err)
		}

		fmt.Println(rover.CurrentPosition())
	}
}
