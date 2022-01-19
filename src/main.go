package main

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"
)

// var srcFile = "../asset/faceimgs/nara.jpg"
// var trgFile = "../asset/faceimgs/nara_new.jpg"

var srcFile = "../asset/faceimgs/n2.jpg"
var trgFile = "../asset/faceimgs/n2_new.jpg"

var classifierFile = "../asset/classifier/cascade_frontalface_default.xml"

func main() {
	img := gocv.IMRead(srcFile, gocv.IMReadColor)
	defer img.Close()

	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(classifierFile)
	defer classifier.Close()

	// detect faces
	rects := classifier.DetectMultiScale(img)
	fmt.Printf("found %d faces\n", len(rects))
	for _, r := range rects {
		fmt.Println("detected", r)
		gocv.Rectangle(&img, r, blue, 3)
	}

	// write image
	gocv.IMWrite(trgFile, img)
}
