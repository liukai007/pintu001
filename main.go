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
	lblMsg := widget.NewLabel("")
	treeData := map[string][]string{
		"":      {"color", "fruit"},
		"color": []string{"red", "blue", "yellow"},
		"fruit": []string{"apple", "b1", "y2"},
		"red":   []string{"red1", "red2", "red3"},
	}
	tree1 := widget.NewTreeWithStrings(treeData)
	//tree1.OpenAllBranches()
	tree1.OpenBranch("red")
	tree1.OnSelected = func(uid widget.TreeNodeID) {
		lblMsg.SetText(uid)
	}
	cc := container.NewBorder(lblMsg, nil, nil, nil, tree1)
	win.SetContent(cc)
	win.SetMaster()
	win.Resize(fyne.NewSize(800, 700))
	win.ShowAndRun()
}
