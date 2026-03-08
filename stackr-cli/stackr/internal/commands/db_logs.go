package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func DBLogs(pos []string, flags map[string]string) {
	if len(pos) == 0 {
		ui.Fail("Informe o ID: stackr db logs <ID>")
		return
	}
	tail := FlagInt(flags, "tail", 100)
	spin := ui.NewSpinner("Buscando logs...")
	var logs api.DBLogs
	data, err := NewClient().Get(fmt.Sprintf("/databases/%s/logs?tail=%d", pos[0], tail))
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &logs); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("")
	ui.Header(fmt.Sprintf("Logs do Database  %s(últimas %d linhas)%s", ui.Gray, tail, ui.Reset))
	fmt.Println(ui.Gray + logs.Logs + ui.Reset)
	fmt.Println()
}
