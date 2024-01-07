package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

/*
go run fyne.io/fyne/v2/cmd/fyne_demo@latest
*/
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs2 := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)
	tabs3 := container.NewAppTabs(
		container.NewTabItem("Tab 3", widget.NewLabel("Hello3")),
		container.NewTabItem("Tab 4", widget.NewLabel("World!3")),
	)
	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", tabs2),
		container.NewTabItem("Tab 2", tabs3),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(800, 700))
	myWindow.ShowAndRun()
}
