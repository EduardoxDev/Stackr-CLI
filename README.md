<div align="center">

```
███████╗████████╗ █████╗  ██████╗██╗  ██╗██████╗
██╔════╝╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝██╔══██╗
███████╗   ██║   ███████║██║     █████╔╝ ██████╔╝
╚════██║   ██║   ██╔══██║██║     ██╔═██╗ ██╔══██╗
███████║   ██║   ██║  ██║╚██████╗██║  ██╗██║  ██║
╚══════╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝
```

**A CLI oficial da plataforma [Stackr](https://stackr.lat) — deploy rápido, simples e sem frescura.**

[![npm version](https://img.shields.io/npm/v/stackr-lat?color=00ff88&style=flat-square)](https://www.npmjs.com/package/stackr-lat)
[![downloads](https://img.shields.io/npm/dm/stackr-lat?color=00aaff&style=flat-square)](https://www.npmjs.com/package/stackr-lat)
[![license](https://img.shields.io/npm/l/stackr-lat?color=ff6b6b&style=flat-square)](LICENSE)
[![platform](https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-lightgrey?style=flat-square)](#instalação)

</div>

---

## ✨ O que é o Stackr CLI?

O **Stackr CLI** é a interface de linha de comando oficial da plataforma Stackr. Com ele você gerencia seus apps, faz deploys, acompanha logs e muito mais — tudo direto do terminal, sem precisar abrir o painel.

---

## ⚡ Instalação

> Requisito: [Node.js](https://nodejs.org) v16 ou superior

```bash
npm install -g stackr-lat
```

Isso é tudo. O instalador detecta seu sistema e baixa o binário correto automaticamente.

| Sistema Operacional | Binário |
|---|---|
| Windows (x64) | `stackr-windows.exe` |
| macOS (Apple Silicon) | `stackr-darwin-arm64` |
| macOS (Intel) | `stackr-darwin-amd64` |
| Linux (x64) | `stackr-linux-amd64` |
| Linux (ARM) | `stackr-linux-arm64` |

---

## 🔐 Autenticação

Após instalar, faça login com seu token da Stackr:

```bash
stackr login <SEU_TOKEN>
```

Seu token é salvo localmente em `~/.stackr/config.json` e usado automaticamente em todos os comandos.

---

## 🚀 Comandos

### Apps

```bash
# Listar todos os seus apps
stackr apps list

# Ver detalhes de um app específico
stackr apps info <app-id>

# Fazer deploy de um app
stackr deploy

# Reiniciar um app
stackr apps restart <app-id>

# Parar um app
stackr apps stop <app-id>
```

### Logs

```bash
# Ver logs em tempo real
stackr logs <app-id>

# Ver as últimas N linhas de log
stackr logs <app-id> --tail 100
```

### Geral

```bash
# Ver versão instalada
stackr --version

# Ver ajuda
stackr --help

# Ver ajuda de um subcomando
stackr apps --help
```

---

## 💡 Exemplo de uso

```bash
# 1. Instale a CLI
npm install -g stackr-lat

# 2. Faça login
stackr login meu-token-aqui

# 3. Liste seus apps
stackr apps list

# Seus Apps (3)
# ─────────────────────────────────────────────────────
#   ID                                    NOME          STATUS
#   8de3ea25-xxxx-xxxx-xxxx-xxxxxxxxxxxx  meu-api       ● online
#   6f2aaa8f-xxxx-xxxx-xxxx-xxxxxxxxxxxx  meu-bot       ● online

# 4. Deploy
stackr deploy
```

---

## 🛠️ Problemas comuns

### `stackr` não é reconhecido como comando (Windows)

O PATH pode não ter sido atualizado. Feche e abra um novo terminal. Se persistir:

```cmd
setx PATH "%PATH%;%APPDATA%\npm"
```

### Erro de download na instalação

Se o instalador falhar ao baixar o binário, baixe manualmente em:
👉 [github.com/EduardoxDev/Stackr-CLI/releases](https://github.com/EduardoxDev/Stackr-CLI/releases)

Coloque o arquivo em:
```
%APPDATA%\npm\node_modules\stackr-lat\bin\stackr.exe
```

### `Esta versão não é compatível com o Windows`

Seu binário pode estar corrompido. Reinstale:
```bash
npm uninstall -g stackr-lat
npm cache clean --force
npm install -g stackr-lat
```

---

## 📦 Sobre

- **Versão atual:** `v1.2.0`
- **Plataforma:** [stackr.lat](https://stackr.lat)
- **npm:** [npmjs.com/package/stackr-lat](https://www.npmjs.com/package/stackr-lat)
- **Repositório:** [github.com/EduardoxDev/Stackr-CLI](https://github.com/EduardoxDev/Stackr-CLI)

---

## 📄 Licença

MIT © [Stackr](https://stackr.lat)
