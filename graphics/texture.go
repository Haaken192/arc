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

type TextureConfig struct {
	Type   TextureType
	Format TextureFormat
	Layers int32
	Size   math.IVec2
}

type TextureType uint8

const (
	TextureType2D TextureType = iota
	TextureType3D
	TextureTypeCubemap
	TextureTypeFont
	TextureTypeColor
)

type TextureFormat uint32

const (
	TextureFormatDefaultColor TextureFormat = iota
	TextureFormatDefaultHDRColor
	TextureFormatDefaultDepth
	TextureFormatR8
	TextureFormatRG8
	TextureFormatRGB8
	TextureFormatRGBA8
	TextureFormatR16
	TextureFormatRG16
	TextureFormatRGB16
	TextureFormatRGBA16
	TextureFormatRGBA16UI
	TextureFormatR32
	TextureFormatRG32
	TextureFormatRGB32
	TextureFormatRGBA32
	TextureFormatRGB32UI
	TextureFormatRGBA32UI
	TextureFormatDepth16
	TextureFormatDepth24
	TextureFormatDepth24Stencil8
	TextureFormatStencil8
)

type UploadFunc func()

type Texture interface {
	core.Object
	Binder
	Sizable

	// Type returns the TextureType of this texture.
	Type() TextureType

	// GLType() returns the OpenGL texture type.
	GLType() uint32

	// Format returns the TextureFormat of this texture.
	Format() TextureFormat

	// SetMagFilter sets the magnification filter.
	SetMagFilter(int32)

	// SetMinFilter sets the minification filter.
	SetMinFilter(int32)

	// FilterMag returns the magnification filter value.
	FilterMag() int32

	// FilterMin returns the minification filter value.
	FilterMin() int32

	// WrapR returns the wrap value for the R dimension.
	WrapR() int32

	// WrapS returns the wrap value for the S dimension.
	WrapS() int32

	// WrapT returns the wrap value for the T dimension.
	WrapT() int32

	// SetWrapR sets the wrap value for the S dimension.
	SetWrapR(int32)

	// SetWrapRST sets the wrap value for the R, S, and T dimensions.
	SetWrapRST(int32, int32, int32)

	// SetWrapS sets the wrap value for the S dimension.
	SetWrapS(int32)

	// SetWrapST sets the wrap value for the S and T dimensions.
	SetWrapST(int32, int32)

	// SetWrapT sets the wrap value for the T dimension.
	SetWrapT(int32)

	// Activate binds and sets the texture active for the provided unit.
	Activate(uint32)

	// MipLevels returns the number of mip maps for this texture.
	MipLevels() uint32

	// GenerateMipmaps will generate mip maps for the texture (if supported).
	GenerateMipmaps()

	// Layers returns the number of layers for this texture.
	Layers() int32

	// SetLayers sets the number of layers for this texture (if supported).
	SetLayers(int32)

	// SetFormat sets the format for this texture (if supported).
	SetFormat(TextureFormat)

	// SetData sets the data for this texture (LDR-values only).
	SetData([]uint8)

	// SetLayerData sets the data for the texture at a specific layer (LDR-values only).
	SetLayerData([]uint8, int32)

	// SetHDRData sets the data for this texture (HDR-values only).
	SetHDRData([]float32)

	// SetHDRLayerData sets the data for the texture at a specific layer (HDR-values only).
	SetHDRLayerData([]float32, int32)
}

func TextureFormatToInternal(format TextureFormat) int32 {
	switch format {
	case TextureFormatR8:
		return gl.R8
	case TextureFormatRG8:
		return gl.RG8
	case TextureFormatRGB8:
		return gl.RGB8
	case TextureFormatDefaultColor:
		fallthrough
	case TextureFormatRGBA8:
		return gl.RGBA8
	case TextureFormatR16:
		return gl.R16F
	case TextureFormatRG16:
		return gl.RG16F
	case TextureFormatRGB16:
		return gl.RGB16F
	case TextureFormatDefaultHDRColor:
		fallthrough
	case TextureFormatRGBA16:
		return gl.RGBA16F
	case TextureFormatR32:
		return gl.R32F
	case TextureFormatRG32:
		return gl.RG32F
	case TextureFormatRGB32:
		return gl.RGB32F
	case TextureFormatRGBA32:
		return gl.RGBA32F
	case TextureFormatRGB32UI:
		return gl.RGB32UI
	case TextureFormatRGBA32UI:
		return gl.RGBA32UI
	case TextureFormatDepth16:
		return gl.DEPTH_COMPONENT16
	case TextureFormatDefaultDepth:
		fallthrough
	case TextureFormatDepth24:
		return gl.DEPTH_COMPONENT24
	case TextureFormatDepth24Stencil8:
		return gl.DEPTH24_STENCIL8
	case TextureFormatStencil8:
		return gl.STENCIL_INDEX8
	case TextureFormatRGBA16UI:
		return gl.RGBA16UI
	}

	return 0
}

func TextureFormatToFormat(format TextureFormat) uint32 {
	switch format {
	case TextureFormatR8:
		fallthrough
	case TextureFormatR16:
		fallthrough
	case TextureFormatR32:
		return gl.RED
	case TextureFormatRG8:
		fallthrough
	case TextureFormatRG16:
		fallthrough
	case TextureFormatRG32:
		return gl.RG
	case TextureFormatRGB8:
		fallthrough
	case TextureFormatRGB16:
		fallthrough
	case TextureFormatRGB32:
		return gl.RGB
	case TextureFormatRGB32UI:
		return gl.RGB_INTEGER
	case TextureFormatDefaultColor:
		fallthrough
	case TextureFormatRGBA8:
		fallthrough
	case TextureFormatDefaultHDRColor:
		fallthrough
	case TextureFormatRGBA16:
		fallthrough
	case TextureFormatRGBA16UI:
		fallthrough
	case TextureFormatRGBA32:
		return gl.RGBA
	case TextureFormatRGBA32UI:
		return gl.RGBA_INTEGER
	case TextureFormatDefaultDepth:
		fallthrough
	case TextureFormatDepth16:
		fallthrough
	case TextureFormatDepth24:
		return gl.DEPTH_COMPONENT
	case TextureFormatDepth24Stencil8:
		fallthrough
	case TextureFormatStencil8:
		return 0
	}

	return 0
}

func TextureFormatToStorage(format TextureFormat) uint32 {
	switch format {
	case TextureFormatDefaultColor:
		fallthrough
	case TextureFormatR8:
		fallthrough
	case TextureFormatRG8:
		fallthrough
	case TextureFormatRGB8:
		fallthrough
	case TextureFormatRGBA8:
		fallthrough
	case TextureFormatStencil8:
		return gl.UNSIGNED_BYTE
	case TextureFormatR16:
		fallthrough
	case TextureFormatRG16:
		fallthrough
	case TextureFormatRGB16:
		fallthrough
	case TextureFormatDefaultHDRColor:
		fallthrough
	case TextureFormatRGBA16:
		return gl.HALF_FLOAT
	case TextureFormatRGBA16UI:
		return gl.UNSIGNED_SHORT
	case TextureFormatR32:
		fallthrough
	case TextureFormatRG32:
		fallthrough
	case TextureFormatRGB32:
		fallthrough
	case TextureFormatRGBA32:
		return gl.FLOAT
	case TextureFormatRGB32UI:
		fallthrough
	case TextureFormatRGBA32UI:
		return gl.UNSIGNED_INT
	case TextureFormatDefaultDepth:
		fallthrough
	case TextureFormatDepth16:
		fallthrough
	case TextureFormatDepth24:
		return gl.FLOAT
	case TextureFormatDepth24Stencil8:
		return gl.UNSIGNED_INT_24_8
	}

	return 0
}

func NewTexture(cfg *TextureConfig) Texture {
	switch cfg.Type {
	case TextureType2D:
		return NewTexture2D(cfg)
	case TextureType3D:
		return NewTexture3D(cfg)
	case TextureTypeCubemap:
		return NewTextureCubemap(cfg)
	case TextureTypeFont:
		return NewTextureFont(cfg)
	case TextureTypeColor:
		return NewTextureColor()
	default:
		return nil
	}
}
