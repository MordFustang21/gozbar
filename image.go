// Package gozbar image bindings for golang.
// Read the ZBar documents for details
package gozbar

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

import (
	"image"
	"runtime"
	"unsafe"
)

// Image contains a zbar image and the grayscale values.
type Image struct {
	image *C.zbar_image_t
	gray  *image.Gray
}

// FromImage will create an ZBar image object from an image.Image.
// To scan the image, call a Scanner.
func FromImage(img image.Image) *Image {
	// allocate the image wrapper
	ret := &Image{
		image: C.zbar_image_create(),
	}

	// get the height and width of the given image
	bounds := img.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y

	// Create a grayscale image
	ret.gray = image.NewGray(bounds)

	// populate all the pixels in the gray image faster than draw.Draw()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			ret.gray.Set(x, y, img.At(x, y))
		}
	}

	C.zbar_image_set_format(ret.image, C.ulong(0x30303859)) // Y800 (grayscale)
	C.zbar_image_set_size(ret.image, C.uint(w), C.uint(h))
	C.zbar_image_set_data(ret.image, unsafe.Pointer(&ret.gray.Pix[0]), C.ulong(len(ret.gray.Pix)), nil)

	// finalizer
	runtime.SetFinalizer(ret, (*Image).Destroy)

	return ret
}

// First will return the first scanned symbol of this image.
// To iterate over the symbols, use Symbol.Each() function
func (i *Image) First() *Symbol {
	s := C.zbar_image_first_symbol(i.image)

	if s == nil {
		return nil
	}

	return &Symbol{
		symbol: s,
	}
}

// Destroy this object
func (i *Image) Destroy() {
	C.zbar_image_destroy(i.image)
}
