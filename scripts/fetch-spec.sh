#!/bin/bash
set -e

SPEC_VERSION="${1:-main}"
BASE_URL="https://raw.githubusercontent.com/vrchatapi/specification/${SPEC_VERSION}/openapi"

echo "Fetching VRChat API specification (version: ${SPEC_VERSION})..."

# Create directories
mkdir -p api/components/schemas

# メイン仕様ファイル取得
curl -fsSL -o api/openapi.yaml "${BASE_URL}/openapi.yaml"

# External reference files
curl -fsSL -o api/components/paths.yaml "${BASE_URL}/components/paths.yaml"
curl -fsSL -o api/components/securitySchemes.yaml "${BASE_URL}/components/securitySchemes.yaml"
curl -fsSL -o api/components/tags.yaml "${BASE_URL}/components/tags.yaml"
curl -fsSL -o api/components/schemas/_index.yaml "${BASE_URL}/components/schemas/_index.yaml"

echo "✓ OpenAPI specification fetched successfully"
