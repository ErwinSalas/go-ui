//go:build js && wasm
// +build js,wasm

package internal

import (
	"fmt"
	"syscall/js"
)

// Drawer es ahora un widget con estado que puede abrirse y cerrarse
type Drawer struct {
	Items   []Widget
	Class   string
	Visible bool // Si el drawer está visible o no
	ID      string
}

// NewDrawer crea un Drawer con estado inicial oculto
func NewDrawer(items []Widget, class string, id string) *Drawer {
	return &Drawer{
		Items:   items,
		Class:   class,
		Visible: false,
		ID:      id,
	}
}

// Render genera el HTML para Drawer, controlando su visibilidad
func (d Drawer) Render() string {
	fmt.Printf("Rendering Drawer %t", d.Visible)
	visibility := "hidden"
	if d.Visible {
		visibility = "visible"
	}

	content := ""
	for _, item := range d.Items {
		content += item.Render()
	}
	return fmt.Sprintf(`
		<div class="drawer" style="visibility: %s;">
			%s
		</div>
	`, visibility, content)
}

// Método para abrir el Drawer
func (d *Drawer) OpenDrawer() {
	d.Visible = true
	d.Rerender()
}

// Método para cerrar el Drawer
func (d *Drawer) CloseDrawer() {
	d.Visible = false
	d.Rerender()
}

// Rerender fuerza el re-renderizado del Drawer
func (d *Drawer) Rerender() {
	document := js.Global().Get("document")
	container := document.Call("getElementById", d.ID)
	container.Set("innerHTML", d.Render())
}
