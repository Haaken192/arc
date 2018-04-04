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

package particle

import (
	"github.com/go-gl/gl/v4.3-core/gl"

	"github.com/haakenlabs/arc/graphics"
	"github.com/haakenlabs/arc/scene"
	"github.com/haakenlabs/arc/system/asset/shader"
	"github.com/haakenlabs/arc/system/asset/texture"
)

type ModuleRenderer struct {
	system       *System
	renderShader *graphics.Shader
	sprite       *graphics.Texture2D
}

func (m *ModuleRenderer) Draw(camera *scene.Camera) {
	m.system.Simulate()

	m.renderShader.Bind()
	m.system.Core.particleBuffer.Bind()

	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE)

	m.renderShader.SetUniform("v_model_matrix", m.system.GetTransform().ActiveMatrix())
	m.renderShader.SetUniform("v_view_matrix", camera.ViewMatrix())
	m.renderShader.SetUniform("v_projection_matrix", camera.ProjectionMatrix())
	m.renderShader.SetUniform("v_offset", m.system.inOffset)

	m.sprite.ActivateTexture(gl.TEXTURE0)
	gl.DrawArrays(gl.POINTS, 0, int32(m.system.Core.alive))

	m.system.Core.particleBuffer.Unbind()
	m.renderShader.Unbind()

	gl.Disable(gl.BLEND)
	gl.Enable(gl.DEPTH_TEST)
}

func NewModuleRenderer(system *System) *ModuleRenderer {
	m := &ModuleRenderer{
		system: system,
	}

	m.renderShader = shader.MustGet("particle/render")
	m.sprite = texture.MustGet("particle.png")

	return m
}
