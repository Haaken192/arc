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

package sg

import (
	"fmt"
	"math"
	"sync"
)

// Graph describes a directed acyclic graph, used to represent a scene graph.
type Graph struct {
	vertices map[Descriptor]*Vertex
	index    Descriptor
	mu       *sync.RWMutex
}

// NewGraph creates a new, empty graph.
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[Descriptor]*Vertex),
		mu:       &sync.RWMutex{},
	}
}

// AddVertex adds a vertex with the provided Node. The descriptor of the
// vertex will be added automatically.
func (g *Graph) AddVertex(n Node) (Descriptor, error) {
	var v *Vertex
	var err error

	d, err := g.nextDescriptor()
	if err != nil {
		return -1, err
	}

	v = &Vertex{
		descriptor: d,
		data:       n,
	}

	g.vertices[d] = v

	return v.descriptor, nil
}

// MoveVertex moves a vertex so its parent is the descriptor in the second argument.
func (g *Graph) MoveVertex(d, p Descriptor) error {
	if err := g.ValidateDescriptor(d); err != nil {
		return err
	}
	if err := g.ValidateDescriptor(p); err != nil {
		return err
	}

	if g.DescendantOf(p, d) {
		return ErrDescriptorDescendant{d: p, p: d}
	}

	oldParent, err := g.Parent(d)
	if err != nil {
		return err
	}

	if err := g.vertices[oldParent].RemoveEdge(d); err != nil {
		return err
	}

	if err := g.vertices[p].AddEdge(d); err != nil {
		return err
	}

	g.vertices[d].parent = g.vertices[p]

	return nil
}

// DeleteVertex removes vertex from the graph.
func (g *Graph) DeleteVertex(d Descriptor) error {
	if !g.HasVertexWithDescriptor(d) {
		return ErrDescriptorNotFound(d)
	}

	if p, err := g.Parent(d); err == nil {
		g.vertices[p].RemoveEdge(d)
	}

	for _, v := range g.DFS(d, true) {
		delete(g.vertices, v)
	}

	delete(g.vertices, d)

	return nil
}

// AddEdge adds edge from p to d.
func (g *Graph) AddEdge(p, d Descriptor) error {
	if err := g.ValidateDescriptor(d); err != nil {
		return err
	}
	if err := g.ValidateDescriptor(p); err != nil {
		return err
	}
	if g.DescendantOf(p, d) {
		return ErrDescriptorDescendant{d: d, p: p}
	}

	g.vertices[d].parent = g.vertices[p]

	if err := g.vertices[p].AddEdge(d); err != nil {
		return err
	}

	return nil
}

// EdgeExists returns true if an edge from p to d exists.
func (g *Graph) EdgeExists(p, d Descriptor) bool {
	if err := g.ValidateDescriptor(d); err != nil {
		return false
	}
	if err := g.ValidateDescriptor(p); err != nil {
		return false
	}

	return g.vertices[p].HasEdge(d)
}

// ValidateDescriptor checks the validity of the descriptor.
func (g *Graph) ValidateDescriptor(d Descriptor) error {
	if d < 0 {
		return ErrDescriptorInvalid(d)
	}
	if !g.HasVertexWithDescriptor(d) {
		return ErrDescriptorNotFound(d)
	}

	return nil
}

// HasVertexWithDescriptor returns true if the graph has a vertex
// with the specified Descriptor.
func (g *Graph) HasVertexWithDescriptor(d Descriptor) bool {
	_, ok := g.vertices[d]

	return ok
}

// HasVertexWithID returns true if the graph has a vertex with
// the specified ID.
func (g *Graph) HasVertexWithID(d int32) bool {
	for i := range g.vertices {
		if g.vertices[i].data.ID() == d {
			return true
		}
	}

	return false
}

func (g *Graph) DescriptorByNode(n Node) (Descriptor, error) {
	for k, v := range g.vertices {
		if v.data.ID() == n.ID() {
			return k, nil
		}
	}

	return -1, fmt.Errorf("descriptor not found for object with ID: %d", n.ID())
}

// ParentOf returns true if p is a parent of d.
func (g *Graph) ParentOf(p, d Descriptor) bool {
	if p < 0 || d < 0 {
		return false
	}
	if !g.HasVertexWithDescriptor(p) || !g.HasVertexWithDescriptor(d) {
		return false
	}

	return g.vertices[p].HasEdge(d)
}

// DescendantOf returns true if d is a descendant of p.
func (g *Graph) DescendantOf(d, p Descriptor) bool {
	for _, v := range g.DFS(p, true) {
		if v == d {
			return true
		}
	}

	return false
}

// Parent gets the parent descriptor of this descriptor.
func (g *Graph) Parent(d Descriptor) (Descriptor, error) {
	if d < 0 {
		return -1, ErrDescriptorInvalid(d)
	}
	if !g.HasVertexWithDescriptor(d) {
		return -1, ErrDescriptorNotFound(d)
	}

	if g.vertices[d].parent == nil {
		return -1, ErrDescriptorNoParent(d)
	}

	return g.vertices[d].parent.descriptor, nil
}

// NodeAtVertex returns the data at the provided descriptor.
func (g *Graph) NodeAtVertex(d Descriptor) (Node, error) {
	if d < 0 {
		return nil, ErrDescriptorInvalid(d)
	}
	if !g.HasVertexWithDescriptor(d) {
		return nil, ErrDescriptorNotFound(d)
	}

	return g.vertices[d].data, nil
}

// VertexActive will return true if this vertex is active. All vertices above
// this vertex will be checked and if any are inactive, this vertex will also
// be considered inactive.
func (g *Graph) VertexActive(d Descriptor) bool {
	if err := g.ValidateDescriptor(d); err != nil {
		return false
	}

	p := g.vertices[d]
	if !p.data.Active() {
		return false
	}

	for p.parent != nil {
		if !p.parent.data.Active() {
			return false
		}
		p = p.parent
	}
	return true
}

// DFS performs a depth-first search of the graph from the provided descriptor.
// If disabled vertices should be included, disabled should be set to true.
func (g *Graph) DFS(d Descriptor, disabled bool) []Descriptor {
	var vertices []Descriptor

	if err := g.ValidateDescriptor(d); err != nil {
		return vertices
	}
	if !disabled {
		if !g.VertexActive(d) {
			return vertices
		}
	}

	vertices = []Descriptor{d}

	for i := range g.vertices[d].edges {
		l := g.DFS(g.vertices[d].edges[i], disabled)
		vertices = append(vertices, l...)
	}

	return vertices
}

// nextDescriptor will return the next available descriptor and
// assign it to the vertices map. The descriptor will not be
// released until it is explicitly released by DeleteVertex.
func (g *Graph) nextDescriptor() (Descriptor, error) {
	descriptor := g.index

	if len(g.vertices) >= math.MaxInt32 {
		return -1, ErrDescriptorLimit
	}

	id := g.index + 1
	_, ok := g.vertices[id]

	for ok {
		id = g.index + 1
		_, ok = g.vertices[id]
	}
	g.index = id

	return descriptor, nil
}
