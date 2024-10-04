//go:build js && wasm
// +build js,wasm

package internal

import (
	"fmt"
	"syscall/js"
)

// Widget base para todos los componentes
type Widget interface {
	Render() string
}

// MaterialApp simula el widget principal en Flutter
type MaterialApp struct {
	Title string
	Home  Widget
}

// Render genera el HTML para MaterialApp
func (m MaterialApp) Render() string {
	return fmt.Sprintf(`
		<div class="app-container">
			%s
		</div>
	`, m.Home.Render())
}

// Scaffold es un contenedor para AppBar, Body y Drawer
type Scaffold struct {
	AppBar AppBar
	Body   Widget
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

// AppBar es la barra superior con título e íconos
type AppBar struct {
	Title   string
	Leading Widget
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

// Text es un widget de texto
type Text struct {
	Content string
}

// Render genera el HTML para el texto
func (t Text) Render() string {
	return fmt.Sprintf(`<p>%s</p>`, t.Content)
}

// IconButton es un botón con ícono
type IconButton struct {
	Icon    string
	Class   string
	OnClick func() // Función de callback para manejar el clic
	ID      string
}

// Render genera el HTML para IconButton y asigna el evento en el DOM
func (i IconButton) Render() string {
	// Asignar el evento al botón
	go assignIconButtonEvent(i.ID, i.OnClick)

	return fmt.Sprintf(`<button id="%s" class="%s">%s</button>`, i.ID, i.Class, i.Icon)
}

// Asignar el evento de clic a un IconButton
func assignIconButtonEvent(id string, onClick func()) {
	// Obtener el elemento y asignar el evento de clic
	button := js.Global().Get("document").Call("getElementById", id)
	if !button.IsNull() {
		button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			onClick() // Llamar al manejador de clic definido
			return nil
		}))
	}
}

// Center alinea su hijo en el centro
type Center struct {
	Child Widget
}

// Render genera el HTML para Center
func (c Center) Render() string {
	return fmt.Sprintf(`<div class="center">%s</div>`, c.Child.Render())
}
