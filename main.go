package main

import (
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"
	"golang.org/x/image/draw"
)

const (
	fontsDir = "archive/train copy"
	defaultFile = "dummy.txt" 
	bgImagePath = "myfont/bg.png"
	outputImage = "output.png"
	lineHeight = 140 
	charSpacing = 0  
	spaceWidth = 70

	
	charWidth = 60  
	charHeight = 150 
)

var index int = 30
var outputFileName int
var similarFont string

func main() {
	rand.Seed(time.Now().UnixNano())

	fileName := getFileName()

	fmt.Println("Do you want similar font for each character (type y) or similar (type n)") 
	fmt.Scanln(&similarFont)

	text, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file, using default text")
		text, _ = os.ReadFile(defaultFile)
	}

	// sheetWidth := bgImg.Bounds().Dx()
	gap, ht := 0, 0 

	renderTextToImage(string(text), 0, gap, ht)
}

func renderTextToImage(text string, startIndex, gap, ht int) {
	bgFile, err := os.Open(bgImagePath)
	if err != nil {
		fmt.Println("Error opening background image ", err)
		return
	}
	defer bgFile.Close()

	bgImg, _, err := image.Decode(bgFile)
	if err != nil {
		fmt.Println("Error decoding background image ", err)
		return
	}

	sheetWidth := bgImg.Bounds().Dx()
	dc := gg.NewContextForImage(bgImg)
	
	for i := startIndex; i < len(text); i++ {
		char := text[i]
		if char == '\n' {
			gap = 0
			ht += lineHeight
			continue
		}
		if char == ' ' {
			gap += spaceWidth
			continue
		}
		charFolder := fmt.Sprintf("%s/%d", fontsDir, char)

		fmt.Println("character from the text ", string(char))
		charImagePath, err := getRandomImage(charFolder)
		if err != nil {
			fmt.Printf("No images found for character %c\n", char)
			continue
		}

		charImgFile, err := os.Open(charImagePath)
		if err != nil {
			fmt.Printf("Error opening image for character %c\n", char)
			continue
		}
		defer charImgFile.Close()

		charImg, err := png.Decode(charImgFile)
		if err != nil {
			fmt.Println("Error decoding character image ", err)
			continue
		}

		resizedCharImg := image.NewRGBA(image.Rect(0, 0, charWidth, charHeight))
		draw.BiLinear.Scale(resizedCharImg, resizedCharImg.Bounds(), charImg, charImg.Bounds(), draw.Over, nil)

		dc.DrawImage(resizedCharImg, gap, ht)
		gap += charWidth + charSpacing - 5

		if gap+charWidth > sheetWidth {
			gap = 0
			ht += lineHeight
		}
		if ht+charHeight > bgImg.Bounds().Dy() {
			fmt.Println("Reached end of image, saving and continuing")
			saveImage(dc)
			renderTextToImage(text, i+1, 0, 0)
			return
		}
	}
	saveImage(dc)
}

func saveImage(dc *gg.Context) {
	outputFileName++
	outputFile := fmt.Sprintf("%d.png", outputFileName)
	err := dc.SavePNG(outputFile)
	if err != nil {
		fmt.Println("Error saving output image", err)
		return
	}
	fmt.Println("Saved", outputFile)
}

func getFileName() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	fmt.Println("No file entered. Using default file...")
	return defaultFile
}

func getRandomImage(dir string) (string, error) {
	files, err := os.ReadDir(dir)

	// fmt.Println("files ",files)
	if err != nil || len(files) == 0 {
		return "", fmt.Errorf("no images found")
	}

	if !(similarFont == "y") {
		index = rand.Intn(len(files))
	}
	
	if index >= len(files) {
		fmt.Println("Index out of bounds. Selecting a random font instead.")
		index = rand.Intn(len(files))
	}

	randomFile := files[index].Name()
	return filepath.Join(dir, randomFile), nil
}
