package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func popUpMenu(w fyne.Window) {
	lblMsg := widget.NewLabel("")
	fyneMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Open", func() {
			lblMsg.SetText("Open")
		}),
		fyne.NewMenuItem("Save", func() {
			lblMsg.SetText("Save")
		}),
	)
	pop := widget.NewPopUpMenu(fyneMenu, w.Canvas())
	btn := widget.NewButton("PopUp", func() {
		pop.Show()
	})
	var btn1 *widget.Button
	btn1 = widget.NewButton("PopUp", func() {
		pop.ShowAtPosition(btn1.Position())
	})
	c := container.NewVBox(lblMsg, btn, btn1)
	w.SetContent(c)
}
