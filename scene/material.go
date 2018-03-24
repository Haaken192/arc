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

package scene

import (
	"github.com/go-gl/gl/v4.5-core/gl"

	"github.com/haakenlabs/arc/core"
	"github.com/haakenlabs/arc/graphics"
	"github.com/haakenlabs/arc/system/asset/shader"
	"github.com/haakenlabs/arc/system/instance"
)

type MaterialTexture uint32

const (
	MaterialTextureAttachment0 MaterialTexture = iota
	MaterialTextureAttachment1
	MaterialTextureDepth
	MaterialTextureEnvironment
	MaterialTextureIrradiance
	MaterialTextureAlbedo
	MaterialTextureNormal
	MaterialTextureMetallic
)

const MaterialMaxTextures = 16

type Material struct {
	core.BaseObject

	textures         [MaterialMaxTextures]graphics.Texture
	shaderProperties map[string]interface{}
	shader           *graphics.Shader
}

func (m *Material) SetTexture(id MaterialTexture, texture graphics.Texture) {
	if id < MaterialMaxTextures {
		m.textures[id] = texture
	}
}

func (m *Material) SetShader(shader *graphics.Shader) {
	m.shader = shader
}

func (m *Material) Texture(id MaterialTexture) graphics.Texture {
	if id >= MaterialMaxTextures {
		return nil
	}

	return m.textures[id]
}

func (m *Material) Shader() *graphics.Shader {
	return m.shader
}

func (m *Material) Bind() {
	if m.shader == nil {
		return
	}

	m.shader.Bind()

	for i := range m.textures {
		if m.textures[i] != nil {
			m.textures[i].ActivateTexture(gl.TEXTURE0 + uint32(i))
		}
	}
	for key, value := range m.shaderProperties {
		m.shader.SetUniform(key, value)
	}
}

func (m *Material) Unbind() {
	m.shader.Unbind()
}

func (m *Material) SupportsDeferredPath() bool {
	if m.shader != nil {
		return m.shader.DeferredCapable()
	}

	return false
}

func (m *Material) SetProperty(property string, value interface{}) {
	m.shaderProperties[property] = value
}

func NewMaterial() *Material {
	m := &Material{
		shaderProperties: make(map[string]interface{}),
	}

	m.SetName("Material")
	instance.MustAssign(m)

	return m
}

func NewMaterialPBR() *Material {
	m := NewMaterial()

	m.shader = shader.DefaultShader()

	m.SetProperty("f_albedo", core.ColorCopper.Vec3())
	m.SetProperty("f_metallic", 1.0)
	m.SetProperty("f_roughness", 0.8)

	return m
}
