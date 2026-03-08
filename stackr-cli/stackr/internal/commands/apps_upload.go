package commands

import (
	"fmt"
	"os"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func AppsUpload(pos []string) {
	if len(pos) == 0 {
		ui.Fail("Informe o arquivo .zip: stackr apps upload <arquivo.zip>")
		ui.Hint("O .zip deve conter um arquivo stackr.config na raiz")
		return
	}
	zipPath := pos[0]
	stat, err := os.Stat(zipPath)
	if os.IsNotExist(err) {
		ui.Fail("Arquivo não encontrado: " + zipPath)
		return
	}
	if stat.Size() > 50*1024*1024 {
		ui.Fail(fmt.Sprintf("Arquivo muito grande: %.1f MB (máx 50 MB)", float64(stat.Size())/1024/1024))
		return
	}
	spin := ui.NewSpinner(fmt.Sprintf("Enviando %s (%.1f MB)...", zipPath, float64(stat.Size())/1024/1024))
	data, err := NewClient().UploadZip("/apps/upload", zipPath)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	var r api.MessageResp
	if err := api.Decode(data, &r); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("App enviado com sucesso! 🚀")
	fmt.Println()
	ui.Label("App ID", r.BotID)
	ui.Label("Status", ui.StatusBadge(r.Status))
	if r.Message != "" {
		ui.Label("Mensagem", r.Message)
	}
	fmt.Println()
	ui.Hint("Acompanhe o build: stackr apps logs " + r.BotID)
	fmt.Println()
}
