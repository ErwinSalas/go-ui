//go:build js && wasm
// +build js,wasm

package types

// Widget base para todos los componentes
type Widget interface {
	Render() string
}
