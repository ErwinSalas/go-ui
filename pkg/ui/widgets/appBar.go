//go:build js && wasm
// +build js,wasm

package widgets

import (
	"fmt"

	"github.com/ErwinSalas/go-ui/pkg/ui/types"
)

// AppBar es la barra superior con título e íconos
type AppBar struct {
	Title   string
	Leading types.Widget
}

// Render genera el HTML para AppBar
func (a AppBar) Render() string {
	return fmt.Sprintf(`
		<div class="app-bar">
			%s
			<h1>%s</h1>
		</div>
	`, a.Leading.Render(), a.Title)
}
