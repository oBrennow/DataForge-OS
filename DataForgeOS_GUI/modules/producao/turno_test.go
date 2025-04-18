package producao_test

import (
	"reflect"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/oBrenn0w/Data-Forge-OS/DataForgeOS_GUI/modules/producao"
)

func TestNewTurnoWindow_ContentHasExpectedWidgets(t *testing.T) {
	// 1) preparamos um app de teste e a janela
	a := test.NewApp()
	defer a.Quit()

	w := producao.NewTurnoWindow(a)

	// 2) o conteúdo deveria ser um *fyne.Container
	cnt, ok := w.Content().(*fyne.Container)
	if !ok {
		t.Fatalf("esperava um *fyne.Container, mas recebi %T", w.Content())
	}

	// 3) percorremos e sinalizamos se encontramos cada tipo
	found := map[string]bool{
		"Label":  false,
		"Entry":  false,
		"Button": false,
	}
	for _, obj := range cnt.Objects {
		switch reflect.TypeOf(obj) {
		case reflect.TypeOf(&widget.Label{}):
			found["Label"] = true
		case reflect.TypeOf(&widget.Entry{}):
			found["Entry"] = true
		case reflect.TypeOf(&widget.Button{}):
			found["Button"] = true
		}
	}

	// 4) verificamos
	for typ, ok := range found {
		if !ok {
			t.Errorf("widget %s não encontrado no conteúdo", typ)
		}
	}

	// 5) e de bônus, checar título
	want := "Produção por Turno"
	if w.Title() != want {
		t.Errorf("Título = %q; queríamos %q", w.Title(), want)
	}
}
