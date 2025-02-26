find archive/train -type f -name "/*/*.jpg" | while read file; do
    mv "$file" "${file%.jpg}.png"
done
