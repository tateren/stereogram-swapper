package main

import (
	"flag"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// jpgっぽいファイル名
var r = regexp.MustCompile("\\.(jpe?g|JPE?G)$")

func swap(img image.Image) *image.RGBA {
	size := img.Bounds()
	rgba := image.NewRGBA(size)

	// 左半分に元画像の右半分を描画
	left := image.Rectangle{image.ZP, image.Point{size.Dx() / 2, size.Dy()}}
	draw.Draw(rgba, left, img, image.Point{size.Dx() / 2, 0}, draw.Src)
	// 右半分に元画像の左半分を描画
	right := image.Rectangle{image.Point{size.Dx() / 2, 0}, size.Max}
	draw.Draw(rgba, right, img, image.ZP, draw.Src)
	return rgba
}

func main() {
	var inputDir, outputDir string
	flag.StringVar(&inputDir, "i", ".", "input directory")
	flag.StringVar(&outputDir, "o", ".", "output directory")
	flag.Parse()

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		log.Fatal("Failed to read directory: ", err)
	}

	for i := range files {
		if r.MatchString(filepath.Ext(files[i].Name())) {
			file, err := os.Open(filepath.Join(inputDir, files[i].Name()))
			if err != nil {
				log.Fatal("Failed to open input file: ", err)
			}
			defer file.Close()
			img, _, err := image.Decode(file)
			if err != nil {
				log.Fatal("Failed to decode image: ", err)
			}

			rgba := swap(img)

			output, err := os.Create(filepath.Join(outputDir, files[i].Name()))
			if err != nil {
				log.Fatal("Failed to create output file: ", err)
			}

			var opt jpeg.Options
			opt.Quality = 100
			jpeg.Encode(output, rgba, &opt)
		}
	}
}
