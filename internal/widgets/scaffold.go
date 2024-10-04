//go:build js && wasm
// +build js,wasm

package widgets

import (
	"fmt"

	"github.com/ErwinSalas/go-ui/internal/types"
)

// Scaffold es un contenedor para AppBar, Body y Drawer
type Scaffold struct {
	AppBar AppBar
	Body   types.Widget
	Drawer Drawer
}

// Render genera el HTML para Scaffold
func (s Scaffold) Render() string {
	return fmt.Sprintf(`
		<div class="scaffold">
			%s
			<div class="body">%s</div>
			%s
		</div>
	`, s.AppBar.Render(), s.Body.Render(), s.Drawer.Render())
}
