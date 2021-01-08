package iui

import (
	"image"

	"github.com/vktec/gll"
)

type UI struct {
	gl   gll.GL300
	box  image.Rectangle
	vao  uint32
	sreg ShaderRegistry

	mouse  image.Point
	clicks []image.Point
}

func New(gl gll.GL300, x0, y0, x1, y1 int) (ui *UI, err error) {
	ui = &UI{gl: gl}
	ui.box = image.Rect(x0, y0, x1, y1)

	ui.sreg, err = buildShaders(gl)
	if err != nil {
		return nil, err
	}

	gl.GenVertexArrays(1, &ui.vao)
	gl.BindVertexArray(ui.vao)

	return ui, nil
}

func (ui *UI) MoveMouse(x, y int) {
	ui.mouse = image.Pt(x, ui.box.Max.Y-y)
}
func (ui *UI) Click() {
	ui.clicks = append(ui.clicks, ui.mouse)
}

func (ui *UI) Draw(comp Component) {
	box := ui.box
	size := comp.Size(box.Size())
	dx, dy := box.Dx()-size.X, box.Dy()-size.Y
	if dx < 0 || dy < 0 {
		panic("Component reported size greater than available area")
	}

	// Center within available space
	box.Min.X += dx / 2
	box.Max.X -= dx - dx/2
	box.Min.Y += dy / 2
	box.Max.Y -= dy - dy/2

	// Constrain rendering to box
	ui.gl.Enable(gll.SCISSOR_TEST)
	defer ui.gl.Disable(gll.SCISSOR_TEST)
	ui.gl.Scissor(scissorBox(box))

	// Make transparency work
	ui.gl.Enable(gll.BLEND)
	defer ui.gl.Disable(gll.BLEND)
	ui.gl.BlendFunc(gll.ONE, gll.ONE_MINUS_SRC_ALPHA)

	comp.Draw(DrawContext{ui.gl, ui.sreg, box, ui.mouse, ui.clicks})
	ui.clicks = ui.clicks[:0]
}
