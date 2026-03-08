package commands

import (
	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func appsAction(pos []string, cmd, spinMsg, successMsg string) {
	id := ResolveID(pos, cmd)
	spin := ui.NewSpinner(spinMsg)
	var r api.MessageResp
	data, err := NewClient().Post("/apps/"+id+"/"+cmd, nil)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	_ = api.Decode(data, &r)
	msg := successMsg
	if r.Message != "" {
		msg = r.Message
	}
	spin.Stop(msg)
}

func AppsStart(pos []string)   { appsAction(pos, "start", "Iniciando app...", "App iniciado!") }
func AppsStop(pos []string)    { appsAction(pos, "stop", "Parando app...", "App parado.") }
func AppsRestart(pos []string) { appsAction(pos, "restart", "Reiniciando app...", "App reiniciando...") }

func AppsRebuild(pos []string) {
	id := ResolveID(pos, "apps rebuild")
	ui.Warn("O app ficará offline durante o rebuild.")
	spin := ui.NewSpinner("Reconstruindo container...")
	var r api.MessageResp
	data, err := NewClient().Post("/apps/"+id+"/rebuild", nil)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	_ = api.Decode(data, &r)
	msg := "Rebuild iniciado!"
	if r.Message != "" {
		msg = r.Message
	}
	spin.Stop(msg)
	ui.Hint("Acompanhe: stackr apps logs " + id)
}
