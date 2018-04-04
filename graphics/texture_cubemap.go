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

type TextureCubemap struct {
	BaseTexture

	data    [6][]uint8
	hdrData [6][]float32
}

func NewTextureCubemap(size math.IVec2, format TextureFormat) *TextureCubemap {
	t := &TextureCubemap{}

	t.data = [6][]uint8{}
	t.hdrData = [6][]float32{}

	t.textureType = gl.TEXTURE_CUBE_MAP

	t.SetName("TextureCubemap")
	instance.MustAssign(t)

	t.size = size
	t.uploadFunc = t.Upload

	t.internalFormat = TextureFormatToInternal(format)
	t.glFormat = TextureFormatToFormat(format)
	t.storageFormat = TextureFormatToStorage(format)

	return t
}

func (t *TextureCubemap) SetData(data []byte, offset int) {
	if offset > 5 {
		return
	}

	t.data[offset] = data
}

func (t *TextureCubemap) SetHDRData(data []float32, offset int) {
	if offset > 5 {
		return
	}

	t.hdrData[offset] = data
}

func (t *TextureCubemap) Upload() {
	t.Bind()

	if len(t.hdrData[0]) > 0 {
		for i := range t.hdrData {
			gl.TexImage2D(
				gl.TEXTURE_CUBE_MAP_POSITIVE_X+uint32(i),
				0,
				t.internalFormat,
				t.size.X(),
				t.size.Y(),
				0,
				t.glFormat,
				t.storageFormat,
				gl.Ptr(t.hdrData[i]),
			)
		}
	} else if len(t.data[0]) > 0 {
		for i := range t.data {
			gl.TexImage2D(
				gl.TEXTURE_CUBE_MAP_POSITIVE_X+uint32(i),
				0,
				t.internalFormat,
				t.size.X(),
				t.size.Y(),
				0,
				t.glFormat,
				t.storageFormat,
				gl.Ptr(t.data[i]),
			)
		}
	} else {
		for i := uint32(0); i < 6; i++ {
			gl.TexImage2D(
				gl.TEXTURE_CUBE_MAP_POSITIVE_X+uint32(i),
				0,
				t.internalFormat,
				t.size.X(),
				t.size.Y(),
				0,
				t.glFormat,
				t.storageFormat,
				nil,
			)
		}
	}
}
