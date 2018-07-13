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

package mesh

import (
	"sync"

	"github.com/haakenlabs/arc/core"
	"github.com/haakenlabs/arc/graphics"
	"github.com/haakenlabs/arc/system/asset"
)

const (
	AssetNameMesh = "mesh" // Identifier is the type name of this asset.
)

// Mesh errors

var _ core.AssetHandler = &Handler{}

type Handler struct {
	core.BaseAssetHandler
}

// Load will load data from the reader.
func (h *Handler) Load(r *core.Resource) error {

}

func (h *Handler) Add(name string, mesh *graphics.Mesh) error {
	h.Mu.Lock()
	defer h.Mu.Unlock()

	if _, dup := h.Items[name]; dup {
		return core.ErrAssetExists(name)
	}

	if err := mesh.Alloc(); err != nil {
		return err
	}

	h.Items[name] = mesh.ID()

	return nil
}

// Get gets an asset by name.
func (h *Handler) Get(name string) (*graphics.Mesh, error) {
	h.Mu.RLock()
	defer h.Mu.RUnlock()

	a, err := h.GetAsset(name)
	if err != nil {
		return nil, err
	}

	a2, ok := a.(*graphics.Mesh)
	if !ok {
		return nil, core.ErrAssetType(name)
	}

	return a2, nil
}

// MustGet is like GetAsset, but panics if an error occurs.
func (h *Handler) MustGet(name string) *graphics.Mesh {
	a, err := h.Get(name)
	if err != nil {
		panic(err)
	}

	return a
}

func (h *Handler) Name() string {
	return AssetNameMesh
}

func NewHandler() *Handler {
	h := &Handler{}
	h.Items = make(map[string]int32)
	h.Mu = &sync.RWMutex{}

	return h
}

func Get(name string) (*graphics.Mesh, error) {
	return mustHandler().Get(name)
}

func MustGet(name string) *graphics.Mesh {
	return mustHandler().MustGet(name)
}

func mustHandler() *Handler {
	h, err := asset.GetHandler(AssetNameMesh)
	if err != nil {
		panic(err)
	}

	return h.(*Handler)
}
