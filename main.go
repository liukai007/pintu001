package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.NewWithID("pintu")
	win := a.NewWindow("PingTu")

	cc := container.NewBorder(nil, nil, nil, nil, nil)
	win.SetContent(cc)
	win.SetMaster()
	win.Resize(fyne.NewSize(800, 700))
	win.ShowAndRun()
}
