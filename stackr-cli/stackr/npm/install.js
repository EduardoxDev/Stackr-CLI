const { execSync } = require('child_process')
const https = require('https')
const fs = require('fs')
const path = require('path')
const os = require('os')

const VERSION = '1.2.0'
const BASE_URL = `https://github.com/stackr-lat/cli/releases/download/v${VERSION}`

function getBinaryName() {
  const platform = os.platform()
  const arch = os.arch()
  if (platform === 'win32') return 'stackr-windows.exe'
  if (platform === 'darwin') return arch === 'arm64' ? 'stackr-darwin-arm64' : 'stackr-darwin-amd64'
  return arch === 'arm64' ? 'stackr-linux-arm64' : 'stackr-linux-amd64'
}

function download(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest)
    https.get(url, (res) => {
      if (res.statusCode === 302 || res.statusCode === 301) {
        file.close()
        return download(res.headers.location, dest).then(resolve).catch(reject)
      }
      res.pipe(file)
      file.on('finish', () => { file.close(); resolve() })
    }).on('error', (err) => { fs.unlink(dest, () => {}); reject(err) })
  })
}

async function main() {
  const binaryName = getBinaryName()
  const url = `${BASE_URL}/${binaryName}`
  const binDir = path.join(__dirname, 'bin')
  const dest = path.join(binDir, os.platform() === 'win32' ? 'stackr.exe' : 'stackr')

  fs.mkdirSync(binDir, { recursive: true })

  process.stdout.write(`  Baixando stackr CLI v${VERSION}...\n`)
  await download(url, dest)

  if (os.platform() !== 'win32') {
    fs.chmodSync(dest, 0o755)
  }

  process.stdout.write(`  ✔ stackr CLI instalado com sucesso!\n`)
  process.stdout.write(`  → Execute: stackr login <SEU_TOKEN>\n\n`)
}

main().catch((err) => {
  console.error('Erro na instalação:', err.message)
  process.exit(1)
})
