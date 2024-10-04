//go:build js && wasm
// +build js,wasm

package widgets

import (
	"fmt"

	"github.com/ErwinSalas/go-ui/internal/dom"
)

// ListTile es un ítem en el Drawer
type ListTile struct {
	Title          string
	Selected       bool
	OnTap          func() // Función de callback para manejar el clic
	ID             string
	MountedChannel chan bool
}

// Render genera el HTML para ListTile y asigna el evento en el DOM
func (l ListTile) Render() string {
	selectedClass := ""
	if l.Selected {
		selectedClass = "selected"
	}

	go l.listenForMount()

	return fmt.Sprintf(`
		<div id="%s" class="list-tile %s">
			%s
		</div>
	`, l.ID, selectedClass, l.Title)
}

// listenForMount escucha el canal y asigna el evento cuando se monta
func (l *ListTile) listenForMount() {
	fmt.Println("ListTile.listenForMount")
	// Esperar la señal de montaje
	for range l.MountedChannel {
		// Asignar evento de clic
		dom.AddEventListener(l.ID, "click", l.OnTap)
	}
}
