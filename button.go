package iui

import "image"

type Button struct {
	Child Component
	Click func()
}

func (btn Button) Size(avail image.Point) image.Point {
	return btn.Child.Size(avail)
}
func (btn Button) Draw(ctx DrawContext) {
	for range ctx.Clicks {
		btn.Click()
	}
	btn.Child.Draw(ctx)
}
