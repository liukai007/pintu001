package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//go:embed swk.jpg
var icon11 []byte

func icon(w fyne.Window) {
	resource := fyne.NewStaticResource("icon", icon11)
	icon1 := widget.NewIcon(resource)

	fi := widget.NewFileIcon(nil)
	fi.SetSelected(true)
	fi.SetURI(fyne.CurrentApp().Storage().RootURI())
	w.SetIcon(resource)
	c := container.NewVBox(icon1, fi)
	w.SetContent(c)
}
