#!/usr/bin/env bash
set -euo pipefail

REPO="stackr-lat/cli"
BINARY="stackr"
INSTALL_DIR="/usr/local/bin"

R='\033[0;31m'; G='\033[0;32m'; Y='\033[1;33m'
C='\033[0;36m'; B='\033[1m'; D='\033[2m'; Z='\033[0m'

ok()   { echo -e "${G}${B}✔${Z}  $1"; }
fail() { echo -e "${R}${B}✖${Z}  $1" >&2; exit 1; }
info() { echo -e "${C}ℹ${Z}  $1"; }
warn() { echo -e "${Y}⚠${Z}  $1"; }
step() { echo -e "\n${B}$1${Z}"; }

echo ""
echo -e "${B}${C}  ███████╗████████╗ █████╗  ██████╗██╗  ██╗██████╗ ${Z}"
echo -e "${B}${C}  ██╔════╝╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝██╔══██╗${Z}"
echo -e "${B}${C}  ███████╗   ██║   ███████║██║     █████╔╝ ██████╔╝ ${Z}"
echo -e "${B}${C}  ╚════██║   ██║   ██╔══██║██║     ██╔═██╗ ██╔══██╗ ${Z}"
echo -e "${B}${C}  ███████║   ██║   ██║  ██║╚██████╗██║  ██╗██║  ██║ ${Z}"
echo -e "${B}${C}  ╚══════╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝${Z}"
echo -e "${D}  Instalador oficial · https://stackr.lat${Z}"
echo ""

step "① Detectando plataforma..."

OS="$(uname -s 2>/dev/null || echo unknown)"
ARCH="$(uname -m 2>/dev/null || echo unknown)"

case "$OS" in
  Linux*)  PLATFORM="linux" ;;
  Darwin*) PLATFORM="darwin" ;;
  MINGW*|MSYS*|CYGWIN*) PLATFORM="windows" ;;
  *) fail "Sistema operacional não suportado: $OS" ;;
esac

case "$ARCH" in
  x86_64|amd64)  ARCH_TAG="amd64" ;;
  arm64|aarch64) ARCH_TAG="arm64" ;;
  *) fail "Arquitetura não suportada: $ARCH" ;;
esac

if [ "$PLATFORM" = "windows" ]; then
  BINARY_FILE="${BINARY}-windows.exe"
else
  BINARY_FILE="${BINARY}-${PLATFORM}-${ARCH_TAG}"
fi

ok "Plataforma: ${B}${PLATFORM}/${ARCH_TAG}${Z}"

if ! command -v curl &>/dev/null; then
  fail "curl não encontrado. Instale curl e tente novamente."
fi

step "② Verificando versão disponível..."

LATEST_URL=$(curl -fsSL -o /dev/null -w '%{url_effective}' \
  "https://github.com/${REPO}/releases/latest" 2>/dev/null) || \
  fail "Não foi possível acessar o GitHub."

LATEST_TAG="${LATEST_URL##*/}"
[ -z "$LATEST_TAG" ] && fail "Não foi possível determinar a versão mais recente."

ok "Versão: ${B}${LATEST_TAG}${Z}"

step "③ Baixando binário..."

DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${BINARY_FILE}"
TMP_FILE="$(mktemp /tmp/stackr-XXXXXX)"
trap 'rm -f "$TMP_FILE"' EXIT

HTTP_CODE=$(curl -fsSL --progress-bar -o "$TMP_FILE" -w "%{http_code}" "$DOWNLOAD_URL") || true
[ "$HTTP_CODE" != "200" ] && fail "Download falhou (HTTP ${HTTP_CODE}). Veja: https://github.com/${REPO}/releases"

chmod +x "$TMP_FILE"
ok "Download completo!"

step "④ Instalando..."

DEST="${INSTALL_DIR}/${BINARY}"

if [ -w "$INSTALL_DIR" ]; then
  mv "$TMP_FILE" "$DEST"
elif command -v sudo &>/dev/null; then
  sudo mv "$TMP_FILE" "$DEST"
else
  ALT="${HOME}/.local/bin"
  mkdir -p "$ALT"
  mv "$TMP_FILE" "${ALT}/${BINARY}"
  DEST="${ALT}/${BINARY}"
  warn "Instalado em ${B}${ALT}${Z}"
  [[ ":$PATH:" != *":${ALT}:"* ]] && warn "Adicione ao PATH: export PATH=\"\$HOME/.local/bin:\$PATH\""
fi

step "⑤ Verificando instalação..."

if [ -x "$DEST" ]; then
  ok "${B}Stackr CLI instalada com sucesso!${Z}"
  ok "Local: ${B}${DEST}${Z}"
  echo ""
  echo -e "  ${B}Próximos passos:${Z}"
  echo -e "  ${C}1.${Z} Gere seu token:  ${B}https://stackr.lat/dashboard/settings${Z}"
  echo -e "  ${C}2.${Z} Autentique:      ${B}stackr login sk_live_xxx${Z}"
  echo -e "  ${C}3.${Z} Liste seus apps: ${B}stackr apps list${Z}"
  echo ""
else
  fail "Instalação falhou."
fi
