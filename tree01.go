package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	a := app.NewWithID("")
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("sfd")
	dirTree(w)
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}

func dirTree(w fyne.Window) {
	lblMsg := widget.NewLabel("")
	tree := widget.NewTree(nil, nil, nil, nil)
	tree.OnSelected = func(uid widget.TreeNodeID) {
		lblMsg.SetText(uid)
	}
	root := "c:/"
	rootTxt := widget.NewEntry()
	rootTxt.SetText(root)
	btn := widget.NewButton("Browse dir", func() {
		refreshTree(tree, rootTxt.Text)
	})
	top := container.NewVBox(lblMsg, rootTxt, btn)
	c := container.NewBorder(top, nil, nil, nil, tree)
	w.SetContent(c)
}

func refreshTree(tree *widget.Tree, root string) {
	tree.Root = root
	tree.ChildUIDs = func(uid widget.TreeNodeID) (c []widget.TreeNodeID) {
		if uid == "" {
			c = getFileList(root)
		} else {
			c = getFileList(uid)
		}
		return
	}
	tree.CreateNode = func(branch bool) (o fyne.CanvasObject) {
		return widget.NewLabel("")
	}
	tree.UpdateNode = func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		lbl := node.(*widget.Label)
		file, err := os.Stat(uid)
		if err != nil {
			lbl.SetText("")
		}
		lbl.SetText(file.Name())
	}
	tree.IsBranch = func(uid widget.TreeNodeID) (ok bool) {
		return isDir(uid)
	}
	tree.Refresh()
}
func isDir(pth string) bool {
	fi, err := os.Stat(pth)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func getFileList(pth string) (lst []string) {
	fi, err := ioutil.ReadDir(pth)
	if err != nil {
		return
	}
	for _, file := range fi {
		path1 := pth + string(filepath.Separator) + file.Name()
		lst = append(lst, path1)
	}
	return
}
