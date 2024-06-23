package main

import (
	"bufio"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: ./jackal9000 IMAGE_PATH [OPTIONAL QUALITY 0-10]")
		waitInput()
		return
	}
	imgPath := os.Args[1]

	filed, err := os.Open(imgPath)
	chk(err)
	defer filed.Close()

	ext := strings.ToUpper(filepath.Ext(filed.Name()))
	ext = strings.TrimPrefix(ext, ".")
	var imgd image.Image
	switch ext {
	case "JPEG", "JPG":
		imgd, err = jpeg.Decode(filed)
	case "PNG":
		imgd, err = png.Decode(filed)
	default:
		println("Unknown format:", ext)
		println("Supports: jpeg, png.")
		return
	}

	chk(err)

	outFile, err := os.Create(getFileName(filed.Name()) + "_jack.jpeg")
	chk(err)

	quality := 8
	if len(os.Args) >= 3 {
		qualStr := os.Args[2]
		qI, err := strconv.Atoi(qualStr)
		if err == nil {
			quality = qI
		}
	}
	err = jpeg.Encode(outFile, imgd, &jpeg.Options{
		Quality: quality,
	})
	chk(err)
}

func getFileName(fileName string) string {
	return strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
}

func chk(err error) {
	if err != nil {
		println("ERROR:", err.Error())
		os.Exit(1)
	}
}

func waitInput() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
