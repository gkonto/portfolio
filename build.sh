#!/usr/bin/env bash

set -euo pipefail

go run ./cmd/web -export-dir dist
