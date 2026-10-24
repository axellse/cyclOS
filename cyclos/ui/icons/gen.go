//go:build ignore

package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	inputFile := os.Args[2]
	inputDirectory := os.Args[3]

	entries, err := os.ReadDir(inputDirectory)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(inputFile)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString("//go:generate go run gen.go -- ./icons.go ./icons/\n")
	file.WriteString("\n")
	file.WriteString("package icons\n")
	file.WriteString("\n")
	file.WriteString("type Icon struct {\n")
	file.WriteString("	Blob string\n")
	file.WriteString("	Width int16\n")
	file.WriteString("	Height int16\n")
	file.WriteString("}\n")
	file.WriteString("\n")
	file.WriteString("// monochrome bitmaps\n")
	file.WriteString("var Icons = map[string]Icon{\n")

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		f, err := os.Open(filepath.Join(inputDirectory, entry.Name()))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		img, err := png.Decode(f)
		if err != nil {
			panic(err)
		}

		width, height := img.Bounds().Max.X, img.Bounds().Max.Y

		ba := []byte{}
		wb := byte(0)
		for y := range height {
			for x := range width {
				_, _, _, alpha := img.At(x,y).RGBA()
				shiftby := ((y * width) + x) % 8
				if alpha != 0 {
					mask := byte(0b10000000) >> (((y * width) + x) % 8)
					wb = wb | mask
				}

				if shiftby == 7 {
					ba = append(ba, wb)
					wb = byte(0)
				}
			}
		}

		if (height * width) % 8 != 0 {
			fmt.Println("appending extra byte for", entry.Name())
			ba = append(ba, wb)
		}

		nameWithoutExtension := strings.Split(entry.Name(), ".")[0]
		file.WriteString(`	"`+ nameWithoutExtension + `": ` + "{\n")
		file.WriteString("		Width: " + strconv.Itoa(width) + ",\n")
		file.WriteString("		Height: " + strconv.Itoa(height) + ",\n")
		file.WriteString("		Blob: " + strconv.QuoteToASCII(string(ba)) + ",\n")
		file.WriteString("	},\n")
	}
	file.WriteString("}")
	fmt.Println("transcoded", len(entries), "file(s)")
}
