//go:build js && wasm
// +build js,wasm

package widgets

import (
	"fmt"

	"github.com/ErwinSalas/go-ui/pkg/ui/dom"
)

// listenForMount escucha el canal y asigna el evento cuando se monta
func (ib *IconButton) listenForMount() {
	fmt.Println("IconButton.listenForMount")
	// Esperar la señal de montaje
	for range ib.MountedChannel {
		fmt.Println("IconButton.listenForMount mounted")
		dom.AddEventListener(ib.ID, "click", ib.OnClick)
	}
}

// IconButton es un botón con ícono
type IconButton struct {
	Icon           string
	Class          string
	OnClick        func() // Función de callback para manejar el clic
	ID             string
	MountedChannel chan bool
}

// Render genera el HTML para IconButton y asigna el evento en el DOM
func (i IconButton) Render() string {

	go i.listenForMount()

	return fmt.Sprintf(`<button id="%s" class="%s">%s</button>`, i.ID, i.Class, i.Icon)
}
