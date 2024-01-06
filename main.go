package main

import (
	"bytes"
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/draw"
)

//go:embed swk.jpg
var swk []byte

//go:embed zbj.png
var zbj []byte

func main() {
	a := app.NewWithID("pintu")
	win := a.NewWindow("embed image browser")
	img := canvas.NewImageFromImage(nil)
	btnSwk := widget.NewButton("swk", func() {
		SetPic(img, swk)
	})
	btnZbj := widget.NewButton("zbj", func() {
		SetPic(img, zbj)
	})
	right := container.NewGridWithColumns(3)
	for i := 0; i < 9; i++ {
		right.Add(canvas.NewImageFromFile("swk.jpg"))
	}
	top := container.NewHBox(btnSwk, btnZbj)
	cc := container.NewBorder(top, nil, nil, nil, img)
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

func ImageToRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	draw.Draw(dst, bounds, src, bounds.Min, draw.Src)
	return dst
}
