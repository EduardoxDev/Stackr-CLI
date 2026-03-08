package commands

import (
	"fmt"
	"strings"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func AppsLogs(pos []string, flags map[string]string) {
	id := ResolveID(pos, "apps logs")
	tail := FlagInt(flags, "tail", 100)
	spin := ui.NewSpinner(fmt.Sprintf("Buscando últimas %d linhas...", tail))
	var logs api.AppLogs
	data, err := NewClient().Get(fmt.Sprintf("/apps/%s/logs?tail=%d", id, tail))
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &logs); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("")
	ui.Header(fmt.Sprintf("Logs  %s(últimas %d linhas)%s", ui.Gray, tail, ui.Reset))
	if len(logs.Logs) == 0 {
		ui.Info("Nenhum log disponível.")
		return
	}
	for _, line := range logs.Logs {
		low := strings.ToLower(line)
		switch {
		case strings.Contains(low, "error") || strings.Contains(low, "fatal") || strings.Contains(low, "panic"):
			fmt.Println("  " + ui.Red + line + ui.Reset)
		case strings.Contains(low, "warn"):
			fmt.Println("  " + ui.Yellow + line + ui.Reset)
		case strings.Contains(low, "[info]") || strings.Contains(low, "info:"):
			fmt.Println("  " + ui.Cyan + line + ui.Reset)
		case strings.Contains(low, "debug"):
			fmt.Println("  " + ui.Magenta + line + ui.Reset)
		default:
			fmt.Println("  " + ui.Gray + line + ui.Reset)
		}
	}
	fmt.Println()
}
