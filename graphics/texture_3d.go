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

	"github.com/haakenlabs/arc/system/instance"
)

var _ Texture = &Texture3D{}

type Texture3D struct {
	BaseTexture
}

func NewTexture3D(cfg *TextureConfig) *Texture3D {
	t := &Texture3D{}

	t.textureType = gl.TEXTURE_3D

	t.SetName("Texture3D")
	instance.MustAssign(t)

	t.size = cfg.Size
	t.uploadFunc = t.Upload

	t.internalFormat = TextureFormatToInternal(cfg.Format)
	t.glFormat = TextureFormatToFormat(cfg.Format)
	t.storageFormat = TextureFormatToStorage(cfg.Format)

	return t
}

func (t *Texture3D) Type() TextureType {
	return TextureType3D
}

func (t *Texture3D) Upload() {
	t.Bind()

	gl.TexImage3D(t.textureType, 0, t.internalFormat, t.size.X(), t.size.Y(), t.layers, 0, t.glFormat, t.storageFormat, nil)
}
