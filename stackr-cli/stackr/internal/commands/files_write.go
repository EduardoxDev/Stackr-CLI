package commands

import (
	"os"

	"github.com/stackr-lat/cli/internal/ui"
)

func FilesWrite(pos []string, flags map[string]string) {
	id := ResolveID(pos, "files write")
	path := Flag(flags, "path", "")
	if path == "" {
		ui.Fail("Informe o caminho com --path")
		return
	}
	content := Flag(flags, "content", "")
	if content == "" {
		if localFile := Flag(flags, "file", ""); localFile != "" {
			raw, err := os.ReadFile(localFile)
			if err != nil {
				ui.Fail("Não foi possível ler o arquivo local: " + err.Error())
				return
			}
			content = string(raw)
		}
	}
	if content == "" {
		ui.Fail("Informe o conteúdo com --content ou --file <arquivo-local>")
		return
	}
	spin := ui.NewSpinner("Escrevendo " + path + "...")
	_, err := NewClient().Put(
		"/apps/"+id+"/files/content",
		map[string]string{"path": path, "content": content},
	)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("Arquivo escrito: " + path)
}
