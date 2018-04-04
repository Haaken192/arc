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

package graphics

import (
	"github.com/go-gl/gl/v4.3-core/gl"

	"github.com/haakenlabs/arc/core"
	"github.com/haakenlabs/arc/system/instance"
)

type TextureColor struct {
	BaseTexture

	color core.Color
}

func NewTextureColor(color core.Color) *TextureColor {
	t := &TextureColor{}

	t.textureType = gl.TEXTURE_3D

	t.SetName("TextureColor")
	instance.MustAssign(t)

	t.uploadFunc = t.Upload

	return t
}

func (t *TextureColor) Upload() {
	t.Bind()
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA32F, t.size.X(), t.size.Y(), 0, gl.RGBA, gl.FLOAT, gl.Ptr(t.color))
}

func (t *TextureColor) Color() core.Color {
	return t.color
}

func (t *TextureColor) SetColor(color core.Color) {
	t.color = color
	t.uploadFunc()
}
