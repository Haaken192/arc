/*
Copyright (c) 2017 HaakenLabs

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

package skybox

import (
	"encoding/json"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"path/filepath"
	"sync"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sirupsen/logrus"

	"github.com/haakenlabs/arc/core"
	"github.com/haakenlabs/arc/graphics"
	"github.com/haakenlabs/arc/pkg/image/hdr"
	"github.com/haakenlabs/arc/pkg/math"
	"github.com/haakenlabs/arc/scene"
	"github.com/haakenlabs/arc/system/asset"
	"github.com/haakenlabs/arc/system/asset/shader"

	_ "image/jpeg"
	_ "image/png"
)

const (
	AssetNameSkybox = "skybox"
)

var rotMatrices = [6]mgl32.Mat4{
	// X
	mgl32.LookAtV(mgl32.Vec3{}, mgl32.Vec3{1, 0, 0}, mgl32.Vec3{0, 1, 0}),
	mgl32.LookAtV(mgl32.Vec3{}, mgl32.Vec3{-1, 0, 0}, mgl32.Vec3{0, 1, 0}),
	// Y
	mgl32.LookAtV(mgl32.Vec3{}, mgl32.Vec3{0, 1, 0}, mgl32.Vec3{0, 0, -1}),
	mgl32.LookAtV(mgl32.Vec3{}, mgl32.Vec3{0, -1, 0}, mgl32.Vec3{0, 0, 1}),
	// Z
	mgl32.LookAtV(mgl32.Vec3{}, mgl32.Vec3{0, 0, 1}, mgl32.Vec3{0, 1, 0}),
	mgl32.LookAtV(mgl32.Vec3{}, mgl32.Vec3{0, 0, -1}, mgl32.Vec3{0, 1, 0}),
}

var _ core.AssetHandler = &Handler{}

type Metadata struct {
	Name       string `json:"name"`
	Radiance   string `json:"radiance"`
	Specular   string `json:"specular"`
	Irradiance string `json:"irradiance"`
}

type Handler struct {
	core.BaseAssetHandler
}

func NewSkyboxHandler() *Handler {
	h := &Handler{}
	h.Items = make(map[string]int32)
	h.Mu = &sync.RWMutex{}

	return h
}

func (h *Handler) Name() string {
	return AssetNameSkybox
}

func (h *Handler) Load(r *core.Resource) error {
	m := &Metadata{}

	if err := json.Unmarshal(r.Bytes(), m); err != nil {
		return err
	}

	if _, dup := h.Items[m.Name]; dup {
		return core.ErrAssetExists(m.Name)
	}

	skybox, err := h.loadMap(m, r.DirPrefix())
	if err != nil {
		return err
	}

	h.Items[m.Name] = skybox.ID()

	return nil
}

func (h *Handler) loadMap(m *Metadata, dir string) (skybox *scene.Skybox, err error) {
	var specR, irrdR *core.Resource
	var specTex, irrdTex *graphics.Texture2D
	var radiance, specular, irradiance *graphics.TextureCubemap

	genSpecular := len(m.Specular) == 0
	genIrradiance := len(m.Irradiance) == 0

	radiR, err := core.NewResource(filepath.Join(dir, m.Radiance))
	if err != nil {
		return nil, err
	}
	if err := asset.ReadResource(radiR); err != nil {
		return nil, err
	}

	if !genSpecular {
		specR, err = core.NewResource(filepath.Join(dir, m.Specular))
		if err != nil {
			return nil, err
		}
		if err := asset.ReadResource(specR); err != nil {
			return nil, err
		}
	}
	if !genIrradiance {
		irrdR, err = core.NewResource(filepath.Join(dir, m.Irradiance))
		if err != nil {
			return nil, err
		}
		if err := asset.ReadResource(irrdR); err != nil {
			return nil, err
		}
	}

	rimg := loadImage(radiR)
	simg := loadImage(specR)
	iimg := loadImage(irrdR)

	radiTex, err := loadTexture(rimg)
	if err != nil {
		return nil, err
	}
	if !genSpecular {
		specTex, err = loadTexture(simg)
		if err != nil {
			return nil, err
		}
	}
	if !genIrradiance {
		irrdTex, err = loadTexture(iimg)
		if err != nil {
			return nil, err
		}
	}

	fbo := graphics.NewFramebufferRaw()
	defer fbo.Dealloc()

	radiance, err = makeCubemap(radiTex, fbo, radiTex.Size().Y()/2)
	if err != nil {
		return nil, err
	}

	if genSpecular {
		specular, err = generateSpecular(radiance, fbo)
	} else {
		specular, err = makeCubemap(specTex, fbo, specTex.Size().Y()/2)
	}
	if err != nil {
		return nil, err
	}

	if genIrradiance {
		irradiance, err = generateIrradiance(radiance, fbo)
	} else {
		irradiance, err = makeCubemap(irrdTex, fbo, irrdTex.Size().Y()/2)
	}
	if err != nil {
		return nil, err
	}

	skybox = scene.NewSkybox(radiance, specular, irradiance)

	return skybox, nil
}

// Get gets an asset by name.
func (h *Handler) Get(name string) (*scene.Skybox, error) {
	a, err := h.GetAsset(name)
	if err != nil {
		return nil, err
	}

	a2, ok := a.(*scene.Skybox)
	if !ok {
		return nil, core.ErrAssetType(name)
	}

	return a2, nil
}

// MustGet is like GetAsset, but panics if an error occurs.
func (h *Handler) MustGet(name string) *scene.Skybox {
	a, err := h.Get(name)
	if err != nil {
		panic(err)
	}

	return a
}

func loadImage(r *core.Resource) (img image.Image) {
	img, _, err := image.Decode(r.Reader())
	if err != nil {
		logrus.Error(err)
		return nil
	}

	return img
}

func loadTexture(img image.Image) (tex *graphics.Texture2D, err error) {
	x := int32(img.Bounds().Dx())
	y := int32(img.Bounds().Dy())

	tex = graphics.NewTexture2D(math.IVec2{x, y}, graphics.TextureFormatDefaultColor)

	switch img.ColorModel() {
	case color.CMYKModel:
		panic("color.CMYKModel unsupported")
	case color.YCbCrModel:
		rgba := image.NewRGBA(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
		tex.SetTexFormat(graphics.TextureFormatRGBA8)
		tex.SetData(rgba.Pix)
	case color.RGBA64Model:
		rgba := image.NewRGBA64(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
		tex.SetTexFormat(graphics.TextureFormatRGBA16)
		tex.SetData(rgba.Pix)
	case color.RGBAModel:
		rgba := image.NewRGBA(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
		tex.SetTexFormat(graphics.TextureFormatRGBA8)
		tex.SetData(rgba.Pix)
	case color.NRGBA64Model:
		rgba := image.NewNRGBA64(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
		tex.SetTexFormat(graphics.TextureFormatRGBA16)
		tex.SetData(rgba.Pix)
	case color.NRGBAModel:
		rgba := image.NewNRGBA(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
		tex.SetTexFormat(graphics.TextureFormatRGBA8)
		tex.SetData(rgba.Pix)
	case hdr.RGB96Model:
		rgba := hdr.NewRGB96(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
		tex.SetTexFormat(graphics.TextureFormatRGB32)

		var data []float32
		for y := 0; y < rgba.Rect.Dy(); y++ {
			for x := 0; x < rgba.Rect.Dx(); x++ {
				c := rgba.RGB96At(x, y)
				data = append(data, c.R)
				data = append(data, c.G)
				data = append(data, c.B)
			}
		}
		tex.SetHDRData(data)

	default:
		return nil, errors.New("unknown image type")
	}

	if err := tex.Alloc(); err != nil {
		return nil, err
	}

	return tex, err
}

func makeCubemap(tex *graphics.Texture2D, fbo *graphics.Framebuffer, faceSize int32) (cubemap *graphics.TextureCubemap, err error) {
	fbo.Bind()
	fbo.SetSize(math.IVec2{faceSize, faceSize})

	cubemap = graphics.NewTextureCubemap(math.IVec2{faceSize, faceSize}, tex.TexFormat())
	if err := cubemap.Alloc(); err != nil {
		return nil, err
	}

	mesh := graphics.NewMeshQuadBack()
	mesh.Bind()

	gl.Disable(gl.DEPTH_TEST)
	gl.DepthMask(false)

	s := shader.MustGet("utils/cubeconv")
	s.Bind()
	s.SetUniform("v_projection_matrix", mgl32.Perspective(math.Pi32/2.0, 1.0, 0.1, 2.0))
	tex.ActivateTexture(gl.TEXTURE0)

	fbo.ClearBuffers()

	for i := uint32(0); i < 6; i++ {
		s.SetUniform("v_view_matrix", rotMatrices[i])
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_CUBE_MAP_POSITIVE_X+i, cubemap.Reference(), 0)
		mesh.Draw()
	}

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, 0, 0)

	gl.DepthMask(true)
	gl.Enable(gl.DEPTH_TEST)

	s.Unbind()
	mesh.Unbind()
	fbo.Unbind()

	return
}

func generateSpecular(radiance *graphics.TextureCubemap, fbo *graphics.Framebuffer) (spec *graphics.TextureCubemap, err error) {
	//return nil, ErrNotImplemented

	return nil, nil
}

func generateIrradiance(radiance *graphics.TextureCubemap, fbo *graphics.Framebuffer) (irrd *graphics.TextureCubemap, err error) {
	//return nil, ErrNotImplemented

	return nil, nil
}

func Get(name string) (*scene.Skybox, error) {
	return mustHandler().Get(name)
}

func MustGet(name string) *scene.Skybox {
	return mustHandler().MustGet(name)
}

func mustHandler() *Handler {
	h, err := asset.GetHandler(AssetNameSkybox)
	if err != nil {
		panic(err)
	}

	return h.(*Handler)
}
