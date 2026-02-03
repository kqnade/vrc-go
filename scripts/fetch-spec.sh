#!/bin/bash
set -e

SPEC_VERSION="${1:-main}"
BASE_URL="https://raw.githubusercontent.com/vrchatapi/specification/${SPEC_VERSION}/openapi"

echo "Fetching VRChat API specification (version: ${SPEC_VERSION})..."

# メイン仕様ファイル取得
curl -fsSL -o api/openapi.yaml "${BASE_URL}/openapi.yaml"

echo "✓ OpenAPI specification fetched successfully"
