package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func AppsDelete(pos []string) {
	id := ResolveID(pos, "apps delete")
	fmt.Println()
	ui.Warn(fmt.Sprintf("Você está prestes a %sDELETAR PERMANENTEMENTE%s o app %s%s%s.",
		ui.Red+ui.Bold, ui.Reset+ui.Yellow, ui.Bold, id, ui.Reset+ui.Yellow))
	ui.Warn("Esta ação NÃO pode ser desfeita.")
	if !Confirm("Digite o ID do app para confirmar", id) {
		ui.Fail("Confirmação incorreta. Operação cancelada.")
		return
	}
	spin := ui.NewSpinner("Deletando app...")
	var r api.MessageResp
	data, err := NewClient().Post("/apps/"+id+"/delete", nil)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	_ = api.Decode(data, &r)
	msg := "App deletado permanentemente."
	if r.Message != "" {
		msg = r.Message
	}
	spin.Stop(msg)
}
