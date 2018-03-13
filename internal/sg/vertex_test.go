package sg

import (
	"fmt"
	"testing"
)

func TestVertex_RemoveEdge(t *testing.T) {
	vertex := NewVertex(0, nil)
	vertex.edges = []Descriptor{1, 2, 3}

	tests := []struct {
		in   Descriptor
		want error
	}{
		{in: 0, want: nil},
		{in: 1, want: nil},
		{in: 2, want: nil},
		{in: 3, want: nil},
		{in: -1, want: ErrDescriptorInvalid(-1)},
		{in: 4, want: ErrDescriptorNotFound(4)},
	}

	for i, v := range tests {
		got := vertex.RemoveEdge(v.in)

		if v.want != got {
			fmt.Errorf("RemoveEdge case %d failed. want: %v  got: %v", i, v.want, got)
		}
	}
}

func TestVertex_HasEdge(t *testing.T) {
	vertex := NewVertex(0, nil)
	vertex.edges = []Descriptor{1, 2, 3}

	tests := []struct {
		in   Descriptor
		want bool
	}{
		{in: 0, want: false},
		{in: 1, want: true},
		{in: 2, want: true},
		{in: 3, want: true},
		{in: -1, want: false},
		{in: 4, want: false},
	}

	for i, v := range tests {
		got := vertex.HasEdge(v.in)

		if v.want != got {
			fmt.Errorf("RemoveEdge case %d failed. want: %v  got: %v", i, v.want, got)
		}
	}
}

func TestVertex_AddEdge(t *testing.T) {
	vertex := NewVertex(0, nil)
	vertex.edges = []Descriptor{1, 2, 3}

	tests := []struct {
		in   Descriptor
		want error
	}{
		{in: 0, want: ErrDescriptorInvalid(0)},
		{in: 1, want: ErrEdgeExists{0, 1}},
		{in: 2, want: ErrEdgeExists{0, 2}},
		{in: 3, want: ErrEdgeExists{0, 3}},
		{in: -1, want: ErrDescriptorInvalid(-1)},
		{in: 4, want: nil},
	}

	for i, v := range tests {
		got := vertex.AddEdge(v.in)

		if v.want != got {
			fmt.Errorf("RemoveEdge case %d failed. want: %v  got: %v", i, v.want, got)
		}
	}
}
