package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func DBStats(pos []string) {
	if len(pos) == 0 {
		ui.Fail("Informe o ID: stackr db stats <ID>")
		return
	}
	spin := ui.NewSpinner("Coletando métricas...")
	var s api.DBStats
	data, err := NewClient().Get("/databases/" + pos[0] + "/stats")
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &s); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("")
	ui.Header("Métricas do Database")
	onOff := ui.Red + "● Offline" + ui.Reset
	if s.Running {
		onOff = ui.Green + "● Online" + ui.Reset
	}
	ui.Label("Status", onOff)
	ui.Label("CPU", s.CPUPercent)
	ui.Label("Memória", s.MemUsage)
	ui.Label("Mem %", s.MemPercent)
	ui.Label("Rede ↓ RX", s.NetRx)
	ui.Label("Rede ↑ TX", s.NetTx)
	ui.Label("Processos", fmt.Sprintf("%d PIDs", s.PIDs))
	fmt.Println()
}
