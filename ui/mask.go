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
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/haakenlabs/arc/graphics"
	"github.com/haakenlabs/arc/scene"
	"github.com/haakenlabs/arc/system/window"
)

type Mask struct {
	BaseComponent

	mesh   *Mesh
	shader *graphics.Shader
	maskID uint8
}

func (m *Mask) Refresh() {
	r := m.RectTransform().Rect()

	verts := MakeQuad(r.SizeElem())

	m.mesh.Upload(verts)
}

func (m *Mask) SetMaskID(maskID uint8) {
	m.maskID = maskID
}

func (m *Mask) MaskID() uint8 {
	return m.maskID
}

func (m *Mask) WriteMask() {
	var parentMask uint8

	if p := m.GameObject().Parent(); p != nil {
		if pm := MaskComponent(p); pm != nil {
			parentMask = pm.MaskID()
		}
	}

	m.shader.Bind()
	m.mesh.Bind()

	gl.StencilMask(0xFF)
	gl.StencilFunc(gl.EQUAL, int32(parentMask), 0xFF)
	gl.StencilOp(gl.KEEP, gl.INCR, gl.INCR)

	gl.ColorMask(false, false, false, false)

	m.shader.SetUniform("v_ortho_matrix", window.OrthoMatrix())
	m.shader.SetUniform("v_model_matrix", m.RectTransform().Rect().Matrix())

	m.mesh.Draw()

	m.mesh.Unbind()
	m.shader.Unbind()

	gl.ColorMask(true, true, true, true)
}

func MaskComponent(g *scene.GameObject) *Mask {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Mask); ok {
			return ct
		}
	}

	return nil
}
