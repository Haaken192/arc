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
	"github.com/haakenlabs/arc/core"
	"github.com/haakenlabs/arc/graphics"
	"github.com/haakenlabs/arc/system/asset/shader"
)

type EnvLightingSource int

const (
	EnvLightingSkybox EnvLightingSource = iota
	EnvLightingColor
)

type EnvironmentLighting struct {
	Source    EnvLightingSource
	Intensity float32
	Ambient   core.Color
}

type Environment struct {
	DeferredShader *graphics.Shader
	Skybox         *Skybox
	SunSource      *Light
}

func NewEnvironment() *Environment {
	e := &Environment{}

	e.DeferredShader = shader.DefaultShader()
	e.Skybox = DefaultSkybox()

	return e
}

func DefaultSkybox() *Skybox {
	//return GetAsset().MustGet(AssetNameSkybox, "default").(*Skybox)
	return nil
}
