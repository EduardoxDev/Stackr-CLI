package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/stackr-lat/cli/internal/ui"
)

func AppsBackup(pos []string, flags map[string]string) {
	id := ResolveID(pos, "apps backup")
	prefix := id
	if len(id) > 8 {
		prefix = id[:8]
	}
	outFile := Flag(flags, "output", fmt.Sprintf("backup-%s.zip", prefix))
	spin := ui.NewSpinner("Gerando backup...")
	body, err := NewClient().StreamGet("/apps/" + id + "/backup")
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	defer body.Close()
	f, err := os.Create(outFile)
	if err != nil {
		spin.Fail("Erro ao criar arquivo: " + err.Error())
		return
	}
	defer f.Close()
	size, err := io.Copy(f, body)
	if err != nil {
		spin.Fail("Erro ao salvar: " + err.Error())
		return
	}
	spin.Stop(fmt.Sprintf("Backup salvo: %s%s%s  %s(%.2f MB)%s",
		ui.Bold, outFile, ui.Reset, ui.Gray, float64(size)/1024/1024, ui.Reset))
	fmt.Println()
}
