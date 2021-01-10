/* Based on https://github.com/pci/talks/blob/master/2020/01-Go-Brum-Go-WASM */

package main

import (
	"fmt"
	"syscall/js"
)

// testWASM is a simple function to test that the WASM is working
func testWASM(this js.Value, i []js.Value) interface{} {
	fmt.Println("WASM is working!")
	fmt.Printf("Received %d args\n", len(i))
	return true
}

// registerFunctions adds any required functions to the js
func registerFunctions() {
	js.Global().Set("testWASM", js.FuncOf(testWASM))
}

func main() {
	c := make(chan struct{})
	fmt.Println("WASM Go Initialised")

	registerFunctions()

	// channel is unused, so this prevents main from terminating and killing the WASM
	<-c
}
