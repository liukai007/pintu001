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

/*
教程来源
https://www.bilibili.com/video/BV1zD4y167UL
*/
func main() {
	a := app.NewWithID("pintu")
	win := a.NewWindow("embed image browser")
	img := canvas.NewImageFromImage(nil)
	n := 3
	right := container.NewGridWithColumns(n)
	pieces := make([]fyne.CanvasObject, 9)
	btn1 := widget.NewButton("Split", func() {
		if img.Image == nil {
			return
		}
		rgb1 := ImageToRGBA(img.Image)
		w := rgb1.Rect.Dx() / n
		h := rgb1.Rect.Dy() / n
		n0 := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				v := canvas.NewImageFromImage(rgb1.SubImage(
					image.Rect(w*j, h*i, w*(j+1), h*(i+1))))
				pieces[n0] = v
				n0++
			}
		}
		right.Objects = pieces
		right.Refresh()
	})
	btnSwk := widget.NewButton("swk", func() {
		SetPic(img, swk)
		btn1.OnTapped()
	})
	btnZbj := widget.NewButton("zbj", func() {
		SetPic(img, zbj)
		btn1.OnTapped()
	})

	top := container.NewHBox(btnSwk, btnZbj, btn1)
	cc := container.NewHSplit(img, right)
	win.SetContent(container.NewBorder(top, nil, nil, nil, cc))
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
