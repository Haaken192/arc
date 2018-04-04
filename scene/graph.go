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
	"github.com/haakenlabs/arc/internal/sg"
	"github.com/haakenlabs/arc/system/instance"
)

type GraphListener interface {
	// OnSceneGraphUpdate is called when the SceneGraph has been updated.
	OnSceneGraphUpdate()
}

type Graph struct {
	root   *GameObject
	graph  *sg.Graph
	aCache []*GameObject
	cCache []Component
	scene  *Scene
	dirty  bool
}

func NewGraph(scene *Scene) *Graph {
	s := &Graph{
		graph:  sg.NewGraph(),
		scene:  scene,
		dirty:  true,
		aCache: []*GameObject{},
		cCache: []Component{},
	}

	s.root = NewGameObject("__root_node__")
	s.graph.AddVertex(s.root)
	s.Update()

	return s
}

func (s *Graph) Dirty() bool {
	return s.dirty
}

func (s *Graph) SetDirty() {
	s.dirty = true
}

func (s *Graph) Update() {
	s.aCache = s.aCache[:0]
	s.cCache = s.cCache[:0]

	for _, v := range s.graph.DFS(0, false) {
		n, err := s.graph.NodeAtVertex(v)
		if err != nil {
			continue
		}
		s.aCache = append(s.aCache, n.(*GameObject))
		s.cCache = append(s.cCache, n.(*GameObject).Components()...)
	}

	s.dirty = false

	s.scene.OnSceneGraphUpdate()
	s.SendMessage(MessageSGUpdate)
}

func (s *Graph) AddObject(object, parent *GameObject) error {
	var err error
	d := sg.Descriptor(-1)
	p := sg.Descriptor(-1)

	if parent != nil {
		p, err = s.graph.DescriptorByNode(parent)
	} else {
		p, err = s.graph.DescriptorByNode(s.root)
		parent = s.root
	}
	if err != nil {
		return err
	}

	if d, err = s.graph.AddVertex(object); err != nil {
		return err
	}

	if err := s.graph.AddEdge(p, d); err != nil {
		return err
	}

	for _, v := range object.children {
		if err := s.AddObject(v, object); err != nil {
			return err
		}
	}

	object.parent = parent
	object.parent.AddChild(object)

	object.scene = s.scene

	s.Update()

	return nil
}

func (s *Graph) RemoveObject(object *GameObject) error {
	var r []int32

	descendants := s.Descendants(object, true)
	for i := len(descendants) - 1; i >= 0; i-- {
		r = append(r, descendants[i].ID())
	}
	r = append(r, object.ID())

	d, err := s.graph.DescriptorByNode(object)
	if err != nil {
		return err
	}

	if err := s.graph.DeleteVertex(d); err != nil {
		return err
	}

	instance.Release(r...)

	s.Update()

	return nil
}

func (s *Graph) MoveObject(object, parent *GameObject) error {
	d, err := s.graph.DescriptorByNode(object)
	if err != nil {
		return err
	}
	p, err := s.graph.DescriptorByNode(parent)
	if err != nil {
		return err
	}

	if err := s.graph.MoveVertex(d, p); err != nil {
		return err
	}

	oldParent := object.parent
	oldParent.RemoveChild(object.ID())

	object.parent = parent
	object.parentChanged()

	s.Update()

	return nil
}

func (s *Graph) Descendants(object *GameObject, disable bool) []*GameObject {
	var descendants []*GameObject

	d, err := s.graph.DescriptorByNode(object)
	if err != nil {
		return descendants
	}

	dfs := s.graph.DFS(d, disable)

	for _, v := range dfs {
		o, err := s.graph.NodeAtVertex(v)
		if err != nil {
			continue
		}
		if o.ID() == object.ID() {
			continue
		}
		descendants = append(descendants, o.(*GameObject))
	}

	return descendants
}

func (s *Graph) Objects() []*GameObject {
	return s.aCache
}

func (s *Graph) Components() []Component {
	return s.cCache
}

func (s *Graph) SendMessage(message Message) {
	for _, v := range s.aCache {
		v.SendMessage(message)
	}
}
