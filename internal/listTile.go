//go:build js && wasm
// +build js,wasm

package internal

import "fmt"

// ListTile es un ítem en el Drawer
type ListTile struct {
	Title    string
	Selected bool
	OnTap    func()
}

// Render genera el HTML para ListTile
func (l ListTile) Render() string {
	selectedClass := ""
	if l.Selected {
		selectedClass = "selected"
	}
	return fmt.Sprintf(`
		<div class="list-tile %s" onclick="goFunction()">
			%s
		</div>
	`, selectedClass, l.Title)
}
