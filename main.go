package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("pintu")
	win := a.NewWindow("PingTu")

	//list2(win)
	//menu(win)
	//popUpMenu(win)
	//popUp(win)
	//icon(win)
	//set11()
	win.Resize(fyne.NewSize(800, 700))
	win.ShowAndRun()
}
