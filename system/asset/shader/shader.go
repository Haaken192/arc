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

package shader

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/haakenlabs/arc/core"
	"github.com/haakenlabs/arc/graphics"
	"github.com/haakenlabs/arc/system/asset"
)

const (
	AssetNameShader = "shader"
)

var _ core.AssetHandler = &Handler{}

type Handler struct {
	core.BaseAssetHandler
}

type Metadata struct {
	Name     string   `json:"name"`
	Deferred bool     `json:"deferred"`
	Files    []string `json:"files"`
}

// Load will load data from the reader.
func (h *Handler) Load(r *core.Resource) error {
	m := &Metadata{}

	data, err := ioutil.ReadAll(r.Reader())
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, m); err != nil {
		return err
	}
	name := m.Name
	if _, dup := h.Items[name]; dup {
		return core.ErrAssetExists(name)
	}

	s := graphics.NewShader(m.Deferred)

	s.SetName(m.Name)

	// Populate shader data.
	for i := range m.Files {
		r, err := core.NewResource(filepath.Join(r.DirPrefix(), m.Files[i]))
		if err != nil {
			return err
		}
		if err := asset.ReadResource(r); err != nil {
			return err
		}

		s.AddData(r.Bytes())
	}

	return h.Add(name, s)
}

func (h *Handler) Add(name string, shader *graphics.Shader) error {
	if _, dup := h.Items[name]; dup {
		return core.ErrAssetExists(name)
	}

	if err := shader.Alloc(); err != nil {
		return err
	}

	h.Items[name] = shader.ID()

	return nil
}

// Get gets an asset by name.
func (h *Handler) Get(name string) (*graphics.Shader, error) {
	a, err := h.GetAsset(name)
	if err != nil {
		return nil, err
	}

	a2, ok := a.(*graphics.Shader)
	if !ok {
		return nil, core.ErrAssetType(name)
	}

	return a2, nil
}

// MustGet is like GetAsset, but panics if an error occurs.
func (h *Handler) MustGet(name string) *graphics.Shader {
	a, err := h.Get(name)
	if err != nil {
		panic(err)
	}

	return a
}

func (h *Handler) Name() string {
	return AssetNameShader
}

func NewHandler() *Handler {
	h := &Handler{}
	h.Items = make(map[string]int32)
	h.Mu = &sync.RWMutex{}

	return h
}

func NewShaderUtilsCopy() *graphics.Shader {
	return MustGet("utils/copy")
}

func NewShaderUtilsSkybox() *graphics.Shader {
	return MustGet("utils/skybox")
}

func DefaultShader() *graphics.Shader {
	return MustGet("standard")
}

func Get(name string) (*graphics.Shader, error) {
	return mustHandler().Get(name)
}

func MustGet(name string) *graphics.Shader {
	return mustHandler().MustGet(name)
}

func mustHandler() *Handler {
	h, err := asset.GetHandler(AssetNameShader)
	if err != nil {
		panic(err)
	}

	return h.(*Handler)
}
