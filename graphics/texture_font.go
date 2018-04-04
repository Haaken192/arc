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

	"github.com/haakenlabs/arc/pkg/math"
	"github.com/haakenlabs/arc/system/instance"
)

type TextureFont struct {
	BaseTexture

	data []uint8
}

func NewTextureFont(size math.IVec2) *TextureFont {
	t := &TextureFont{}

	t.textureType = gl.TEXTURE_2D

	t.SetName("TextureFont")
	instance.MustAssign(t)

	t.size = size
	t.uploadFunc = t.Upload

	t.internalFormat = gl.RGBA8
	t.glFormat = gl.RGBA
	t.storageFormat = gl.UNSIGNED_BYTE

	return t
}

func (t *TextureFont) Upload() {
	t.Bind()

	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	if t.data != nil && len(t.data) > 0 {
		gl.TexImage2D(gl.TEXTURE_2D, 0, t.internalFormat, t.size.X(), t.size.Y(), 0, t.glFormat, t.storageFormat, gl.Ptr(t.data))
	} else {
		gl.TexImage2D(gl.TEXTURE_2D, 0, t.internalFormat, t.size.X(), t.size.Y(), 0, t.glFormat, t.storageFormat, nil)
	}
}

func (t *TextureFont) SetData(data []uint8) {
	t.data = data
}
