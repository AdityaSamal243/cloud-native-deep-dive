//mandlebrot emits a png image to mandlebrot fractal 
// make use of goroutines and channels to make the compute faster


package main	

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main(){
	const(
		xmin,ymin,xmax,ymax = -2,-2,+2,+2
		width, height = 1024, 1024
	)
	img :=image.NewRGBA(image.Rect(0,0,width,height))
	
	for i:=0;i<5;i++{
        
	}
}