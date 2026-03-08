const fs = require('fs')
const path = require('path')
const os = require('os')
const https = require('https')

const VERSION = '1.2.0'
const REPO = 'EduardoxDev/Stackr-CLI'
const BASE_URL = `https://github.com/${REPO}/releases/download/v${VERSION}`

function getBinaryName() {
  const platform = os.platform()
  const arch = os.arch()
  if (platform === 'win32') return 'stackr-windows.exe'
  if (platform === 'darwin') return arch === 'arm64' ? 'stackr-darwin-arm64' : 'stackr-darwin-amd64'
  return arch === 'arm64' ? 'stackr-linux-arm64' : 'stackr-linux-amd64'
}

function getDestName() {
  return os.platform() === 'win32' ? 'stackr.exe' : 'stackr'
}

function download(url, dest, redirects) {
  redirects = redirects || 0
  if (redirects > 10) {
    throw new Error('Muitos redirecionamentos')
  }

  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest)

    https.get(url, { headers: { 'User-Agent': 'stackr-installer' } }, (res) => {
      // Seguir redirecionamentos (301, 302, 307, 308)
      if (res.statusCode >= 300 && res.statusCode < 400 && res.headers.location) {
        file.close()
        fs.unlink(dest, () => {})
        return resolve(download(res.headers.location, dest, redirects + 1))
      }

      if (res.statusCode !== 200) {
        file.close()
        fs.unlink(dest, () => {})
        return reject(new Error(`HTTP ${res.statusCode} ao baixar ${url}`))
      }

      const total = parseInt(res.headers['content-length'] || '0', 10)
      let downloaded = 0

      res.on('data', (chunk) => {
        downloaded += chunk.length
        if (total > 0) {
          const pct = Math.round((downloaded / total) * 100)
          process.stdout.write(`\r  Baixando... ${pct}%`)
        }
      })

      res.pipe(file)

      file.on('finish', () => {
        process.stdout.write('\n')
        file.close(resolve)
      })

      file.on('error', (err) => {
        fs.unlink(dest, () => {})
        reject(err)
      })
    }).on('error', (err) => {
      fs.unlink(dest, () => {})
      reject(err)
    })
  })
}

async function main() {
  const binaryName = getBinaryName()
  const destName = getDestName()
  const url = `${BASE_URL}/${binaryName}`
  const binDir = path.join(__dirname, 'bin')
  const dest = path.join(binDir, destName)

  fs.mkdirSync(binDir, { recursive: true })

  process.stdout.write(`\n  Baixando stackr CLI v${VERSION}...\n`)

  try {
    await download(url, dest)

    if (os.platform() !== 'win32') {
      fs.chmodSync(dest, 0o755)
    }

    const size = fs.statSync(dest).size
    if (size < 1000) {
      throw new Error('Arquivo muito pequeno — download falhou. Tente novamente ou baixe manualmente.')
    }

    process.stdout.write(`  ✔  stackr CLI instalado com sucesso! (${Math.round(size / 1024 / 1024)}MB)\n`)
    process.stdout.write(`  >>  Execute: stackr login <SEU_TOKEN>\n\n`)
  } catch (err) {
    process.stderr.write(`\n  ERRO: ${err.message}\n`)
    process.stderr.write(`  Baixe manualmente: https://github.com/${REPO}/releases\n\n`)
    process.exit(1)
  }
}

main()
