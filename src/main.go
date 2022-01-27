package main

import (
	"fmt"
	"image"
	"image/color"
	internal "skkim-01/gocv-wrapper/src/internals"

	"gocv.io/x/gocv"
)

// var srcFile = "../asset/faceimgs/nara.jpg"
// var trgFile = "../asset/faceimgs/nara_new.jpg"

var srcFile = "../asset/faceimgs/n2.jpg"
var trgFile = "../asset/faceimgs/n2_new.jpg"

var classifierFile = "../asset/classifier/cascade_frontalface_default.xml"

func main() {
	fmt.Println("gocv-wrapper start")

	webcam, _ := gocv.VideoCaptureDevice(0)
	defer webcam.Close()

	window := gocv.NewWindow("Cartoonize")
	defer window.Close()

	//img := gocv.NewMat()
	//img := gocv.NewMatWithSize(1024, 780, gocv.MatTypeCV32F)
	img := gocv.NewMatWithSize(1024, 780, gocv.MatTypeCV8U)
	//img.ConvertTo(&img, gocv.MatTypeCV32F)
	defer img.Close()

	edges := gocv.NewMatWithSize(1024, 780, gocv.MatTypeCV8U)
	defer img.Close()

	// cl := gocv.NewMatWithSize(1024, 780, gocv.MatTypeCV8U)
	// defer img.Close()

	for {
		webcam.Read(&img)

		/// sample 4
		internal.EdgeMask(&img, &edges, 7, 13, 11)
		//internal.EdgeMask(&img, &edges, 7, 5, 3)
		//internal.EdgeMask(&img, &edges, 7, 11, 5.1)

		//internal.Edge2(&img, &edges, 1, 2, 3)

		//internal.ColorQuantization(&edges, &cl, 8)
		// display
		window.IMShow(edges)

		/// sample 1
		// gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
		// gocv.MedianBlur(img, &img, 7)
		// gocv.Laplacian(img, &img, gocv.MatTypeCV8U, 5, 1, 0, gocv.BorderReplicate)
		// gocv.Threshold(img, &img, 80, 255, gocv.ThresholdBinary)

		// //colorpainting
		// gocv.Resize(img, &img, image.Point{X: img.Rows() / 2, Y: img.Cols() / 2}, float64(img.Rows()/2), float64(img.Cols()/2), gocv.InterpolationLinear)
		// for i := 0; i < 7; i++ {
		// 	gocv.BilateralFilter(img, &tmp, 9, 9, 7)
		// 	gocv.BilateralFilter(tmp, &img, 9, 9, 7)
		// }
		// gocv.Resize(img, &img, image.Point{X: img.Rows() * 2, Y: img.Cols() * 2}, float64(img.Rows()*2), float64(img.Cols()*2), gocv.InterpolationLinear)

		/// sample 2
		// gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
		// gocv.MedianBlur(img, &edges, 7)
		// gocv.AdaptiveThreshold(edges, &edges, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, 13, 7)

		// filt := img.Reshape(1, img.Rows())
		// filt.ConvertTo(&filt, gocv.MatTypeCV32F)

		// crit := gocv.NewTermCriteria(gocv.EPS+gocv.MaxIter, 20, 0.001)
		// total_color := 8

		// //gocv.ConvertScaleAbs(img, &filterImg, 1, 0)
		// gocv.KMeans(filt.T(), total_color, &bestLabels, crit, 10, gocv.KMeansRandomCenters, &center)
		// //result := gocv.NewMatWithSize(1024, 780, gocv.MatTypeCV8U)
		// //gocv.TextureFlattening(center, bestLabels, &result, 30, 45, 3)
		// img := filt.Reshape(2, filt.Rows())

		// sample 3
		// const MedianBlurFilterSz int = 7
		// const BlockSize int = 13
		// const ThresholdConst float32 = 7

		// gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
		// gocv.MedianBlur(img, &img, MedianBlurFilterSz)
		// //gocv.Laplacian(img, &img, gocv.MatTypeCV8UC1, 5, 1, 0, gocv.BorderReplicate)
		// gocv.AdaptiveThreshold(img, &edges, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, BlockSize, ThresholdConst)

		// gocv.BilateralFilter(img, &center, 7, 200, 200)
		// gocv.BitwiseAnd(img, center, &img)

		// gocv.Threshold(img, &filterImg, 80, 250, gocv.ThresholdBinary)

		//gocv.BilateralFilter(img, &center, 9, 250, 250)

		// display
		// window.IMShow(img)
		keyPress := window.WaitKey(1)

		if keyPress == 27 {
			break
		}
	}
}

func _detectFace() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	defer webcam.Close()

	// open display window

	//webcam.Set(gocv.VideoCaptureProperties())
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(classifierFile) {
		fmt.Printf("Error reading cascade file: %v\n", classifierFile)
		return
	}

	fmt.Printf("start reading camera device: %v\n", 0)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", 0)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image,
		// along with text identifying as "Human"
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)

			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}

func _helloVideo() {
	fmt.Println("gocv-wrapper start")

	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("test window")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}

func _faceReader() {
	fmt.Println("gocv-wrapper start")

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
