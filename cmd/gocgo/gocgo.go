package main

import (
	"fmt"
	"github.com/EngineeringSoftware/gocgo/pkg/visitors"
	"os"
)

func main() {
	v := visitors.NewFuncDeleteVisitor()
	err := visitors.VisitFile(v, os.Args[1])

	fmt.Printf("%v\n", v.Buff.String())
	fmt.Printf("Done with error: %v\n", err)
}
