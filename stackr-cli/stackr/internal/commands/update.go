package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/stackr-lat/cli/internal/config"
	"github.com/stackr-lat/cli/internal/ui"
)

func Update() {
	ui.Header("Atualizar Stackr CLI")
	ui.Info(fmt.Sprintf("Versão instalada: %sv%s%s", ui.Bold, config.Version, ui.Reset))

	spin := ui.NewSpinner("Verificando última versão...")
	resp, err := http.Get(config.ReleaseLatestURL)
	if err != nil {
		spin.Fail("Não foi possível conectar ao GitHub: " + err.Error())
		return
	}
	defer resp.Body.Close()

	parts := strings.Split(resp.Request.URL.String(), "/")
	latestTag := parts[len(parts)-1]
	latestVer := strings.TrimPrefix(latestTag, "v")

	if latestVer == config.Version {
		spin.Stop(fmt.Sprintf("Você já está na versão mais recente! %s(v%s)%s", ui.Green, config.Version, ui.Reset))
		return
	}
	spin.Stop(fmt.Sprintf("Nova versão: %sv%s%s → %s%s%s",
		ui.Gray, config.Version, ui.Reset, ui.Green+ui.Bold, latestTag, ui.Reset))

	goos, goarch := runtime.GOOS, runtime.GOARCH
	var binaryName string
	switch {
	case goos == "windows":
		binaryName = "stackr-windows.exe"
	case goos == "darwin" && goarch == "arm64":
		binaryName = "stackr-darwin-arm64"
	case goos == "darwin":
		binaryName = "stackr-darwin-amd64"
	case goarch == "arm64":
		binaryName = "stackr-linux-arm64"
	default:
		binaryName = "stackr-linux-amd64"
	}

	spin2 := ui.NewSpinner(fmt.Sprintf("Baixando %s...", binaryName))
	dlResp, err := http.Get(fmt.Sprintf("%s/%s/%s", config.ReleaseBaseURL, latestTag, binaryName))
	if err != nil {
		spin2.Fail("Erro no download: " + err.Error())
		return
	}
	defer dlResp.Body.Close()
	if dlResp.StatusCode != 200 {
		spin2.Fail(fmt.Sprintf("Download falhou (HTTP %d)", dlResp.StatusCode))
		ui.Hint("Baixe manualmente em: https://github.com/stackr-lat/cli/releases")
		return
	}

	tmp, err := os.CreateTemp("", "stackr-update-*")
	if err != nil {
		spin2.Fail("Erro ao criar arquivo temporário: " + err.Error())
		return
	}
	defer os.Remove(tmp.Name())
	if _, err := io.Copy(tmp, dlResp.Body); err != nil {
		spin2.Fail("Erro ao salvar binário: " + err.Error())
		return
	}
	tmp.Close()
	os.Chmod(tmp.Name(), 0755)
	spin2.Stop("Download completo!")

	selfPath, err := exec.LookPath("stackr")
	if err != nil {
		selfPath, _ = os.Executable()
	}

	spin3 := ui.NewSpinner("Instalando...")
	if goos == "windows" {
		dest := selfPath + ".new.exe"
		if err := os.Rename(tmp.Name(), dest); err != nil {
			spin3.Fail("Substitua manualmente " + tmp.Name() + " por " + selfPath)
			return
		}
		spin3.Stop("Pronto! Substitua " + selfPath + " por " + dest)
		return
	}
	if err := os.Rename(tmp.Name(), selfPath); err != nil {
		spin3.Stop("")
		ui.Warn("Requer sudo para instalar em " + selfPath)
		cmd := exec.Command("sudo", "mv", tmp.Name(), selfPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			ui.Fail("Falha na instalação.")
			ui.Hint(fmt.Sprintf("Instale manualmente: sudo mv %s %s", tmp.Name(), selfPath))
			return
		}
	}
	spin3.Stop(fmt.Sprintf("CLI atualizada para %s%s%s! 🚀", ui.Green+ui.Bold, latestTag, ui.Reset))
}
