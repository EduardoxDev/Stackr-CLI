package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func AppsStats(pos []string) {
	id := ResolveID(pos, "apps stats")
	spin := ui.NewSpinner("Coletando métricas...")
	var s api.AppStats
	data, err := NewClient().Get("/apps/" + id + "/stats")
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &s); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("")
	ui.Header("Métricas em Tempo Real")
	onOff := ui.Red + "● Offline" + ui.Reset
	if s.Running {
		onOff = ui.Green + "● Online" + ui.Reset
	}
	ui.Label("Status", onOff)
	ui.Label("CPU", s.CPU)
	ui.Label("RAM", s.RAM)
	ui.Label("Rede ↓ RX", s.NetworkRx)
	ui.Label("Rede ↑ TX", s.NetworkTx)
	fmt.Println()
}
