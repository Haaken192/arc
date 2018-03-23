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
	"encoding/binary"
	"unsafe"

	"github.com/go-gl/gl/v4.3-core/gl"

	"github.com/haakenlabs/arc/graphics"
	"github.com/haakenlabs/forge"
)

type ModuleCore struct {
	StartColor    forge.Color
	Duration      float32
	StartDelay    float32
	StartLifetime float32
	StartSpeed    float32
	StartSize     float32
	PlaybackSpeed float32
	RandomSeed    uint32
	Looping       bool

	alive           uint32
	dead            uint32
	emit            uint32
	maxParticles    uint32
	particleBuffer  *buffer
	lifecycleShader *graphics.Shader
	simulateShader  *graphics.Shader
}

const (
	sizeOfParticle = 80
)

func (m *ModuleCore) SetMaxParticles(value uint32) {
	m.maxParticles = value
	m.dead = m.maxParticles
	m.alive = 0

	m.particleBuffer.SetSize(m.maxParticles)
}

func (m *ModuleCore) syncCounts() {
	m.particleBuffer.Bind()

	gl.BindBuffer(gl.SHADER_STORAGE_BUFFER, m.particleBuffer.idIndex)
	p := gl.MapBuffer(gl.SHADER_STORAGE_BUFFER, gl.READ_ONLY)

	out := make([]byte, 12)

	for i := range out {
		out[i] = *((*byte)(unsafe.Pointer(uintptr(p) + uintptr(i))))
	}

	m.alive = binary.LittleEndian.Uint32(out[0:])
	m.dead = binary.LittleEndian.Uint32(out[4:])
	m.emit = binary.LittleEndian.Uint32(out[8:])

	gl.UnmapBuffer(gl.SHADER_STORAGE_BUFFER)
	m.particleBuffer.Unbind()
}

func (m *ModuleCore) MaxParticles() uint32 {
	return m.maxParticles
}

func (m *ModuleCore) ParticleCount() uint32 {
	return m.alive
}

func NewModuleCore(maxParticles uint32) *ModuleCore {
	m := &ModuleCore{
		StartColor:    forge.ColorWhite,
		StartSpeed:    10.0,
		StartSize:     1.0,
		PlaybackSpeed: 1.0,
		maxParticles:  maxParticles,
		Looping:       true,
	}

	m.lifecycleShader = forge.GetAsset().MustGet(forge.AssetNameShader, "particle/lifecycle").(*forge.Shader)
	m.simulateShader = forge.GetAsset().MustGet(forge.AssetNameShader, "particle/simulate").(*forge.Shader)

	m.particleBuffer = newBuffer(m.maxParticles)
	m.particleBuffer.Alloc()

	m.SetMaxParticles(m.maxParticles)

	return m
}
