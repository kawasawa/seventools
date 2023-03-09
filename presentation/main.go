//go:build js && wasm
// +build js,wasm

package main

import (
	"context"
	"fmt"
	"go-onion-sample/domain/entity"
	"go-onion-sample/usecase"
	"strconv"
	"syscall/js"
)

// func main() {
// 	err := execute()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	c := make(chan struct{})

	fmt.Println("Hello, WebAssembly!")
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func add(this js.Value, args []js.Value) interface{} {
	println(args[0].Int() + args[1].Int())
	return nil
}

func subtract(this js.Value, args []js.Value) interface{} {
	println(args[0].Int() - args[1].Int())
	return nil
}

func execute() error {
	ctx := context.Background()
	//params := entity.ParseOSArgs(os.Args)
	params := &entity.Params{}
	params.A = 10
	params.B = 20

	service, err := di(ctx)
	if err != nil {
		return err
	}

	addResult, err := service.Add(ctx, params)
	if err != nil {
		return err
	}

	subResult, err := service.Subtract(ctx, params)
	if err != nil {
		return err
	}

	fmt.Println("add: " + strconv.Itoa(addResult))
	fmt.Println("sub: " + strconv.Itoa(subResult))
	return nil
}

func di(ctx context.Context) (calcService usecase.ICalcService, err error) {
	return usecase.NewCalcService(), nil
}
