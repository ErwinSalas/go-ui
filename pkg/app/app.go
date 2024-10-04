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
	drawerOpen    bool
}

func (h *MyHomePage) Render() string {
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
				OnClick: h.toggleDrawer,
				Class:   "menu-icon",
				ID:      "menuButton", // Asignar un ID al botón
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
			Visible: h.drawerOpen, // Controlar la visibilidad del Drawer según el estado drawerOpen
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
		ID: fmt.Sprintf("drawerItem%d", index), // Asignar un ID a cada ítem del Drawer
	}
}

// Método para manejar la selección de ítems del Drawer
func (h *MyHomePage) onItemTapped(index int) {
	h.selectedIndex = index
	h.drawerOpen = false // Cerrar el Drawer al seleccionar un ítem
	h.Rerender()
}

// Método para abrir y cerrar el Drawer
func (h *MyHomePage) toggleDrawer() {
	h.drawerOpen = !h.drawerOpen
	h.Rerender()
}

// Rerender fuerza el re-renderizado del componente
func (h *MyHomePage) Rerender() {
	document := js.Global().Get("document")
	container := document.Call("getElementById", "app")
	container.Set("innerHTML", h.Render())

	// Asignar los manejadores de eventos después de renderizar
	h.assignEventHandlers()
}

// Asignar los eventos a los elementos interactivos después de renderizar
func (h *MyHomePage) assignEventHandlers() {
	// Asignar el evento al botón de menú
	menuButton := js.Global().Get("document").Call("getElementById", "menuButton")
	menuButton.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("Menu button clicked")
		h.toggleDrawer()
		return nil
	}))

	// Asignar los eventos a los ítems del Drawer
	for i := 0; i < 3; i++ { // Sabemos que hay 3 ítems
		index := i
		item := js.Global().Get("document").Call("getElementById", fmt.Sprintf("drawerItem%d", index))
		item.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			h.onItemTapped(index)
			return nil
		}))
	}
}

func NewApp() {
	// Crear la instancia de la aplicación
	app := App{}

	// Renderizar la aplicación inicial
	document := js.Global().Get("document")
	container := document.Call("getElementById", "app")
	materialApp := app.Build().(internal.MaterialApp) // Convertir a MaterialApp

	// Asegurarnos de que el Home es un MyHomePage
	if homePage, ok := materialApp.Home.(*MyHomePage); ok {
		container.Set("innerHTML", materialApp.Render())
		// Asignar manejadores de eventos iniciales
		homePage.assignEventHandlers()
	} else {
		panic("Home is not *MyHomePage")
	}

	// Mantener el programa en ejecución
	select {}
}
