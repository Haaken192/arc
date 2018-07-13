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
	"github.com/haakenlabs/arc/pkg/math"
)

type BaseTexture struct {
	core.BaseObject

	uploadFunc     UploadFunc
	internalFormat int32
	storageFormat  uint32
	glFormat       uint32
	filterMag      int32
	filterMin      int32
	wrapR          int32
	wrapS          int32
	wrapT          int32
	layers         int32
	reference      uint32
	textureFormat  TextureFormat
	size           math.IVec2
	resizable      bool
	textureType    uint32
}

func (t *BaseTexture) Alloc() error {
	if t.reference != 0 {
		return nil
	}

	gl.GenTextures(1, &t.reference)

	t.filterMag = gl.LINEAR
	t.filterMin = gl.LINEAR
	t.wrapR = gl.CLAMP_TO_EDGE
	t.wrapS = gl.CLAMP_TO_EDGE
	t.wrapT = gl.CLAMP_TO_EDGE
	t.resizable = true
	t.layers = 1

	t.uploadFunc()

	t.SetFilter(t.filterMag, t.filterMin)
	t.SetWrapRST(t.wrapR, t.wrapS, t.wrapT)

	return nil
}

func (t *BaseTexture) Dealloc() {
	if t.reference != 0 {
		gl.DeleteTextures(1, &t.reference)
		t.reference = 0
	}
}

func (t *BaseTexture) Activate(unit uint32) {
	gl.ActiveTexture(unit)
	t.Bind()
}

func (t *BaseTexture) Bind() {
	gl.BindTexture(t.textureType, t.reference)
}

func (t *BaseTexture) FilterMag() int32 {
	return t.filterMag
}

func (t *BaseTexture) FilterMin() int32 {
	return t.filterMin
}

func (t *BaseTexture) GenerateMipmaps() {}

func (t *BaseTexture) GLFormat() uint32 {
	return t.glFormat
}

func (t *BaseTexture) GLInternalFormat() int32 {
	return t.internalFormat
}

func (t *BaseTexture) GLStorageFormat() uint32 {
	return t.storageFormat
}

func (t *BaseTexture) GLType() uint32 {
	return t.textureType
}

func (t *BaseTexture) Height() int32 {
	return t.size.Y()
}

func (t *BaseTexture) Layers() int32 {
	return t.layers
}

func (t *BaseTexture) SetLayers(layers int32) {
	t.layers = layers
}

func (t *BaseTexture) MipLevels() uint32 {
	return 1
}

func (t *BaseTexture) Resize() {}

func (t *BaseTexture) Resizable() bool {
	return t.resizable
}

func (t *BaseTexture) SetFormat(format TextureFormat) {
	t.SetGLFormats(TextureFormatToInternal(format), TextureFormatToFormat(format), TextureFormatToStorage(format))
}

func (t *BaseTexture) SetFilter(magFilter, minFilter int32) {
	t.SetMagFilter(magFilter)
	t.SetMinFilter(minFilter)
}

func (t *BaseTexture) SetGLFormats(internalFormat int32, format uint32, storageFormat uint32) {
	t.internalFormat = internalFormat
	t.glFormat = format
	t.storageFormat = storageFormat

	t.uploadFunc()
}

func (t *BaseTexture) SetMagFilter(magFilter int32) {
	t.filterMag = magFilter
	gl.TexParameteri(t.textureType, gl.TEXTURE_MAG_FILTER, t.filterMag)
}

func (t *BaseTexture) SetMinFilter(minFilter int32) {
	t.filterMin = minFilter
	gl.TexParameteri(t.textureType, gl.TEXTURE_MIN_FILTER, t.filterMin)
}

func (t *BaseTexture) SetResizable(resizable bool) {
	t.resizable = resizable
}

func (t *BaseTexture) SetSize(size math.IVec2) {
	t.size = size
	t.uploadFunc()
}

func (t *BaseTexture) SetTexFormat(format TextureFormat) {
	t.SetGLFormats(TextureFormatToInternal(format), TextureFormatToFormat(format), TextureFormatToStorage(format))
}

func (t *BaseTexture) SetWrapR(wrapR int32) {
	t.wrapR = wrapR
	gl.TexParameteri(t.textureType, gl.TEXTURE_WRAP_R, t.wrapR)
	if t.wrapR == gl.CLAMP_TO_BORDER {
		color := [4]float32{}
		gl.TexParameterfv(t.textureType, gl.TEXTURE_BORDER_COLOR, &color[0])
	}
}

func (t *BaseTexture) SetWrapRST(wrapR, wrapS, wrapT int32) {
	t.SetWrapR(wrapR)
	t.SetWrapS(wrapS)
	t.SetWrapT(wrapT)
}

func (t *BaseTexture) SetWrapS(wrapS int32) {
	t.wrapS = wrapS
	gl.TexParameteri(t.textureType, gl.TEXTURE_WRAP_S, t.wrapS)
	if t.wrapS == gl.CLAMP_TO_BORDER {
		color := [4]float32{}
		gl.TexParameterfv(t.textureType, gl.TEXTURE_BORDER_COLOR, &color[0])
	}
}

func (t *BaseTexture) SetWrapST(wrapS, wrapT int32) {
	t.SetWrapS(wrapS)
	t.SetWrapT(wrapT)
}

func (t *BaseTexture) SetWrapT(wrapT int32) {
	t.wrapT = wrapT
	gl.TexParameteri(t.textureType, gl.TEXTURE_WRAP_T, t.wrapT)
	if t.wrapT == gl.CLAMP_TO_BORDER {
		color := [4]float32{}
		gl.TexParameterfv(t.textureType, gl.TEXTURE_BORDER_COLOR, &color[0])
	}
}

func (t *BaseTexture) Size() math.IVec2 {
	return t.size
}

func (t *BaseTexture) Format() TextureFormat {
	return t.textureFormat
}

func (t *BaseTexture) Reference() uint32 {
	return t.reference
}

func (t *BaseTexture) Upload() {
	panic("Unimplemented!")
}

func (t *BaseTexture) Unbind() {
	gl.BindTexture(t.textureType, 0)
}

func (t *BaseTexture) Width() int32 {
	return t.size.X()
}

func (t *BaseTexture) WrapR() int32 {
	return t.wrapR
}

func (t *BaseTexture) WrapS() int32 {
	return t.wrapS
}

func (t *BaseTexture) WrapT() int32 {
	return t.wrapT
}

func (t *BaseTexture) SetData([]uint8) {}

func (t *BaseTexture) SetLayerData([]uint8, int32) {}

func (t *BaseTexture) SetHDRData([]float32) {}

func (t *BaseTexture) SetHDRLayerData([]float32, int32) {}
