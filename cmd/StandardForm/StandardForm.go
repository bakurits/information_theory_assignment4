package main

import (
	"fmt"
	"information_theory_assignment3/linearcodes"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please enter correct parameters")
		return
	}
	argsWithoutProg := os.Args[1:]
	inp := argsWithoutProg[0]
	outp := argsWithoutProg[1]

	inpf, err := os.Open(inp)
	if err != nil {
		log.Fatal("error in opening file")
	}
	defer func() {
		err = inpf.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	outpf, err := os.OpenFile(outp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("error in opening file")
	}
	defer func() {
		err = outpf.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	matrix, perm := linearcodes.NormalizeMatrix(linearcodes.ReadMatrix(inpf))
	linearcodes.PrintBooleanMatrix(outpf, matrix)

	for _, val := range perm {
		_, _ = fmt.Fprintf(outpf, "%d ", val)
	}
}
