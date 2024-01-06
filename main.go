package main

import (
	"bytes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image"
)

var swk []byte
var zbj []byte

func main() {
	a := app.NewWithID("pintu")
	win := a.NewWindow("PingTu")
	msg := widget.NewLabel("")
	entry := widget.NewEntry()
	btn1 := widget.NewButton("swk", func() {

	})
	btn2 := widget.NewButton("zbj", func() {

	})
	btnRefresh := widget.NewButton("Refresh", func() {

	})
	left := canvas.NewImageFromFile("swk.jpg")
	right := container.NewGridWithColumns(3)
	for i := 0; i < 9; i++ {
		right.Add(canvas.NewImageFromFile("swk.jpg"))
	}
	split := container.NewHSplit(left, right)
	top := container.NewHBox(entry, btn1, btn2, btnRefresh)
	cc := container.NewBorder(top, msg, nil, nil, split)
	win.SetContent(cc)
	win.SetMaster()
	win.Resize(fyne.NewSize(800, 700))
	win.ShowAndRun()
}

func SetPic(img *canvas.Image, picData []byte) (err error) {
	buf := bytes.NewReader(picData)
	img.Image, _, err = image.Decode(buf)
	if err != nil {
		return err
	}
	img.Refresh()
	return nil
}
