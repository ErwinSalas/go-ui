//go:build js && wasm
// +build js,wasm

package widgets

import (
	"fmt"

	"github.com/ErwinSalas/go-ui/internal/types"
)

// Text es un widget de texto
type Text struct {
	Content string
}

// Render genera el HTML para el texto
func (t Text) Render() string {
	return fmt.Sprintf(`<p>%s</p>`, t.Content)
}

// Center alinea su hijo en el centro
type Center struct {
	Child types.Widget
}

// Render genera el HTML para Center
func (c Center) Render() string {
	return fmt.Sprintf(`<div class="center">%s</div>`, c.Child.Render())
}
