package menu_test

import (
	"reflect"
	"testing"

	"fyne.io/fyne/v2/app"
	"github.com/oBrenn0w/Data-Forge-OS/DataForgeOS_GUI/menu"
)

func TestMainMenuiItens(t *testing.T) {
	win := app.New().NewWindow("DataForgeOS")
	win.SetMainMenu(menu.NewAppMenu())

	menu := win.MainMenu()
	titles := []string{}

	for _, m := range menu.Items {
		titles = append(titles, m.Label)
	}

	esperado := []string{"Produção", "Vendas", "Produtos"}
	if !reflect.DeepEqual(titles, esperado) {
		t.Errorf("esperado menu %v, got %v", esperado, titles)
	}
}
