package core

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"

	"github.com/devectron/sunlight/log"
)

//PngToJpeg convert png image to jpg format.
func PngToJpeg(img string) {
	log.Inf("Converting from PNG to JPEG ...")
	name := strings.Split(img, path.Ext(img))
	imgF, err := os.Open(img)
	if err != nil {
		log.Err("Cannot read file.... %v", err)
	}
	defer imgF.Close()
	imgsrc, err := png.Decode(imgF)
	if err != nil {
		log.Err("Cannot decode to png file %v", err)
	}
	imgdst := image.NewRGBA(imgsrc.Bounds())
	draw.Draw(imgdst, imgdst.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(imgdst, imgdst.Bounds(), imgsrc, imgsrc.Bounds().Min, draw.Over)
	check := saveImage(imgdst, name[0]+".jpg", path.Ext(img))
	if !check {
		log.Err("Error While converting ...")
	}
}
func saveImage(imgsrc image.Image, imgname string, format string) bool {
	out, err := os.Create(imgname)
	if err != nil {
		log.Err("Can't save image %v", err)
	}
	defer out.Close()
	log.Inf("Saving %s", imgsrc)
	switch format {
	case "png":
		png.Encode(out, imgsrc)
		return true
	case "jpg", "jpeg":
		jpeg.Encode(out, imgsrc, nil)
		return true
	}
	return false
}
