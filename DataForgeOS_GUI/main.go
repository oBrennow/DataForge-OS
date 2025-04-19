package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/oBrenn0w/Data-Forge-OS/DataForgeOS_GUI/menu"
	"github.com/oBrenn0w/Data-Forge-OS/DataForgeOS_GUI/window"
)

func main() {
	myApp := app.New()
	win := window.CreateMainWindow(myApp)
	win.SetMainMenu(menu.NewAppMenu(myApp))
	win.ShowAndRun()

}
