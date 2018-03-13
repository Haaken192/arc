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

// Node represents an object at vertex.
type Node interface {
	ID() int32
	Active() bool
}

// Descriptor is a unique identifier for vertex in a graph.
type Descriptor int32

// Vertex describes a vertex on the graph.
type Vertex struct {
	edges      []Descriptor
	data       Node
	descriptor Descriptor
	parent     *Vertex
}

// NewVertex creates a new Vertex using the provided descriptor and data.
func NewVertex(descriptor Descriptor, data Node) *Vertex {
	return &Vertex{
		descriptor: descriptor,
		data:       data,
	}
}

func (v *Vertex) AddEdge(d Descriptor) error {
	if d < 0 {
		return ErrDescriptorInvalid(d)
	}
	if d == v.descriptor {
		return ErrDescriptorInvalid(d)
	}
	if v.HasEdge(d) {
		return ErrEdgeExists{p: v.descriptor, d: d}
	}

	v.edges = append(v.edges, d)

	return nil
}

// RemoveEdge will remove the provided descriptor from this vertex's edge list.
func (v *Vertex) RemoveEdge(d Descriptor) error {
	if d < 0 {
		return ErrDescriptorInvalid(d)
	}

	for i, x := range v.edges {
		if x == d {
			v.edges[i] = v.edges[len(v.edges)-1]
			v.edges = v.edges[:len(v.edges)-1]
			return nil
		}
	}

	return ErrDescriptorNotFound(d)
}

// HasEdge returns true if this vertex forms an out edge with the provided
// descriptor.
func (v *Vertex) HasEdge(d Descriptor) bool {
	for i := range v.edges {
		if v.edges[i] == d {
			return true
		}
	}

	return false
}
