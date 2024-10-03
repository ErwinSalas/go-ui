//go:build js && wasm
// +build js,wasm

package internal

import "fmt"

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
	OnClick func()
}

// Render genera el HTML para IconButton
func (i IconButton) Render() string {
	return fmt.Sprintf(`<button class="%s" onclick="goFunction()">%s</button>`, i.Class, i.Icon)
}

// Center alinea su hijo en el centro
type Center struct {
	Child Widget
}

// Render genera el HTML para Center
func (c Center) Render() string {
	return fmt.Sprintf(`<div class="center">%s</div>`, c.Child.Render())
}
