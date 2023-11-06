#!/bin/bash

add_link() {
    if [ $# -ne 2 ]; then
        echo "Usage: add_link <description> <link>"
        return 1
    fi
    
    local description="$1"
    local link="$2"
    
    # Use double quotes for variables to expand them
    curl -X POST -H "Content-Type: application/json" -d "{\"description\": \"$description\", \"address\": {\"url\": \"$link\"} }" http://localhost:8088/links
}

# Check if no arguments are provided, then print help
if [ "$#" -ne 2 ]; then
    echo "Usage: add_link <description> <link>"
    exit 1
fi

add_link "$1" "$2"


#curl -X POST -H "Content-Type: application/json" \
#     -d '{"description": "New Link", "address": {"url": "https://example.com"}}' \
#     http://localhost:8088/links

