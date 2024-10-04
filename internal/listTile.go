//go:build js && wasm
// +build js,wasm

package internal

import (
	"fmt"
	"syscall/js"
)

// ListTile es un ítem en el Drawer
type ListTile struct {
	Title    string
	Selected bool
	OnTap    func() // Función de callback para manejar el clic
	ID       string
}

// Render genera el HTML para ListTile y asigna el evento en el DOM
func (l ListTile) Render() string {
	selectedClass := ""
	if l.Selected {
		selectedClass = "selected"
	}

	// Asignar el evento al ítem
	go assignListTileEvent(l.ID, l.OnTap)

	return fmt.Sprintf(`
		<div id="%s" class="list-tile %s">
			%s
		</div>
	`, l.ID, selectedClass, l.Title)
}

// Asignar el evento de clic a un ListTile
func assignListTileEvent(id string, onTap func()) {
	// Obtener el elemento y asignar el evento de clic
	item := js.Global().Get("document").Call("getElementById", id)
	if !item.IsNull() {
		item.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			onTap() // Llamar al manejador de clic definido
			return nil
		}))
	}
}
