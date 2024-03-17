package colors

// Taken from https://github.com/esimov/gobrot/blob/master/internal/palette/color.go
// under MIT license
import "image/color"

type Color struct {
	Step  float64
	Color color.Color
}

type ColorMap struct {
	Keyword string
	Colors  []Color
}

var ColorPalettes = []ColorMap{
	{"AfternoonBlue", []Color{
		{0.0, color.RGBA{0x93, 0xd2, 0xca, 0xff}},
		{0.2, color.RGBA{0x6c, 0x98, 0xb8, 0xff}},
		{0.5, color.RGBA{0x38, 0x68, 0x85, 0xff}},
		{0.8, color.RGBA{0x17, 0x4f, 0x72, 0xff}},
		{1.0, color.RGBA{0x08, 0x2a, 0x4f, 0xff}},
	}},
	{"SummerBeach", []Color{
		{0.0, color.RGBA{0xff, 0xf0, 0x94, 0xff}},
		{0.3, color.RGBA{0xff, 0xb7, 0x2d, 0xff}},
		{0.6, color.RGBA{0xff, 0x8d, 0x00, 0xff}},
		{0.8, color.RGBA{0x2d, 0x69, 0xae, 0xff}},
		{1.0, color.RGBA{0x1e, 0x2c, 0x60, 0xff}},
	}},
	{"Biochemist", []Color{
		{Color: color.RGBA{0x51, 0x99, 0x25, 0xff}},
		{Color: color.RGBA{0x50, 0xa6, 0x1c, 0xff}},
		{Color: color.RGBA{0x4b, 0xb8, 0x0b, 0xff}},
		{Color: color.RGBA{0x4d, 0xcd, 0x00, 0xff}},
		{Color: color.RGBA{0x53, 0xdf, 0x00, 0xff}},
	}},
	{"Fiesta", []Color{
		{Color: color.RGBA{0x03, 0x67, 0xa6, 0xff}},
		{Color: color.RGBA{0x04, 0xad, 0xbf, 0xff}},
		{Color: color.RGBA{0x93, 0xa6, 0x03, 0xff}},
		{Color: color.RGBA{0xe5, 0xce, 0x1b, 0xff}},
		{Color: color.RGBA{0xc8, 0x3f, 0x2a, 0xff}},
	}},
	{"Hippi", []Color{
		{Color: color.RGBA{0x00, 0x04, 0x0f, 0xff}},
		{Color: color.RGBA{0x03, 0x26, 0x28, 0xff}},
		{Color: color.RGBA{0x07, 0x3e, 0x1e, 0xff}},
		{Color: color.RGBA{0x18, 0x55, 0x08, 0xff}},
		{Color: color.RGBA{0x5f, 0x6e, 0x0f, 0xff}},
		{Color: color.RGBA{0x84, 0x50, 0x19, 0xff}},
		{Color: color.RGBA{0x9b, 0x30, 0x22, 0xff}},
		{Color: color.RGBA{0xb4, 0x92, 0x2f, 0xff}},
		{Color: color.RGBA{0x94, 0xca, 0x3d, 0xff}},
		{Color: color.RGBA{0x4f, 0xd5, 0x51, 0xff}},
		{Color: color.RGBA{0x66, 0xff, 0xb3, 0xff}},
		{Color: color.RGBA{0x82, 0xc9, 0xe5, 0xff}},
		{Color: color.RGBA{0x9d, 0xa3, 0xeb, 0xff}},
		{Color: color.RGBA{0xd7, 0xb5, 0xf3, 0xff}},
		{Color: color.RGBA{0xfd, 0xd6, 0xf6, 0xff}},
		{Color: color.RGBA{0xff, 0xf0, 0xf2, 0xff}},
	}},
	{"Vivid", []Color{
		{Color: color.RGBA{0x02, 0x3b, 0x2b, 0xff}},
		{Color: color.RGBA{0x36, 0x34, 0x48, 0xff}},
		{Color: color.RGBA{0x04, 0x8b, 0x64, 0xff}},
		{Color: color.RGBA{0x81, 0x6a, 0x6e, 0xff}},
		{Color: color.RGBA{0x8d, 0x5e, 0x67, 0xff}},
		{Color: color.RGBA{0x98, 0x52, 0x60, 0xff}},
		{Color: color.RGBA{0xa4, 0x46, 0x59, 0xff}},
		{Color: color.RGBA{0xd3, 0x17, 0x3d, 0xff}},
		{Color: color.RGBA{0xe1, 0x08, 0x34, 0xff}},
		{Color: color.RGBA{0xde, 0x17, 0x30, 0xff}},
		{Color: color.RGBA{0xdc, 0x26, 0x2d, 0xff}},
		{Color: color.RGBA{0xd9, 0x36, 0x2a, 0xff}},
		{Color: color.RGBA{0xd7, 0x45, 0x27, 0xff}},
		{Color: color.RGBA{0xd5, 0x55, 0x24, 0xff}},
		{Color: color.RGBA{0xd2, 0x64, 0x21, 0xff}},
		{Color: color.RGBA{0xd0, 0x74, 0x1e, 0xff}},
		{Color: color.RGBA{0xce, 0x83, 0x1b, 0xff}},
		{Color: color.RGBA{0xcb, 0x92, 0x18, 0xff}},
		{Color: color.RGBA{0xc9, 0xa2, 0x15, 0xff}},
		{Color: color.RGBA{0x11, 0x03, 0x12, 0xff}},
		{Color: color.RGBA{0x33, 0x08, 0x35, 0xff}},
		{Color: color.RGBA{0xf5, 0xca, 0xf7, 0xff}},
	}},
	{"Plan9", []Color{
		{Color: color.RGBA{0x00, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0x44, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0x88, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0xcc, 0xff}},
		{Color: color.RGBA{0x00, 0x44, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0x44, 0x44, 0xff}},
		{Color: color.RGBA{0x00, 0x44, 0x88, 0xff}},
		{Color: color.RGBA{0x00, 0x44, 0xcc, 0xff}},
		{Color: color.RGBA{0x00, 0x88, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0x88, 0x44, 0xff}},
		{Color: color.RGBA{0x00, 0x88, 0x88, 0xff}},
		{Color: color.RGBA{0x00, 0x88, 0xcc, 0xff}},
		{Color: color.RGBA{0x00, 0xcc, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0xcc, 0x44, 0xff}},
		{Color: color.RGBA{0x00, 0xcc, 0x88, 0xff}},
		{Color: color.RGBA{0x00, 0xcc, 0xcc, 0xff}},
		{Color: color.RGBA{0x00, 0xdd, 0xdd, 0xff}},
		{Color: color.RGBA{0x11, 0x11, 0x11, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0x55, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0x99, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0xdd, 0xff}},
		{Color: color.RGBA{0x00, 0x55, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0x55, 0x55, 0xff}},
		{Color: color.RGBA{0x00, 0x4c, 0x99, 0xff}},
		{Color: color.RGBA{0x00, 0x49, 0xdd, 0xff}},
		{Color: color.RGBA{0x00, 0x99, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0x99, 0x4c, 0xff}},
		{Color: color.RGBA{0x00, 0x99, 0x99, 0xff}},
		{Color: color.RGBA{0x00, 0x93, 0xdd, 0xff}},
		{Color: color.RGBA{0x00, 0xdd, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0xdd, 0x49, 0xff}},
		{Color: color.RGBA{0x00, 0xdd, 0x93, 0xff}},
		{Color: color.RGBA{0x00, 0xee, 0x9e, 0xff}},
		{Color: color.RGBA{0x00, 0xee, 0xee, 0xff}},
		{Color: color.RGBA{0x22, 0x22, 0x22, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0x66, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0xaa, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0xee, 0xff}},
		{Color: color.RGBA{0x00, 0x66, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0x66, 0x66, 0xff}},
		{Color: color.RGBA{0x00, 0x55, 0xaa, 0xff}},
		{Color: color.RGBA{0x00, 0x4f, 0xee, 0xff}},
		{Color: color.RGBA{0x00, 0xaa, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0xaa, 0x55, 0xff}},
		{Color: color.RGBA{0x00, 0xaa, 0xaa, 0xff}},
		{Color: color.RGBA{0x00, 0x9e, 0xee, 0xff}},
		{Color: color.RGBA{0x00, 0xee, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0xee, 0x4f, 0xff}},
		{Color: color.RGBA{0x00, 0xff, 0x55, 0xff}},
		{Color: color.RGBA{0x00, 0xff, 0xaa, 0xff}},
		{Color: color.RGBA{0x00, 0xff, 0xff, 0xff}},
		{Color: color.RGBA{0x33, 0x33, 0x33, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0x77, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0xbb, 0xff}},
		{Color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
		{Color: color.RGBA{0x00, 0x77, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0x77, 0x77, 0xff}},
		{Color: color.RGBA{0x00, 0x5d, 0xbb, 0xff}},
		{Color: color.RGBA{0x00, 0x55, 0xff, 0xff}},
		{Color: color.RGBA{0x00, 0xbb, 0x00, 0xff}},
		{Color: color.RGBA{0x00, 0xbb, 0x5d, 0xff}},
		{Color: color.RGBA{0x00, 0xbb, 0xbb, 0xff}},
		{Color: color.RGBA{0x00, 0xaa, 0xff, 0xff}},
		{Color: color.RGBA{0x00, 0xff, 0x00, 0xff}},
		{Color: color.RGBA{0x44, 0x00, 0x44, 0xff}},
		{Color: color.RGBA{0x44, 0x00, 0x88, 0xff}},
		{Color: color.RGBA{0x44, 0x00, 0xcc, 0xff}},
		{Color: color.RGBA{0x44, 0x44, 0x00, 0xff}},
		{Color: color.RGBA{0x44, 0x44, 0x44, 0xff}},
		{Color: color.RGBA{0x44, 0x44, 0x88, 0xff}},
		{Color: color.RGBA{0x44, 0x44, 0xcc, 0xff}},
		{Color: color.RGBA{0x44, 0x88, 0x00, 0xff}},
		{Color: color.RGBA{0x44, 0x88, 0x44, 0xff}},
		{Color: color.RGBA{0x44, 0x88, 0x88, 0xff}},
		{Color: color.RGBA{0x44, 0x88, 0xcc, 0xff}},
		{Color: color.RGBA{0x44, 0xcc, 0x00, 0xff}},
		{Color: color.RGBA{0x44, 0xcc, 0x44, 0xff}},
		{Color: color.RGBA{0x44, 0xcc, 0x88, 0xff}},
		{Color: color.RGBA{0x44, 0xcc, 0xcc, 0xff}},
		{Color: color.RGBA{0x44, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0x55, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0x55, 0x00, 0x55, 0xff}},
		{Color: color.RGBA{0x4c, 0x00, 0x99, 0xff}},
		{Color: color.RGBA{0x49, 0x00, 0xdd, 0xff}},
		{Color: color.RGBA{0x55, 0x55, 0x00, 0xff}},
		{Color: color.RGBA{0x55, 0x55, 0x55, 0xff}},
		{Color: color.RGBA{0x4c, 0x4c, 0x99, 0xff}},
		{Color: color.RGBA{0x49, 0x49, 0xdd, 0xff}},
		{Color: color.RGBA{0x4c, 0x99, 0x00, 0xff}},
		{Color: color.RGBA{0x4c, 0x99, 0x4c, 0xff}},
		{Color: color.RGBA{0x4c, 0x99, 0x99, 0xff}},
		{Color: color.RGBA{0x49, 0x93, 0xdd, 0xff}},
		{Color: color.RGBA{0x49, 0xdd, 0x00, 0xff}},
		{Color: color.RGBA{0x49, 0xdd, 0x49, 0xff}},
		{Color: color.RGBA{0x49, 0xdd, 0x93, 0xff}},
		{Color: color.RGBA{0x49, 0xdd, 0xdd, 0xff}},
		{Color: color.RGBA{0x4f, 0xee, 0xee, 0xff}},
		{Color: color.RGBA{0x66, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0x66, 0x00, 0x66, 0xff}},
		{Color: color.RGBA{0x55, 0x00, 0xaa, 0xff}},
		{Color: color.RGBA{0x4f, 0x00, 0xee, 0xff}},
		{Color: color.RGBA{0x66, 0x66, 0x00, 0xff}},
		{Color: color.RGBA{0x66, 0x66, 0x66, 0xff}},
		{Color: color.RGBA{0x55, 0x55, 0xaa, 0xff}},
		{Color: color.RGBA{0x4f, 0x4f, 0xee, 0xff}},
		{Color: color.RGBA{0x55, 0xaa, 0x00, 0xff}},
		{Color: color.RGBA{0x55, 0xaa, 0x55, 0xff}},
		{Color: color.RGBA{0x55, 0xaa, 0xaa, 0xff}},
		{Color: color.RGBA{0x4f, 0x9e, 0xee, 0xff}},
		{Color: color.RGBA{0x4f, 0xee, 0x00, 0xff}},
		{Color: color.RGBA{0x4f, 0xee, 0x4f, 0xff}},
		{Color: color.RGBA{0x4f, 0xee, 0x9e, 0xff}},
		{Color: color.RGBA{0x55, 0xff, 0xaa, 0xff}},
		{Color: color.RGBA{0x55, 0xff, 0xff, 0xff}},
		{Color: color.RGBA{0x77, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0x77, 0x00, 0x77, 0xff}},
		{Color: color.RGBA{0x5d, 0x00, 0xbb, 0xff}},
		{Color: color.RGBA{0x55, 0x00, 0xff, 0xff}},
		{Color: color.RGBA{0x77, 0x77, 0x00, 0xff}},
		{Color: color.RGBA{0x77, 0x77, 0x77, 0xff}},
		{Color: color.RGBA{0x5d, 0x5d, 0xbb, 0xff}},
		{Color: color.RGBA{0x55, 0x55, 0xff, 0xff}},
		{Color: color.RGBA{0x5d, 0xbb, 0x00, 0xff}},
		{Color: color.RGBA{0x5d, 0xbb, 0x5d, 0xff}},
		{Color: color.RGBA{0x5d, 0xbb, 0xbb, 0xff}},
		{Color: color.RGBA{0x55, 0xaa, 0xff, 0xff}},
		{Color: color.RGBA{0x55, 0xff, 0x00, 0xff}},
		{Color: color.RGBA{0x55, 0xff, 0x55, 0xff}},
		{Color: color.RGBA{0x88, 0x00, 0x88, 0xff}},
		{Color: color.RGBA{0x88, 0x00, 0xcc, 0xff}},
		{Color: color.RGBA{0x88, 0x44, 0x00, 0xff}},
		{Color: color.RGBA{0x88, 0x44, 0x44, 0xff}},
		{Color: color.RGBA{0x88, 0x44, 0x88, 0xff}},
		{Color: color.RGBA{0x88, 0x44, 0xcc, 0xff}},
		{Color: color.RGBA{0x88, 0x88, 0x00, 0xff}},
		{Color: color.RGBA{0x88, 0x88, 0x44, 0xff}},
		{Color: color.RGBA{0x88, 0x88, 0x88, 0xff}},
		{Color: color.RGBA{0x88, 0x88, 0xcc, 0xff}},
		{Color: color.RGBA{0x88, 0xcc, 0x00, 0xff}},
		{Color: color.RGBA{0x88, 0xcc, 0x44, 0xff}},
		{Color: color.RGBA{0x88, 0xcc, 0x88, 0xff}},
		{Color: color.RGBA{0x88, 0xcc, 0xcc, 0xff}},
		{Color: color.RGBA{0x88, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0x88, 0x00, 0x44, 0xff}},
		{Color: color.RGBA{0x99, 0x00, 0x4c, 0xff}},
		{Color: color.RGBA{0x99, 0x00, 0x99, 0xff}},
		{Color: color.RGBA{0x93, 0x00, 0xdd, 0xff}},
		{Color: color.RGBA{0x99, 0x4c, 0x00, 0xff}},
		{Color: color.RGBA{0x99, 0x4c, 0x4c, 0xff}},
		{Color: color.RGBA{0x99, 0x4c, 0x99, 0xff}},
		{Color: color.RGBA{0x93, 0x49, 0xdd, 0xff}},
		{Color: color.RGBA{0x99, 0x99, 0x00, 0xff}},
		{Color: color.RGBA{0x99, 0x99, 0x4c, 0xff}},
		{Color: color.RGBA{0x99, 0x99, 0x99, 0xff}},
		{Color: color.RGBA{0x93, 0x93, 0xdd, 0xff}},
		{Color: color.RGBA{0x93, 0xdd, 0x00, 0xff}},
		{Color: color.RGBA{0x93, 0xdd, 0x49, 0xff}},
		{Color: color.RGBA{0x93, 0xdd, 0x93, 0xff}},
		{Color: color.RGBA{0x93, 0xdd, 0xdd, 0xff}},
		{Color: color.RGBA{0x99, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0xaa, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0xaa, 0x00, 0x55, 0xff}},
		{Color: color.RGBA{0xaa, 0x00, 0xaa, 0xff}},
		{Color: color.RGBA{0x9e, 0x00, 0xee, 0xff}},
		{Color: color.RGBA{0xaa, 0x55, 0x00, 0xff}},
		{Color: color.RGBA{0xaa, 0x55, 0x55, 0xff}},
		{Color: color.RGBA{0xaa, 0x55, 0xaa, 0xff}},
		{Color: color.RGBA{0x9e, 0x4f, 0xee, 0xff}},
		{Color: color.RGBA{0xaa, 0xaa, 0x00, 0xff}},
		{Color: color.RGBA{0xaa, 0xaa, 0x55, 0xff}},
		{Color: color.RGBA{0xaa, 0xaa, 0xaa, 0xff}},
		{Color: color.RGBA{0x9e, 0x9e, 0xee, 0xff}},
		{Color: color.RGBA{0x9e, 0xee, 0x00, 0xff}},
		{Color: color.RGBA{0x9e, 0xee, 0x4f, 0xff}},
		{Color: color.RGBA{0x9e, 0xee, 0x9e, 0xff}},
		{Color: color.RGBA{0x9e, 0xee, 0xee, 0xff}},
		{Color: color.RGBA{0xaa, 0xff, 0xff, 0xff}},
		{Color: color.RGBA{0xbb, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0xbb, 0x00, 0x5d, 0xff}},
		{Color: color.RGBA{0xbb, 0x00, 0xbb, 0xff}},
		{Color: color.RGBA{0xaa, 0x00, 0xff, 0xff}},
		{Color: color.RGBA{0xbb, 0x5d, 0x00, 0xff}},
		{Color: color.RGBA{0xbb, 0x5d, 0x5d, 0xff}},
		{Color: color.RGBA{0xbb, 0x5d, 0xbb, 0xff}},
		{Color: color.RGBA{0xaa, 0x55, 0xff, 0xff}},
		{Color: color.RGBA{0xbb, 0xbb, 0x00, 0xff}},
		{Color: color.RGBA{0xbb, 0xbb, 0x5d, 0xff}},
		{Color: color.RGBA{0xbb, 0xbb, 0xbb, 0xff}},
		{Color: color.RGBA{0xaa, 0xaa, 0xff, 0xff}},
		{Color: color.RGBA{0xaa, 0xff, 0x00, 0xff}},
		{Color: color.RGBA{0xaa, 0xff, 0x55, 0xff}},
		{Color: color.RGBA{0xaa, 0xff, 0xaa, 0xff}},
		{Color: color.RGBA{0xcc, 0x00, 0xcc, 0xff}},
		{Color: color.RGBA{0xcc, 0x44, 0x00, 0xff}},
		{Color: color.RGBA{0xcc, 0x44, 0x44, 0xff}},
		{Color: color.RGBA{0xcc, 0x44, 0x88, 0xff}},
		{Color: color.RGBA{0xcc, 0x44, 0xcc, 0xff}},
		{Color: color.RGBA{0xcc, 0x88, 0x00, 0xff}},
		{Color: color.RGBA{0xcc, 0x88, 0x44, 0xff}},
		{Color: color.RGBA{0xcc, 0x88, 0x88, 0xff}},
		{Color: color.RGBA{0xcc, 0x88, 0xcc, 0xff}},
		{Color: color.RGBA{0xcc, 0xcc, 0x00, 0xff}},
		{Color: color.RGBA{0xcc, 0xcc, 0x44, 0xff}},
		{Color: color.RGBA{0xcc, 0xcc, 0x88, 0xff}},
		{Color: color.RGBA{0xcc, 0xcc, 0xcc, 0xff}},
		{Color: color.RGBA{0xcc, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0xcc, 0x00, 0x44, 0xff}},
		{Color: color.RGBA{0xcc, 0x00, 0x88, 0xff}},
		{Color: color.RGBA{0xdd, 0x00, 0x93, 0xff}},
		{Color: color.RGBA{0xdd, 0x00, 0xdd, 0xff}},
		{Color: color.RGBA{0xdd, 0x49, 0x00, 0xff}},
		{Color: color.RGBA{0xdd, 0x49, 0x49, 0xff}},
		{Color: color.RGBA{0xdd, 0x49, 0x93, 0xff}},
		{Color: color.RGBA{0xdd, 0x49, 0xdd, 0xff}},
		{Color: color.RGBA{0xdd, 0x93, 0x00, 0xff}},
		{Color: color.RGBA{0xdd, 0x93, 0x49, 0xff}},
		{Color: color.RGBA{0xdd, 0x93, 0x93, 0xff}},
		{Color: color.RGBA{0xdd, 0x93, 0xdd, 0xff}},
		{Color: color.RGBA{0xdd, 0xdd, 0x00, 0xff}},
		{Color: color.RGBA{0xdd, 0xdd, 0x49, 0xff}},
		{Color: color.RGBA{0xdd, 0xdd, 0x93, 0xff}},
		{Color: color.RGBA{0xdd, 0xdd, 0xdd, 0xff}},
		{Color: color.RGBA{0xdd, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0xdd, 0x00, 0x49, 0xff}},
		{Color: color.RGBA{0xee, 0x00, 0x4f, 0xff}},
		{Color: color.RGBA{0xee, 0x00, 0x9e, 0xff}},
		{Color: color.RGBA{0xee, 0x00, 0xee, 0xff}},
		{Color: color.RGBA{0xee, 0x4f, 0x00, 0xff}},
		{Color: color.RGBA{0xee, 0x4f, 0x4f, 0xff}},
		{Color: color.RGBA{0xee, 0x4f, 0x9e, 0xff}},
		{Color: color.RGBA{0xee, 0x4f, 0xee, 0xff}},
		{Color: color.RGBA{0xee, 0x9e, 0x00, 0xff}},
		{Color: color.RGBA{0xee, 0x9e, 0x4f, 0xff}},
		{Color: color.RGBA{0xee, 0x9e, 0x9e, 0xff}},
		{Color: color.RGBA{0xee, 0x9e, 0xee, 0xff}},
		{Color: color.RGBA{0xee, 0xee, 0x00, 0xff}},
		{Color: color.RGBA{0xee, 0xee, 0x4f, 0xff}},
		{Color: color.RGBA{0xee, 0xee, 0x9e, 0xff}},
		{Color: color.RGBA{0xee, 0xee, 0xee, 0xff}},
		{Color: color.RGBA{0xee, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0xff, 0x00, 0x00, 0xff}},
		{Color: color.RGBA{0xff, 0x00, 0x55, 0xff}},
		{Color: color.RGBA{0xff, 0x00, 0xaa, 0xff}},
		{Color: color.RGBA{0xff, 0x00, 0xff, 0xff}},
		{Color: color.RGBA{0xff, 0x55, 0x00, 0xff}},
		{Color: color.RGBA{0xff, 0x55, 0x55, 0xff}},
		{Color: color.RGBA{0xff, 0x55, 0xaa, 0xff}},
		{Color: color.RGBA{0xff, 0x55, 0xff, 0xff}},
		{Color: color.RGBA{0xff, 0xaa, 0x00, 0xff}},
		{Color: color.RGBA{0xff, 0xaa, 0x55, 0xff}},
		{Color: color.RGBA{0xff, 0xaa, 0xaa, 0xff}},
		{Color: color.RGBA{0xff, 0xaa, 0xff, 0xff}},
		{Color: color.RGBA{0xff, 0xff, 0x00, 0xff}},
		{Color: color.RGBA{0xff, 0xff, 0x55, 0xff}},
		{Color: color.RGBA{0xff, 0xff, 0xaa, 0xff}},
		{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}},
	}},
}
