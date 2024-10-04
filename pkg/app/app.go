//go:build js && wasm
// +build js,wasm

package app

import (
	"github.com/ErwinSalas/go-ui/internal/dom"
	"github.com/ErwinSalas/go-ui/internal/types"
	"github.com/ErwinSalas/go-ui/internal/widgets"
)

// App es un StatelessWidget
type App struct {
	mountedChannel chan bool
}

// Build crea la interfaz de App
func (m App) Build() types.Widget {
	return widgets.MaterialApp{
		Title: "Drawer Demo",
		ID:    "materialApp",
		Home: &MyHomePage{
			Title:    "Drawer Demo",
			parentID: "materialApp",
		},
	}
}

func NewApp() {
	// Crear la instancia de la aplicación
	app := App{
		mountedChannel: make(chan bool),
	}

	// Renderizar la aplicación inicial

	materialApp := app.Build().(widgets.MaterialApp)
	dom.Mount("app", materialApp)
	// Mantener el programa en ejecución
	select {}
}
