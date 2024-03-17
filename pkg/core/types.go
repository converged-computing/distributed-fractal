package core

// A mandel iteration holds values that a worker needs to calculate a single pixel value
type MandelIteration struct {
	Xmin         float64
	Ymin         float64
	Xmax         float64
	Ymax         float64
	IndexY       int
	Width        int
	MaxIteration int
}

type IterationResult struct {
	Norm         []float64
	IndexX       int
	IndexY       int
	MaxIteration int32
	ResultIt     []int
}
