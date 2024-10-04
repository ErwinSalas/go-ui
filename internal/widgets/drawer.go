//go:build js && wasm
// +build js,wasm

package widgets

import (
	"fmt"

	"github.com/ErwinSalas/go-ui/internal/types"
)

// Drawer es ahora un widget con estado que puede abrirse y cerrarse
type Drawer struct {
	Items   []types.Widget
	Class   string
	Visible bool // Si el drawer está visible o no
	ID      string
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
