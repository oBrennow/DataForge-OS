package menu

import (
	"fyne.io/fyne/v2"
	"github.com/oBrenn0w/Data-Forge-OS/DataForgeOS_GUI/modules/producao"
)

func NewAppMenu(a fyne.App) *fyne.MainMenu {
	producaoTurno := fyne.NewMenuItem("Produção", func() {
		wCadastroProducao := producao.CadastroProducao(a)
		wCadastroProducao.Show()
	})
	cadastroMaquina := fyne.NewMenuItem("Cadastro de Maquinas", func() {
		// Criar um subtelinha para r3ealizar o cadastro de maquinas
	})

	cadastroVendas := fyne.NewMenuItem("Cadastro de vendas", func() {
		// Criar um subtelinha para r3ealizar o cadastro das vendas
	})

	relatorioVendas := fyne.NewMenuItem("Relatorio de vendas", func() {
		// Criar um subtelinha para r3ealizar o cadastro das vendas
	})

	cadastroProduto := fyne.NewMenuItem("Cadastro de Produto", func() {
		// Criar um subtelinha para r3ealizar o cadastro das vendas
	})

	relatorioSaidaProdutos := fyne.NewMenuItem("Saida de Produtos", func() {
		// Criar um subtelinha para r3ealizar o cadastro das vendas
	})

	relatorioEntradaProdutos := fyne.NewMenuItem("Entrada de Produtos", func() {
		// Criar um subtelinha para r3ealizar o cadastro das vendas
	})

	menuProducao := fyne.NewMenu("Produção", producaoTurno, cadastroMaquina)
	menuVendas := fyne.NewMenu("Vendas", cadastroVendas, relatorioVendas)
	menuProdutos := fyne.NewMenu("Produtos", cadastroProduto, relatorioSaidaProdutos, relatorioEntradaProdutos)

	return fyne.NewMainMenu(menuProducao, menuVendas, menuProdutos)
}
