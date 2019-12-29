package main

import (
	"fmt"
	"information_theory_assignment4/cyclic_codes"
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
	var n, p int
	_, _ = fmt.Fscanf(inpf, "%d\n%d", &p, &n)
	pol := cyclic_codes.Polynomial{
		Base:         p,
		Degree:       n - 1,
		Coefficients: nil,
	}
	pol.Coefficients = make([]int, n)
	for i := n - 1; i >= 0; i-- {
		_, _ = fmt.Fscanf(inpf, "%d", &pol.Coefficients[i])
	}

	res, err := cyclic_codes.ParityCheck(pol)
	if err != nil {
		_, _ = fmt.Fprintln(outpf, "NO")
		return
	}
	_, _ = fmt.Fprintln(outpf, "YES")
	left := n
	for i := len(res.Coefficients) - 1; i >= 0; i-- {
		_, _ = fmt.Fprintf(outpf, "%d ", res.Coefficients[i])
		left--
	}
	for i := 0; i < left; i++ {
		_, _ = fmt.Fprintf(outpf, "0 ")
	}

	//matrix, perm :=
	//
	//parityMatrix := linearcodes.GetParityCheckMatrix(matrix, perm)
	//linearcodes.PrintBooleanMatrix(outpf, parityMatrix)

}
