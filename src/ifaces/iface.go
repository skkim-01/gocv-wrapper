package ifaces

import (
	"fmt"
	internal "skkim-01/gocv-wrapper/src/internals"

	"gocv.io/x/gocv"
)

type IFace struct {
	bStart bool
}

func (o *IFace) Init() {
	o.bStart = false
}

func (o *IFace) GetAvaliCamIdxs() []int {
	var nIndex int = 0
	var arrDevices []int = make([]int, 0)

	for {
		_, err := gocv.VideoCaptureDevice(nIndex)
		if err != nil {
			break
		}
		arrDevices = append(arrDevices, nIndex)
		nIndex++
	}
	return arrDevices
}

func (o *IFace) StartCam(nCamIndex int) string {
	if o.bStart {
		return "Already started"
	}

	webcam, err := gocv.VideoCaptureDevice(nCamIndex)
	if err != nil {
		return err.Error()
	}

	window := gocv.NewWindow("gocv-wrapper")
	retv := window.GetWindowProperty(gocv.WindowPropertyVisible)
	fmt.Println(retv)

	o.bStart = true

	src := gocv.NewMat()
	dst := gocv.NewMat()

	for {
		if !o.bStart {
			break
		}
		webcam.Read(&src)
		internal.EdgeMask(&src, &dst, 7, 5, 3)
		window.IMShow(dst)

		keyPress := window.WaitKey(1)
		if keyPress == 27 {
			break
		}
	}

	o.bStart = false

	var e error
	e = dst.Close()
	fmt.Println(e)
	e = src.Close()
	fmt.Println(e)
	e = webcam.Close()
	fmt.Println(e)
	e = window.Close()
	fmt.Println(e)
	retv = window.GetWindowProperty(gocv.WindowPropertyVisible)
	fmt.Println(retv)

	return "fin"
}

func (o *IFace) Stop() {
	o.bStart = false
}
