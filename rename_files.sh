#!/bin/bash

FOLDER="cropped_characters"

if [ ! -d "$FOLDER" ]; then
    echo "Folder '$FOLDER' not found!"
    exit 1
fi

for file in "$FOLDER"/*; do
    [ -f "$file" ] || continue

    filename=$(basename "$file")

    if command -v feh &> /dev/null; then
        feh "$file" &
        FEH_PID=$!
    elif command -v display &> /dev/null; then
        display "$file" &
        DISPLAY_PID=$!
    fi

    echo "Current file: $filename"
    read -p "Enter the character for this image: " newname

    if [ -n "$FEH_PID" ]; then
        kill "$FEH_PID" 2>/dev/null
    fi
    if [ -n "$DISPLAY_PID" ]; then
        kill "$DISPLAY_PID" 2>/dev/null
    fi

    ascii_value=$(printf "%d" "'$newname")

    if [ -n "$newname" ]; then
        mv "$file" "$FOLDER/${ascii_value}.png"
        echo "Renamed to: ${ascii_value}.png"
    else
        echo "Skipped."
    fi
done

echo "Renaming process completed!"
