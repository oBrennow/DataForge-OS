package producao_test

import (
	"reflect"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"

	"github.com/oBrenn0w/Data-Forge-OS/DataForgeOS_GUI/modules/producao"
)

func TestNewCadastroProducaoWindow_Structure(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	w := producao.CadastroProducao(a)

	if w.Title() != "Cadastro de Produção por turno" {
		t.Errorf("titulo = %q; esperava %q", w.Title(), "Cadastro de Produção")
	}

	border, ok := w.Content().(*fyne.Container)
	if !ok {
		t.Fatalf("esperava *fyne.Container, mas foi %T", w.Content())
	}

	tb, ok := border.Objects[0].(*widget.Toolbar)
	if !ok {
		t.Fatalf("esperava toolbar no topo, mas é %T", border.Objects[0])
	}
	if len(tb.Items) != 5 {
		t.Errorf("toolbar.Items = %d; esperava 5", len(tb.Items))
	}

	formCnt, ok := border.Objects[1].(*fyne.Container)
	if !ok {
		t.Fatalf("esperava container de formulário, mas é %T", border.Objects[1])
	}

	found := map[string]bool{"Label": false, "Entry": false, "Select": false}
	for _, obj := range formCnt.Objects {
		switch reflect.TypeOf(obj) {
		case reflect.TypeOf(&widget.Label{}):
			found["Label"] = true
		case reflect.TypeOf(&widget.Entry{}):
			found["Entry"] = true
		case reflect.TypeOf(&widget.Select{}):
			found["Select"] = true
		}
	}
	for typ, ok := range found {
		if !ok {
			t.Errorf("esperava pelo menos um %s no formulário", typ)
		}
	}
}
