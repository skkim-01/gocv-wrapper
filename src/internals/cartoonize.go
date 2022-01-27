package internal

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func EdgeMask(src *gocv.Mat, dst *gocv.Mat, nBlur int, nLine int, f32Blur float32) {
	gocv.CvtColor(*src, dst, gocv.ColorBGRToGray)
	gocv.MedianBlur(*dst, dst, nBlur)
	// for edge test
	// gocv.AdaptiveThreshold(*dst, dst, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, nLine, f32Blur)
	// return

	edges := gocv.NewMat()
	defer edges.Close()

	gocv.AdaptiveThreshold(*dst, &edges, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, nLine, f32Blur)

	edgePress := gocv.NewMatWithSize(1024, 780, gocv.MatTypeCV8U)
	defer edgePress.Close()
	gocv.BilateralFilter(*src, &edgePress, 3, 7, 7)

	gocv.BitwiseAndWithMask(edgePress, edgePress, dst, edges)
}

func Edge2(src *gocv.Mat, dst *gocv.Mat, nBlur int, nLine int, f32Blur float32) {
	gocv.CvtColor(*src, dst, gocv.ColorBGRToGray)
	gocv.GaussianBlur(*dst, dst, image.Point{X: 3, Y: 3}, 0, 0, gocv.BorderDefault)

	edges := gocv.NewMatWithSize(1024, 780, gocv.MatTypeCV8U)
	defer edges.Close()

	gocv.Laplacian(*dst, &edges, -1, 5, 1, 0, gocv.BorderDefault)
	gocv.ConvertScaleAbs(edges, &edges, 1, 0)
	gocv.Threshold(*dst, dst, 150, 255, gocv.ThresholdBinary)

	edgePress := gocv.NewMatWithSize(1024, 780, gocv.MatTypeCV8U)
	defer edgePress.Close()
	gocv.BilateralFilter(*src, &edgePress, 2, 50, 0.4)

	//edgePress.ConvertTo(&edgePress, gocv.MatTypeCV8U)
	//edges.ConvertTo(&edges, gocv.MatTypeCV8U)

	fmt.Println(edgePress.Type())
	fmt.Println(edges.Type())
	fmt.Println(dst.Type())

	fmt.Println(edgePress.Size())
	fmt.Println(edges.Size())
	fmt.Println(dst.Size())

	gocv.BitwiseAndWithMask(edgePress, edgePress, dst, edges)

	//gocv.BitwiseAnd(edgePress, edgePress, dst)
}

func ColorQuant(src *gocv.Mat, dst *gocv.Mat, nColor int) {

}

func ColorQuantization(src *gocv.Mat, dst *gocv.Mat, nColor int) {
	criteria := gocv.NewTermCriteria(gocv.EPS+gocv.MaxIter, 10, 1.03)
	l := gocv.NewMat()
	cn := gocv.NewMat()
	p := gocv.Zeros(src.Cols()*src.Rows(), 5, gocv.MatTypeCV32F)
	bgr := gocv.Split(*src)

	for _, v := range bgr {
		v.ConvertTo(&v, gocv.MatTypeCV32F)
	}

	for i := 0; i < src.Cols()*src.Rows(); i++ {
		pbgr0, _ := bgr[0].DataPtrFloat32()
		pbgr1, _ := bgr[1].DataPtrFloat32()
		pbgr2, _ := bgr[2].DataPtrFloat32()

		p.SetFloatAt(i, 0, float32((i/src.Cols())/src.Rows()))
		p.SetFloatAt(i, 1, float32((i/src.Cols())/src.Cols()))
		p.SetFloatAt(i, 2, float32(pbgr0[i]/255.0))
		p.SetFloatAt(i, 3, float32(pbgr1[i]/255.0))
		p.SetFloatAt(i, 4, float32(pbgr2[i]/255.0))
	}

	gocv.KMeans(p, nColor, &l, criteria, 10, gocv.KMeansRandomCenters, &cn)

	colors := make([]int, nColor)
	for i := 0; i < nColor; i++ {
		colors[i] = 255 / (i + 1)
	}

	//clustered = Mat(src.rows, src.cols, CV_32F);
	cl := gocv.NewMatWithSize(src.Rows(), src.Cols(), gocv.MatTypeCV32F)
	for i := 0; i < src.Cols()*src.Rows(); i++ {
		cl.SetFloatAt(i/src.Cols(), i%src.Cols(), float32(colors[l.GetIntAt(0, i)]))
	}

	cl.ConvertTo(dst, gocv.MatTypeCV8U)

	// filt := src.Reshape(1, src.Rows())
	// filt.ConvertTo(&filt, gocv.MatTypeCV32F)
	// gocv.KMeans(filt.T(), nColor, &l, criteria, 10, gocv.KMeansRandomCenters, &cn)
	//*dst = filt.Reshape(2, filt.Rows())
}
