package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/ui"
)

func FilesMkdir(pos []string, flags map[string]string) {
	id := ResolveID(pos, "files mkdir")
	path := Flag(flags, "path", "")
	if path == "" {
		ui.Fail("Informe o caminho com --path")
		return
	}
	spin := ui.NewSpinner("Criando " + path + "...")
	_, err := NewClient().Post(
		fmt.Sprintf("/apps/%s/files/mkdir", id),
		map[string]string{"path": path},
	)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("Pasta criada: " + path)
}
