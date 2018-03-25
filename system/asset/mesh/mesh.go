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
	"encoding/gob"
	"sync"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/juju/errors"

	"github.com/haakenlabs/arc/core"
	"github.com/haakenlabs/arc/graphics"
	"github.com/haakenlabs/arc/pkg/math"
	"github.com/haakenlabs/arc/system/asset"
)

const (
	AssetNameMesh = "mesh" // Identifier is the type name of this asset.
)

// Mesh errors
var (
	ErrMeshInvalidFaceType = errors.New("invalid model face type")
	ErrMeshMissingFaces    = errors.New("model has no faces")
)

const (
	FaceVertex = iota
	FaceTexture
	FaceNormal
)

type FaceType int

const (
	FaceTypeV FaceType = iota
	FaceTypeVT
	FaceTypeVN
	FaceTypeVTN
)

type Face [3]math.IVec3

type Metadata struct {
	Name  string       `json:"name"`
	FType FaceType     `json:"face_type"`
	V     []mgl32.Vec3 `json:"v"`
	N     []mgl32.Vec3 `json:"n"`
	T     []mgl32.Vec2 `json:"t"`
	F     []Face       `json:"f"`
}

var _ core.AssetHandler = &Handler{}

type Handler struct {
	core.BaseAssetHandler
}

// Load will load data from the reader.
func (h *Handler) Load(r *core.Resource) error {
	metadata := &Metadata{}
	m := graphics.NewMesh()

	dec := gob.NewDecoder(r.Reader())
	err := dec.Decode(&metadata)
	if err != nil {
		return err
	}

	name := metadata.Name

	if _, dup := h.Items[name]; dup {
		return core.ErrAssetExists(name)
	}

	if len(metadata.F) == 0 {
		return ErrMeshMissingFaces
	}

	v := make([]mgl32.Vec3, len(metadata.F)*3)
	n := make([]mgl32.Vec3, len(metadata.F)*3)
	t := make([]mgl32.Vec2, len(metadata.F)*3)

	for i := range metadata.F {
		for j := range metadata.F[i] {
			switch metadata.FType {
			case FaceTypeV:
				v[i*3+j] = metadata.V[metadata.F[i][j][FaceVertex]]
			case FaceTypeVT:
				v[i*3+j] = metadata.V[metadata.F[i][j][FaceVertex]]
				t[i*3+j] = metadata.T[metadata.F[i][j][FaceTexture]]
			case FaceTypeVN:
				v[i*3+j] = metadata.V[metadata.F[i][j][FaceVertex]]
				n[i*3+j] = metadata.N[metadata.F[i][j][FaceNormal]]
			case FaceTypeVTN:
				v[i*3+j] = metadata.V[metadata.F[i][j][FaceVertex]]
				t[i*3+j] = metadata.T[metadata.F[i][j][FaceTexture]]
				n[i*3+j] = metadata.N[metadata.F[i][j][FaceNormal]]
			default:
				return ErrMeshInvalidFaceType
			}
		}
	}

	m.SetVertices(v)
	m.SetNormals(n)
	m.SetUvs(t)

	return h.Add(name, m)
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
