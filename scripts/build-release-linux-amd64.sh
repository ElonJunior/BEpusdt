#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

BASE_VERSION="$(git describe --tags --abbrev=0 2>/dev/null || true)"
if [[ -z "$BASE_VERSION" ]]; then
  BASE_VERSION="v0.0.0"
fi

COMMIT_DATE="$(git show -s --date=format:%Y%m%d --format=%cd HEAD)"
COMMIT_SHORT="$(git rev-parse --short HEAD)"
FULL_VERSION="${BASE_VERSION}-${COMMIT_DATE}-${COMMIT_SHORT}"

OUT_DIR="dist/linux-amd64"
PKG_DIR="dist/linux-amd64-package/bepusdt"
BIN_PATH="${OUT_DIR}/bepusdt"
TAR_PATH="dist/linux-amd64-BEpusdt.tar.gz"

mkdir -p "$OUT_DIR" "$PKG_DIR"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -trimpath -ldflags "-s -w -X github.com/v03413/bepusdt/app.Version=${FULL_VERSION}" -o "$BIN_PATH" ./main

cp "$BIN_PATH" "$PKG_DIR/bepusdt"
cp "docs/bepusdt.service" "$PKG_DIR/bepusdt.service"
tar -czf "$TAR_PATH" -C "dist/linux-amd64-package" "bepusdt"

echo "Build complete"
echo "Version: ${FULL_VERSION}"
echo "Binary : ${BIN_PATH}"
echo "Package: ${TAR_PATH}"
