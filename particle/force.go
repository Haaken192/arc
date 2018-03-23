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

import "github.com/go-gl/mathgl/mgl32"

type AttractorMode uint32

const (
	AttractorNormal AttractorMode = iota
	AttractorRepell
	AttractorBlackhole
	AttractorGlobal
)

type ModuleForce struct {
	system     *System
	attractors []*Attractor

	EnableAttractors bool
}

type Attractor struct {
	Position  mgl32.Vec4
	Direction mgl32.Vec4
	Mode      AttractorMode
	Force     float32
	Range     float32
	Unused    float32
}

func (m *ModuleForce) AddAttractor(a *Attractor) {
	m.attractors = append(m.attractors, a)
}

func NewModuleForce() *ModuleForce {
	m := &ModuleForce{}

	return m
}

func NewAttractor() *Attractor {
	return &Attractor{}
}
