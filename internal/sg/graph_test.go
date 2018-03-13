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

import "testing"

type dummyNode struct {
	id     int32
	active bool
}

func (n *dummyNode) ID() int32 { return n.id }

func (n *dummyNode) Active() bool { return n.active }

func setupTestGraph() *Graph {
	g := NewGraph()

	g.AddVertex(&dummyNode{id: 0, active: true})
	g.AddVertex(&dummyNode{id: 1, active: true})
	g.AddVertex(&dummyNode{id: 2, active: false})
	g.AddVertex(&dummyNode{id: 3, active: true})
	g.AddVertex(&dummyNode{id: 4, active: false})
	g.AddVertex(&dummyNode{id: 5, active: true})
	g.AddVertex(&dummyNode{id: 6, active: true})
	g.AddVertex(&dummyNode{id: 7, active: true})
	g.AddVertex(&dummyNode{id: 8, active: false})
	g.AddVertex(&dummyNode{id: 9, active: true})
	g.AddVertex(&dummyNode{id: 10, active: true})
	g.AddVertex(&dummyNode{id: 11, active: true})
	g.AddVertex(&dummyNode{id: 12, active: true})
	g.AddVertex(&dummyNode{id: 13, active: true})

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)
	g.AddEdge(3, 9)
	g.AddEdge(4, 10)
	g.AddEdge(5, 11)
	g.AddEdge(5, 12)
	g.AddEdge(6, 13)

	return g
}

func TestGraph_HasVertexWithDescriptor(t *testing.T) {
	g := setupTestGraph()

	tests := []struct {
		in   Descriptor
		want bool
	}{
		{in: 0, want: true},
		{in: 1, want: true},
		{in: 2, want: true},
		{in: -1, want: false},
		{in: 20, want: false},
	}

	for i, v := range tests {
		got := g.HasVertexWithDescriptor(v.in)
		if v.want != got {
			t.Errorf(
				"HasVertexWithID case %d failed. want: %v got: %v",
				i, v.want, got)
		}
	}
}

func TestGraph_HasVertexWithID(t *testing.T) {
	g := setupTestGraph()

	tests := []struct {
		in   int32
		want bool
	}{
		{in: 0, want: true},
		{in: 1, want: true},
		{in: 2, want: true},
		{in: -1, want: false},
		{in: 20, want: false},
	}

	for i, v := range tests {
		got := g.HasVertexWithID(v.in)
		if v.want != got {
			t.Errorf(
				"HasVertexWithID case %d failed. want: %v got: %v",
				i, v.want, got)
		}
	}
}

func TestGraph_ValidateDescriptor(t *testing.T) {
	g := setupTestGraph()

	tests := []struct {
		in   Descriptor
		want error
	}{
		{in: 0, want: nil},
		{in: 1, want: nil},
		{in: 2, want: nil},
		{in: -1, want: ErrDescriptorInvalid(-1)},
		{in: 20, want: ErrDescriptorNotFound(20)},
	}

	for i, v := range tests {
		got := g.ValidateDescriptor(v.in)
		if v.want != got {
			t.Errorf(
				"HasVertexWithID case %d failed. want: %v got: %v",
				i, v.want, got)
		}
	}
}

func TestGraph_AddVertex(t *testing.T) {
	g := NewGraph()

	tests := []struct {
		in   Node
		want Descriptor
	}{
		{in: &dummyNode{id: 0, active: true}, want: 0},
		{in: &dummyNode{id: 1, active: true}, want: 1},
		{in: &dummyNode{id: 2, active: true}, want: 2},
		{in: &dummyNode{id: 3, active: true}, want: 3},
		{in: &dummyNode{id: 4, active: true}, want: 4},
		{in: &dummyNode{id: 5, active: true}, want: 5},
		{in: &dummyNode{id: 6, active: true}, want: 6},
	}

	for i, v := range tests {
		got, err := g.AddVertex(v.in)

		if v.want != got {
			t.Errorf(
				"AddVertex case %d failed. want: %d got: %d error: %v",
				i, v.want, got, err)
		}
	}
}

func TestGraph_EdgeExists(t *testing.T) {
	g := setupTestGraph()

	tests := []struct {
		inD  Descriptor
		inP  Descriptor
		want bool
	}{
		{0, 1, true},
		{3, 8, true},
		{2, 6, true},
		{11, 1, false},
		{-1, 1, false},
		{-1, 100, false},
	}

	for i, v := range tests {
		got := g.EdgeExists(v.inD, v.inP)

		if v.want != got {
			t.Errorf(
				"EdgeExists case %d failed. want: %v got: %v",
				i, v.want, got)
		}
	}
}

func TestGraph_ParentOf(t *testing.T) {
	g := setupTestGraph()

	tests := []struct {
		inP  Descriptor
		inD  Descriptor
		want bool
	}{
		{0, 1, true},
		{3, 8, true},
		{2, 6, true},
		{0, 0, false},
		{0, 3, false},
		{6, 2, false},
		{11, 1, false},
		{-1, 1, false},
		{-1, 100, false},
		{100, 99, false},
	}

	for i, v := range tests {
		got := g.ParentOf(v.inP, v.inD)

		if v.want != got {
			t.Errorf(
				"ParentOf case %d failed. want: %v got: %v",
				i, v.want, got)
		}
	}
}

func TestGraph_DFS(t *testing.T) {
	g := setupTestGraph()

	tests := []struct {
		in      Descriptor
		want    []Descriptor
		disable bool
	}{
		{0, []Descriptor{0, 1, 3, 7, 8, 9, 4, 10, 2, 5, 11, 12, 6, 13}, true},
		{1, []Descriptor{1, 3, 7, 8, 9, 4, 10}, true},
		{2, []Descriptor{2, 5, 11, 12, 6, 13}, true},
		{3, []Descriptor{3, 7, 8, 9}, true},
		{4, []Descriptor{4, 10}, true},
		{5, []Descriptor{5, 11, 12}, true},
		{6, []Descriptor{6, 13}, true},
		{0, []Descriptor{0, 1, 3, 7, 9}, false},
		{1, []Descriptor{1, 3, 7, 9}, false},
		{2, []Descriptor{}, false},
		{3, []Descriptor{3, 7, 9}, false},
		{4, []Descriptor{}, false},
		{5, []Descriptor{}, false},
		{6, []Descriptor{}, false},
	}

	for i, v := range tests {
		got := g.DFS(v.in, v.disable)

		if len(v.want) != len(got) {
			t.Errorf("DFS case %d failed (size mismatch).\n\tgot: %v\n\twant: %v", i, got, v.want)
		}
		for j := range got {
			if got[j] != v.want[j] {
				t.Errorf("DFS case %d failed (value mismatch).\n\tgot: %v\n\twant: %v", i, got, v.want)
			}
		}
	}
}

func TestGraph_Movement(t *testing.T) {
	tests := []struct {
		in          Descriptor
		parent      Descriptor
		err         error
		wantInclude []Descriptor
		wantExclude []Descriptor
	}{
		{3, 2, nil, []Descriptor{2, 5, 11, 12, 6, 13, 3, 7, 8, 9}, []Descriptor{}},
		{2, 13, ErrDescriptorDescendant{13, 2}, []Descriptor{13}, []Descriptor{}},
		{5, 9, nil, []Descriptor{9, 5, 11, 12}, []Descriptor{9, 5, 11, 12}},
		{0, 100, ErrDescriptorNotFound(100), nil, nil},
		{100, 13, ErrDescriptorNotFound(100), []Descriptor{13}, []Descriptor{}},
	}

	for i, v := range tests {
		g := setupTestGraph()
		gotE := g.MoveVertex(v.in, v.parent)

		if v.err != gotE {
			t.Errorf("GraphMovement case %d failed (error mismatch).\n\tgot: %v\n\twant: %v", i, gotE, v.err)
		} else {
			gotInclude := g.DFS(v.parent, true)
			gotExclude := g.DFS(v.parent, false)

			if len(v.wantInclude) != len(gotInclude) {
				t.Errorf("GraphMovement case %d failed (include size mismatch).\n\tgot: %v\n\twant: %v", i, gotInclude, v.wantInclude)
			} else {
				for j := range gotInclude {
					if gotInclude[j] != v.wantInclude[j] {
						t.Errorf("GraphMovement case %d failed (include value mismatch).\n\tgot: %v\n\twant: %v", i, gotInclude, v.wantInclude)
					}
				}
			}

			if len(v.wantExclude) != len(gotExclude) {
				t.Errorf("GraphMovement case %d failed (exclude size mismatch).\n\tgot: %v\n\twant: %v", i, gotExclude, v.wantExclude)
			} else {
				for j := range gotExclude {
					if gotExclude[j] != v.wantExclude[j] {
						t.Errorf("GraphMovement case %d failed (exclude value mismatch).\n\tgot: %v\n\twant: %v", i, gotExclude, v.wantExclude)
					}
				}
			}
		}
	}
}

func TestGraph_DeleteVertex(t *testing.T) {
	tests := []struct {
		in   Descriptor
		want []Descriptor
		err  error
	}{
		{0, []Descriptor{}, nil},
		{1, []Descriptor{0, 2, 5, 11, 12, 6, 13}, nil},
	}

	for i, v := range tests {
		g := setupTestGraph()
		gotE := g.DeleteVertex(v.in)

		if v.err != gotE {
			t.Errorf("DeleteVertex case %d failed (error mismatch).\n\tgot: %v\n\twant: %v", i, gotE, v.err)
		} else {
			got := g.DFS(0, true)

			if len(v.want) != len(got) {
				t.Errorf("DeleteVertex case %d failed (exclude size mismatch).\n\tgot: %v\n\twant: %v", i, got, v.want)
			} else {
				for j := range got {
					if got[j] != v.want[j] {
						t.Errorf("DeleteVertex case %d failed (exclude value mismatch).\n\tgot: %v\n\twant: %v", i, got, v.want)
					}
				}
			}
		}
	}
}

func BenchmarkGraph_DFS_Disabled(b *testing.B) {
	g := setupTestGraph()

	for i := 0; i < b.N; i++ {
		g.DFS(0, true)
	}

}

func BenchmarkGraph_DFS(b *testing.B) {
	g := setupTestGraph()

	for i := 0; i < b.N; i++ {
		g.DFS(0, false)
	}
}
