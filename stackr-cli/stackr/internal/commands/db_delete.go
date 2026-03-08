package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/ui"
)

func DBDelete(pos []string) {
	if len(pos) == 0 {
		ui.Fail("Informe o ID: stackr db delete <ID>")
		return
	}
	id := pos[0]
	fmt.Println()
	ui.Warn(fmt.Sprintf("Você está prestes a %sDELETAR PERMANENTEMENTE%s o database %s%s%s.",
		ui.Red+ui.Bold, ui.Reset+ui.Yellow, ui.Bold, id, ui.Reset+ui.Yellow))
	ui.Warn("TODOS OS DADOS SERÃO PERDIDOS.")
	if !Confirm("Digite o ID para confirmar", id) {
		ui.Fail("Confirmação incorreta. Operação cancelada.")
		return
	}
	spin := ui.NewSpinner("Removendo database...")
	_, err := NewClient().Delete("/databases/" + id)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("Database deletado.")
}
