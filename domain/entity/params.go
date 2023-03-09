package entity

import (
	"flag"
	"fmt"
	"strconv"
)

type Params struct {
	A int
	B int
}

func ParseOSArgs(args []string) *Params {
	params := Params{}
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.IntVar(&params.A, "a", 0, "number a")
	flagSet.IntVar(&params.B, "b", 0, "number b")
	flagSet.Parse(args[1:])
	fmt.Println(args)
	return &params
}

func ParseJSArgs(args []string) *Params {
	params := Params{}
	params.A, _ = strconv.Atoi(args[0])
	params.B, _ = strconv.Atoi(args[1])
	return &params
}
