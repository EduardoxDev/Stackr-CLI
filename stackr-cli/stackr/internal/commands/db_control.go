package commands

import (
	"github.com/stackr-lat/cli/internal/ui"
)

func dbAction(pos []string, cmd, spinMsg, successMsg string) {
	if len(pos) == 0 {
		ui.Fail("Informe o ID: stackr db " + cmd + " <ID>")
		return
	}
	spin := ui.NewSpinner(spinMsg)
	_, err := NewClient().Post("/databases/"+pos[0]+"/"+cmd, nil)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop(successMsg)
}

func DBStart(pos []string)   { dbAction(pos, "start", "Iniciando database...", "Database iniciado!") }
func DBStop(pos []string)    { dbAction(pos, "stop", "Parando database...", "Database parado.") }
func DBRestart(pos []string) { dbAction(pos, "restart", "Reiniciando database...", "Database reiniciado!") }
