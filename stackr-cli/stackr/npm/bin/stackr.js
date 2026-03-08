#!/usr/bin/env node
const { execFileSync } = require('child_process')
const path = require('path')
const os = require('os')

const bin = path.join(__dirname, os.platform() === 'win32' ? 'stackr.exe' : 'stackr')

try {
  execFileSync(bin, process.argv.slice(2), { stdio: 'inherit' })
} catch (err) {
  process.exit(err.status || 1)
}
