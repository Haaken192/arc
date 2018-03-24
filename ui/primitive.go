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
	"github.com/go-gl/mathgl/mgl32"
	"github.com/haakenlabs/forge"
)

type Primitive interface {
	Rect() forge.Rect
	Draw(mgl32.Mat4)
	Refresh()
	Position() mgl32.Vec2
	Size() mgl32.Vec2
	SetRect(forge.Rect)
	SetSize(mgl32.Vec2)
	SetPosition(mgl32.Vec2)
}

var _ Primitive = &BasePrimitive{}

type BasePrimitive struct {
	rect     forge.Rect
	material *forge.Material
	mesh     *Mesh
}

func (p *BasePrimitive) Rect() forge.Rect {
	return p.rect
}

func (p *BasePrimitive) SetRect(rect forge.Rect) {
	p.rect = rect
	p.Refresh()
}

func (p *BasePrimitive) SetSize(size mgl32.Vec2) {
	p.rect.SetSize(size)
}

func (p *BasePrimitive) SetPosition(position mgl32.Vec2) {
	p.rect.SetOrigin(position)
}

func (p *BasePrimitive) Position() mgl32.Vec2 {
	return p.rect.Origin()
}

func (p *BasePrimitive) Size() mgl32.Vec2 {
	return p.rect.Size()
}

func (p *BasePrimitive) Refresh() {}

func (p *BasePrimitive) Draw(mgl32.Mat4) {}

func (p *BasePrimitive) SetMaterial(material *forge.Material) {
	p.material = material
}

func (p *BasePrimitive) SetMesh(mesh *Mesh) {
	p.mesh = mesh
}

func (p *BasePrimitive) Mesh() *Mesh {
	return p.mesh
}

func (p *BasePrimitive) Shader() *forge.Material {
	return p.material
}
