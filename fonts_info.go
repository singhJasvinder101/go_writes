package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Fonts() {
	dirs, err := os.ReadDir("archive/train copy")
	if err != nil {
		fmt.Println(err)
		return
	}

	minSize := int64(99999999999)
	minDir := ""

	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Println("Checking folder:", dir.Name())

			charFolder := filepath.Join("archive/train copy", dir.Name())

			files, err := os.ReadDir(charFolder)
			if err != nil {
				fmt.Println("Error reading folder:", err)
				continue
			}

			fmt.Println("size of ",  dir.Name(), " is ", len(files))
			if int64(len(files)) < minSize {
				minSize = int64(len(files))
				minDir = charFolder
			}
		}
	}

	if minDir != "" {
		fmt.Printf("minimum size %d (Directory: %s)\n", minSize, minDir)
	} else {
		fmt.Println("no directories found")
	}
}

