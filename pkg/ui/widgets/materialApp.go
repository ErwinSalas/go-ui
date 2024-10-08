//go:build js && wasm
// +build js,wasm

package widgets

import (
	"fmt"

	"github.com/ErwinSalas/go-ui/pkg/ui/types"
)

// MaterialApp simula el widget principal en Flutter
type MaterialApp struct {
	Title string
	Home  types.Widget
	ID    string
}

// Render genera el HTML para MaterialApp
func (m MaterialApp) Render() string {
	return fmt.Sprintf(`
		<div id="%s" class="app-container">
			%s
		</div>
	`, m.ID, m.Home.Render())
}
