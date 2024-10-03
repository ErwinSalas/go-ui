//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/ErwinSalas/go-ui/pkg/app"
)

func main() {
	app.NewApp()
}
