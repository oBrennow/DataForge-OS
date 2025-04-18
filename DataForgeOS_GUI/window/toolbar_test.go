package window

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func TestFormToolbar(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	w := a.NewWindow("FormToolbar Test")
	defer w.Close()

	actions := FormActions{
		Include: func() {},
		Search:  func() {},
		Delete:  func() {},
		Confirm: func() {},
		Cancel:  func() {},
	}

	tb := FormToolbar(w, actions)

	if tb == nil {
		t.Fatal("FormToolbar retornou nil — pipeline interrompido")
	}
	if len(tb.Items) != 5 {
		t.Fatalf("Número de itens = %d; esperava 5", len(tb.Items))
	}

	for i, it := range tb.Items {
		if _, ok := it.(*widget.ToolbarAction); !ok {
			t.Errorf("Item[%d] = %T; esperava *widget.ToolbarAction", i, it)
		}
	}
}
