#!/bin/bash

# Set strict mode for better error handling and security
set -euo pipefail

# Define the input OpenAPI specification file
SPEC_FILE="openapi.json"

# Define output directories for clients and documentation
OUTPUT_DIR_TS="./reference/code/typescript"
OUTPUT_DIR_JS="./reference/code/javascript"
OUTPUT_DIR_BASH="./reference/code/bash"
OUTPUT_DIR_RUST="./reference/code/rust"
OUTPUT_DIR_PY="./reference/code/python"
OUTPUT_DIR_GO="./reference/code/go"
OUTPUT_DIR_MD="./reference/"  # Markdown documentation output directory

# Check if OpenAPI Generator CLI is installed
if ! command -v openapi-generator-cli &> /dev/null; then
    echo "openapi-generator-cli could not be found, please install it to proceed."
    exit 1
fi

# Generate Markdown documentation
echo "Generating Markdown documentation..."
openapi-generator-cli generate -i "$SPEC_FILE" -g markdown -o "$OUTPUT_DIR_MD"

# Function to generate API client
generate_client() {
    local lang="$1"
    local generator="$2"
    local output_dir="$3"

    echo "Generating $lang client..."
    if ! openapi-generator-cli generate -i "$SPEC_FILE" -g "$generator" -o "$output_dir"; then
        echo "Failed to generate $lang client, stopping script."
        exit 1
    fi
}

# Generate clients for each language
generate_client "TypeScript" "typescript-node" "$OUTPUT_DIR_TS"
generate_client "JavaScript" "javascript" "$OUTPUT_DIR_JS"
generate_client "Bash" "bash" "$OUTPUT_DIR_BASH"
generate_client "Rust" "rust" "$OUTPUT_DIR_RUST"
generate_client "Python" "python" "$OUTPUT_DIR_PY"
generate_client "Go" "go" "$OUTPUT_DIR_GO"

echo "All clients and documentation generated successfully."
