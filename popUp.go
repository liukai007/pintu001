package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func popUp(w fyne.Window) {
	lblMsg := widget.NewLabel("")
	var mpu *widget.PopUp
	mpuc := container.NewVBox(
		widget.NewLabel("modal pop up label"),
		widget.NewButton("modal pop up btn", func() {
			lblMsg.SetText("Click pop up btn")
		}),
		widget.NewButton("close", func() {
			mpu.Hide()
		}),
	)
	mpu = widget.NewModalPopUp(mpuc, w.Canvas())
	btn1 := widget.NewButton("show modal pop up", func() {
		mpu.Show()
	})
	puc := container.NewVBox(
		widget.NewLabel("pop up label"),
		widget.NewButton("Pop up btn", func() {

		}),
	)
	pu := widget.NewPopUp(puc, w.Canvas())
	btn := widget.NewButton("show popup", func() {
		pu.Show()
	})
	c := container.NewVBox(lblMsg, btn, btn1)
	w.SetContent(c)
}
