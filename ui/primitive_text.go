/*
Copyright (c) 2018 HaakenLabs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package ui

import (
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/haakenlabs/forge"
	"github.com/haakenlabs/forge/system/asset/font"
	"github.com/haakenlabs/forge/system/asset/shader"
)

var _ Primitive = &Text{}

type Text struct {
	BasePrimitive

	font      *forge.Font
	fontSize  int32
	color     forge.Color
	value     string
	maskLayer uint8
}

func (t *Text) Font() *forge.Font {
	return t.font
}

func (t *Text) FontSize() int32 {
	return t.fontSize
}

func (t *Text) Value() string {
	return t.value
}

func (t *Text) SetFont(font *forge.Font) {
	t.font = font
}

func (t *Text) SetFontSize(size int32) {
	if size < 1 {
		size = 1
	}
	t.fontSize = size
}

func (t *Text) SetValue(value string) {
	t.value = value
}

func (t *Text) SetColor(color forge.Color) {
	t.color = color
}

func (t *Text) Color() forge.Color {
	return t.color
}

func (t *Text) Refresh() {
	if t.font == nil {
		return
	}

	vertices, bounds := t.font.DrawText(t.value, float64(t.fontSize))
	fa := t.font.Atlas(float64(t.fontSize))

	t.rect.SetSize(bounds)

	t.material.SetTexture(0, fa.Texture())
	t.mesh.Upload(vertices)
}

func (t *Text) Draw(matrix mgl32.Mat4) {
	if t.material == nil || t.mesh.size == 0 {
		return
	}

	t.material.Bind()
	t.mesh.Bind()

	t.material.SetProperty("v_ortho_matrix", forge.GetWindow().OrthoMatrix())
	t.material.SetProperty("v_model_matrix", matrix.Mul4(t.rect.Matrix()))
	t.material.SetProperty("f_alpha", float32(1.0))
	t.material.SetProperty("f_color", t.color.Vec4())

	gl.StencilFunc(gl.ALWAYS, int32(t.maskLayer), 0xFF)
	gl.StencilMask(0)

	t.mesh.Draw()

	t.mesh.Unbind()
	t.material.Unbind()
}

func NewText() *Text {
	t := &Text{
		color:    Styles.TextColor,
		fontSize: Styles.TextSize,
	}

	t.material = forge.NewMaterial()
	t.material.SetShader(shader.MustGet("ui/text"))

	t.font = font.MustGet("SourceCodePro-Regular.ttf")

	t.mesh = NewMesh()
	t.mesh.Alloc()

	return t
}
