package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func menu(w fyne.Window) {
	lblMsg := widget.NewLabel("")
	menu1 := widget.NewMenu(fyne.NewMenu("File",
		fyne.NewMenuItem("Open", func() {
			lblMsg.SetText("Open")
		}),
		fyne.NewMenuItem("Save", func() {
			lblMsg.SetText("Save")
		}),
		fyne.NewMenuItem("Close", func() {
			lblMsg.SetText("Close")
		}),
	))
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Open", func() {
			lblMsg.SetText("Open")
		}),
		fyne.NewMenuItem("Save", func() {
			lblMsg.SetText("Save")
		}),
		fyne.NewMenuItem("Close", func() {
			lblMsg.SetText("Close")
		}),
	)
	editMenu := fyne.NewMenu("Edit",
		fyne.NewMenuItem("Open", func() {
			lblMsg.SetText("111")
		}),
		fyne.NewMenuItem("Save", func() {
			lblMsg.SetText("Save")
		}),
		fyne.NewMenuItem("Close", func() {
			lblMsg.SetText("Close")
		}),
	)
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		editMenu)
	w.SetMainMenu(mainMenu)
	c := container.NewVBox(lblMsg, menu1)
	w.SetContent(c)
}
