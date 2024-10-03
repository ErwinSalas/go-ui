//go:build js && wasm
// +build js,wasm

package app

import (
	"fmt"
	"syscall/js"

	"github.com/ErwinSalas/go-ui/internal"
)

// App es un StatelessWidget
type App struct {
}

// Build crea la interfaz de App
func (m App) Build() internal.Widget {
	return internal.MaterialApp{
		Title: "Drawer Demo",
		Home:  &MyHomePage{Title: "Drawer Demo"},
	}
}

// MyHomePage es un StatefulWidget
type MyHomePage struct {
	Title         string
	selectedIndex int
}

func (h *MyHomePage) Render() string {
	// Llamar a Build y renderizar el resultado
	return h.Build().Render()
}

// Build define la interfaz de MyHomePage
func (h *MyHomePage) Build() internal.Widget {
	// Opciones de widgets
	widgetOptions := []internal.Widget{
		internal.Text{Content: "Index 0: Home"},
		internal.Text{Content: "Index 1: Business"},
		internal.Text{Content: "Index 2: School"},
	}

	// Crear el Scaffold
	return internal.Scaffold{
		AppBar: internal.AppBar{
			Title: h.Title,
			Leading: internal.IconButton{
				Icon:    "☰",
				OnClick: h.openDrawer,
				Class:   "menu-icon",
			},
		},
		Body: internal.Center{
			Child: widgetOptions[h.selectedIndex],
		},
		Drawer: internal.Drawer{
			Items: []internal.Widget{
				h.drawerItem("Home", 0),
				h.drawerItem("Business", 1),
				h.drawerItem("School", 2),
			},
		},
	}
}

// drawerItem crea un elemento para el Drawer
func (h *MyHomePage) drawerItem(label string, index int) internal.Widget {
	return internal.ListTile{
		Title:    label,
		Selected: h.selectedIndex == index,
		OnTap: func() {
			h.onItemTapped(index)
		},
	}
}

// Método para manejar la selección de ítems del Drawer
func (h *MyHomePage) onItemTapped(index int) {
	h.selectedIndex = index
	h.Rerender()
}

// Método para abrir el Drawer (simulado en este caso)
func (h *MyHomePage) openDrawer() {
	fmt.Println("Opening Drawer")
}

// Rerender fuerza el re-renderizado del componente
func (h *MyHomePage) Rerender() {
	document := js.Global().Get("document")
	container := document.Call("getElementById", "app")
	container.Set("innerHTML", h.Build().Render())
}

// main inicia la aplicación
func NewApp() {
	// Crear la instancia de la aplicación
	app := App{}

	// Renderizar la aplicación inicial
	document := js.Global().Get("document")
	container := document.Call("getElementById", "app")
	container.Set("innerHTML", app.Build().Render())

	// Mantener el programa en ejecución
	select {}
}
