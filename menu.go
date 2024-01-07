package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func menu(w fyne.Window) {
	lblMsg := widget.NewLabel("")
	menu := widget.NewMenu(fyne.NewMenu("File",
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

	c := container.NewVBox(lblMsg, menu)
	w.SetContent(c)
}
