package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Line")

	line := canvas.NewLine(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	line.StrokeWidth = 1
	//line.Move(fyne.NewPos(1, 1))

	line2 := canvas.NewLine(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	line2.StrokeWidth = 1
	//line2.Move(fyne.NewPos(100, 100))

	line3 := canvas.NewLine(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	line3.StrokeWidth = 1

	line4 := canvas.NewLine(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	line4.StrokeWidth = 1

	multilineEntry := widget.NewMultiLineEntry()
	multilineEntry.Wrapping = fyne.TextWrapBreak
	const loren = "The sunset and the solitary bird fly together, and the autumn water is the same color as the sky"
	multilineEntry.SetText(loren)
	multilineEntry.Move(fyne.NewPos(150, 30))
	multilineEntry.Resize(fyne.NewSize(140, 290))
	info := widget.NewLabel("设备调试消息列表")
	info.Move(fyne.NewPos(150, 0))
	info1 := widget.NewLabel("协议组")
	info1.Move(fyne.NewPos(0, 0))

	info2 := widget.NewLabel("命令CODE列表")
	info2.Move(fyne.NewPos(0, 150))

	c := container.NewWithoutLayout(info1, info2, info, multilineEntry)
	cc := container.NewBorder(line, line2, line3, line4, c)

	w.SetContent(cc)

	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}
