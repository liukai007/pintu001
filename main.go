package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
	slider := widget.NewSlider(0, 1)
	slider.Resize(fyne.NewSize(111, 11))
	msgEntry := widget.NewMultiLineEntry()
	msgEntry1 := widget.NewMultiLineEntry()
	slider.Step = 0.01
	slider.OnChanged = func(f float64) {
		fmt.Println(f)
	}
	win.SetOnDropped(func(p fyne.Position, u []fyne.URI) {
		paths := ""
		for _, v := range u {
			paths = paths + v.Path() + "\n"
		}
		fmt.Println(paths)
		msgEntry.SetText(paths)
	})
	//title := canvas.NewText("往窗口上拖拉文件演示，滑块控件点击演示", theme.PrimaryColor())
	title := canvas.NewText("desktop", theme.PrimaryColor())
	top := container.NewVBox(msgEntry, msgEntry1)
	win.SetContent(container.NewBorder(title, slider, nil, nil, top))
	win.SetMaster()
	win.Resize(fyne.NewSize(800, 700))
	win.ShowAndRun()
}
