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
	"github.com/haakenlabs/arc/internal/sg"
	"github.com/haakenlabs/arc/system/instance"
)

type Message uint8

const (
	MessageActivate Message = iota
	MessageStart
	MessageAwake
	MessageUpdate
	MessageLateUpdate
	MessageFixedUpdate
	MessageGUIRender
	MessageSGUpdate
)

var _ sg.Node = &GameObject{}

type GameObject struct {
	core.BaseObject

	components []Component
	children   []*GameObject
	parent     *GameObject
	scene      *Scene
	active     bool
}

func (g *GameObject) Active() bool {
	return g.active
}

func (g *GameObject) SetActive(active bool) {
	if g.active != active {
		g.active = active

		//if g.graph != nil {
		//	g.graph.SetDirty()
		//}
	}
}

func (g *GameObject) Scene() *Scene {
	return g.scene
}

func (g *GameObject) Transform() Transform {
	return g.components[0].(Transform)
}

func (g *GameObject) SetTransform(transform Transform) {
	if transform == nil {
		return
	}
	if g.Transform().ID() == transform.ID() {
		return
	}

	g.components[0] = transform
	g.components[0].SetGameObject(g)
	g.components[0].OnParentChanged()
}

// AddChild will add a child object to this object.
// Note: This will not modify the scene graph in anyway. It is recommended to
// use the AddObject or MoveObject functions in SceneGraph instead.
func (g *GameObject) AddChild(object *GameObject) {
	if object == nil {
		return
	}
	for _, v := range g.children {
		if v.ID() == object.ID() {
			return
		}
	}

	g.children = append(g.children, object)
}

// RemoveChild will remove a child object with matching ID from this object.
// Note: This will not modify the scene graph in anyway. It is recommended to
// use the RemoveObject function in SceneGraph instead.
func (g *GameObject) RemoveChild(id int32) {
	for i, v := range g.children {
		if v.ID() == id {
			g.children[i] = g.children[len(g.children)-1]
			g.children = g.children[:len(g.children)-1]
		}
	}
}

// AddComponent attaches a component to this object.
func (g *GameObject) AddComponent(component Component) {
	if component == nil {
		return
	}
	for _, v := range g.components {
		if v.ID() == component.ID() {
			return
		}
	}

	g.components = append(g.components, component)
	component.SetGameObject(g)
	component.OnParentChanged()
}

// AddComponent removes a component from this object.
func (g *GameObject) RemoveComponent(id int32) {
	for i, v := range g.components {
		if v.ID() == id {
			g.components[i] = g.components[len(g.components)-1]
			g.components = g.components[:len(g.components)-1]
			v.SetGameObject(nil)
		}
	}
}

// Parent returns the parent of this object.
func (g *GameObject) Parent() *GameObject {
	return g.parent
}

// Components returns the components of this object.
func (g *GameObject) Components() []Component {
	return g.components
}

// Ancestors lists all ancestor objects of this game object.
func (g *GameObject) Ancestors() []*GameObject {
	var ancestors []*GameObject

	if g.Parent() != nil {
		ancestors = append(ancestors, g.Parent())
		ancestors = append(ancestors, g.Parent().Ancestors()...)
	}

	return ancestors
}

// SendMessage calls the function associated with the given message.
func (g *GameObject) SendMessage(msg Message) {
	if !g.active {
		return
	}

	if msg == MessageActivate {
		g.SendMessage(MessageAwake)
		return
	}

	for i := range g.components {
		switch msg {
		case MessageStart:
			if c, ok := g.components[i].(ScriptComponent); ok {
				c.Start()
			}
		case MessageAwake:
			if c, ok := g.components[i].(ScriptComponent); ok {
				c.Awake()
			}
		case MessageUpdate:
			if c, ok := g.components[i].(ScriptComponent); ok {
				c.Update()
			}
		case MessageLateUpdate:
			if c, ok := g.components[i].(ScriptComponent); ok {
				c.LateUpdate()
			}
		case MessageFixedUpdate:
			if c, ok := g.components[i].(ScriptComponent); ok {
				c.FixedUpdate()
			}
		case MessageGUIRender:
			if c, ok := g.components[i].(ScriptComponent); ok {
				c.GUIRender()
			}
		case MessageSGUpdate:
			if c, ok := g.components[i].(GraphListener); ok {
				c.OnSceneGraphUpdate()
			}
		}
	}
}

func (g *GameObject) ComponentsInChildren() []Component {
	var components []Component

	if len(g.children) == 0 {
		return components
	}

	if g.scene == nil {
		for _, v := range g.children {
			components = append(components, v.Components()...)
			components = append(components, v.ComponentsInChildren()...)
		}
		return components
	}

	for _, v := range g.scene.Descendants(g, false) {
		components = append(components, v.Components()...)
	}

	return components
}

func (g *GameObject) ComponentsInParent() []Component {
	var components []Component

	if g.parent == nil {
		return components
	}

	for _, v := range g.Ancestors() {
		components = append(components, v.Components()...)
	}

	return components
}

func (g *GameObject) Environment() *Environment {
	if g.scene != nil {
		return g.scene.Environment()
	}

	return nil
}

func (g *GameObject) parentChanged() {
	for _, v := range g.components {
		v.OnParentChanged()
	}
}

func (g *GameObject) transformChanged() {
	for _, v := range g.components {
		v.OnTransformChanged()
	}
}

func NewGameObject(name string) *GameObject {
	g := &GameObject{
		active: true,
	}

	g.SetName(name)
	instance.MustAssign(g)

	g.components = []Component{NewTransform()}
	g.components[0].SetGameObject(g)

	return g
}
