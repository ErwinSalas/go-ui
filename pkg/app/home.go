//go:build js && wasm
// +build js,wasm

package app

import (
	"fmt"

	"github.com/ErwinSalas/go-ui/pkg/ui/dom"
	"github.com/ErwinSalas/go-ui/pkg/ui/types"
	"github.com/ErwinSalas/go-ui/pkg/ui/widgets"
)

// MyHomePage es un StatefulWidget
type MyHomePage struct {
	parentID        string
	Title           string
	selectedIndex   int
	drawerOpen      bool
	mountedChannels []chan bool // Canal para notificar montaje

}

func (h *MyHomePage) Render() string {
	defer func() {
		go func() {
			for _, ch := range h.mountedChannels {
				ch <- true
			}
			h.mountedChannels = []chan bool{}
		}()
	}()
	return h.Build().Render()
}

// drawerItem crea un elemento para el Drawer
func (h *MyHomePage) drawerItem(label string, index int) types.Widget {
	return widgets.ListTile{
		Title:    label,
		Selected: h.selectedIndex == index,
		OnTap: func() {
			h.onItemTapped(index)
		},
		ID:             fmt.Sprintf("drawerItem%d", index), // Asignar un ID a cada ítem del Drawer
		MountedChannel: h.addChannel(),
	}
}

// Método para manejar la selección de ítems del Drawer
func (h *MyHomePage) onItemTapped(index int) {
	fmt.Printf("Item %d seleccionado\n", index)
	h.selectedIndex = index
	h.drawerOpen = false // Cerrar el Drawer al seleccionar un ítem
	h.Rerender()
}

// Método para abrir y cerrar el Drawer
func (h *MyHomePage) toggleDrawer() {
	h.drawerOpen = !h.drawerOpen
	h.Rerender()
}

// Añade un canal al slice y lo retorna
func (h *MyHomePage) addChannel() chan bool {
	channel := make(chan bool)
	h.mountedChannels = append(h.mountedChannels, channel)
	return channel
}

// Rerender fuerza el re-renderizado del componente
func (h *MyHomePage) Rerender() {
	dom.Mount(h.parentID, h)
}

func (h *MyHomePage) Build() types.Widget {
	// Opciones de widgets
	widgetOptions := []types.Widget{
		widgets.Text{Content: "Index 0: Home"},
		widgets.Text{Content: "Index 1: Business"},
		widgets.Text{Content: "Index 2: School"},
	}

	// Crear el Scaffold
	return widgets.Scaffold{
		AppBar: widgets.AppBar{
			Title: h.Title,
			Leading: widgets.IconButton{
				Icon:           "☰",
				OnClick:        h.toggleDrawer,
				Class:          "menu-icon",
				ID:             "menuButton", // Asignar un ID al botón
				MountedChannel: h.addChannel(),
			},
		},
		Body: widgets.Center{
			Child: widgetOptions[h.selectedIndex],
		},
		Drawer: widgets.Drawer{
			Items: []types.Widget{
				h.drawerItem("Home", 0),
				h.drawerItem("Business", 1),
				h.drawerItem("School", 2),
			},
			Visible: h.drawerOpen, // Controlar la visibilidad del Drawer según el estado drawerOpen
		},
	}
}
