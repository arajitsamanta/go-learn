package interfaces

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

/*
GoImages - GO image package example.

Package image(https://golang.org/pkg/image/#Image) defines the Image interface:

package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.

(See the documentation https://golang.org/pkg/image/#Image for all the details.)

The color.Color and color.Model types are also interfaces, but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package
*/
func GoImages() {
	fmt.Printf("\n****Running interfaces.GoImages(), Go image example. \n")
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

//Image type
type Image struct {
	Width, Height int
	colr          uint8
}

//Bounds - Implement Bounds() on Image type
func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

//ColorModel - Implement ColorModel() on Image type
func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

//At - Implement At() on Image type
func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.colr + uint8(x), r.colr + uint8(y), 255, 255}
}

/*ImageExcercise - Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one
*/
func ImageExcercise() {
	fmt.Printf("****Running interfaces.ImageExcercise(), Go image excercise. \n")
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
}
