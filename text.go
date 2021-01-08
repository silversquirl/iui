package iui

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/vktec/gll"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

type Text struct {
	Text      string
	Color     color.Color
	PointSize float64

	cached      bool
	faceCache   font.Face
	boundsCache image.Rectangle
}

var textFont *opentype.Font

func init() {
	var err error
	textFont, err = opentype.Parse(goregular.TTF)
	if err != nil {
		panic("Error parsing font: " + err.Error())
	}
}

var textShaderId = RegisterShader(`
#version 150
uniform sampler2DRect tex;
uniform vec2 pos;
out vec4 color;
void main() {
	vec2 texCoord = gl_FragCoord.xy - pos;
	float texh = float(textureSize(tex).y);
	texCoord.y = texh - texCoord.y;
	color = texture(tex, texCoord);
}
`)

func (t *Text) face() font.Face {
	if !t.cached {
		var err error
		t.faceCache, err = opentype.NewFace(textFont, &opentype.FaceOptions{
			Size:    t.PointSize,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		if err != nil {
			panic("Error creating face: " + err.Error())
		}
		t.cached = true
	}
	return t.faceCache
}

// TODO: text wrapping
func (t *Text) bounds() image.Rectangle {
	if t.PointSize == 0 {
		t.PointSize = 12
	}
	if !t.cached {
		face := t.face()
		bounds, _ := font.BoundString(face, t.Text)
		t.boundsCache = image.Rect(
			bounds.Min.X.Floor(), bounds.Min.Y.Floor(),
			bounds.Max.X.Ceil(), bounds.Max.Y.Ceil(),
		)
		t.cached = true
	}
	return t.boundsCache
}
func (t *Text) Size(avail image.Point) image.Point {
	size := t.bounds().Size()
	if size.X > avail.X {
		size.X = avail.X
	}
	if size.Y > avail.Y {
		size.Y = avail.Y
	}
	return size
}

func (t *Text) Draw(ctx DrawContext) {
	if t.Color == nil {
		t.Color = color.White
	}

	// Draw text
	bounds := t.bounds()
	dst := image.NewRGBA(bounds)
	draw := font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(t.Color),
		Face: t.face(),
	}
	draw.DrawString(t.Text)
	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	png.Encode(f, dst)
	f.Close()

	// Upload texture
	ctx.ActiveTexture(gll.TEXTURE0)
	var tex uint32
	ctx.GenTextures(1, &tex)
	defer ctx.DeleteTextures(1, &tex)
	ctx.BindTexture(gll.TEXTURE_RECTANGLE, tex)
	ctx.TexImage2D(gll.TEXTURE_RECTANGLE, 0, gll.RGBA, int32(bounds.Dx()), int32(bounds.Dy()), 0, gll.RGBA, gll.UNSIGNED_BYTE, gll.Ptr(dst.Pix))

	prog := ctx.GetShader(textShaderId)
	ctx.UseProgram(prog)
	ctx.Uniform1i(ctx.GetUniformLocation(prog, gll.Str("tex\000")), 0)
	ctx.Uniform2f(ctx.GetUniformLocation(prog, gll.Str("pos\000")), float32(ctx.Box.Min.X), float32(ctx.Box.Min.Y))
	ctx.DrawArrays(gll.TRIANGLES, 0, 3)
}
