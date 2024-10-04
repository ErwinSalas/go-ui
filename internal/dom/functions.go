//go:build js && wasm
// +build js,wasm

package dom

import (
	"fmt"
	"syscall/js"

	"github.com/ErwinSalas/go-ui/internal/types"
)

func Mount(rootID string, w types.Widget) {
	root := js.Global().Get("document")
	container := root.Call("getElementById", "app")
	container.Set("innerHTML", w.Render())
}

func AddEventListener(id string, event string, callback func()) {
	fmt.Printf("AddEventListener %s, %s", id, event)
	document := js.Global().Get("document")
	element := document.Call("getElementById", id)
	element.Call("addEventListener", event, js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		callback()
		return nil
	}))
}
