package iui

import "image"

type Size struct {
	W, H  int
	Child Component
}

func (size Size) Size(avail image.Point) image.Point {
	if size.W < avail.X {
		avail.X = size.W
	}
	if size.H < avail.Y {
		avail.Y = size.H
	}
	return size.Child.Size(avail)
}

func (size Size) Draw(ctx DrawContext) {
	size.Child.Draw(ctx)
}
