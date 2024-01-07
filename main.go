package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("pintu")
	win := a.NewWindow("PingTu")

	List1(win)
	win.Resize(fyne.NewSize(800, 700))
	win.ShowAndRun()
}

type person struct {
	name, age string
}

func List1(w fyne.Window) {
	lblMsg := widget.NewLabel("")
	//data := []string{"red", "blue", "yellow"}
	data := []person{
		{"name1", "age1"},
		{"name2", "age2"},
		{"name3", "age3"},
	}
	lst := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			//return widget.NewLabel("")
			return container.NewHBox(widget.NewLabel(""), widget.NewLabel(""))
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			//lbl := co.(*widget.Label)
			hbox := co.(*fyne.Container)
			lblLeft := hbox.Objects[0].(*widget.Label)
			lblRight := hbox.Objects[1].(*widget.Label)
			lblLeft.SetText(data[lii].name)
			lblRight.SetText(data[lii].age)
		},
	)
	lst.OnSelected = func(id widget.ListItemID) {
		lblMsg.SetText(data[id].name + ":" + data[id].age)
	}
	//c := container.NewHBox(lblMsg, lst)
	c := container.NewBorder(lblMsg, nil, nil, nil, lst)
	w.SetContent(c)
}
