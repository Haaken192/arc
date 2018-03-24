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

package ui

import (
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/haakenlabs/forge"
)

type Controller struct {
	forge.BaseScriptComponent

	wCache      []Widget
	selected    Widget
	highlighted Widget
}

func (c *Controller) UpdateCache() {
	c.wCache = c.wCache[:0]

	components := c.GameObject().ComponentsInChildren()
	for i := range components {
		if w, ok := components[i].(Widget); ok {
			c.wCache = append(c.wCache, w)
		}
	}
}

func (c *Controller) OnSceneGraphUpdate() {
	c.UpdateCache()
}

func (c *Controller) GUIRender() {
	if len(c.wCache) == 0 {
		return
	}

	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFuncSeparate(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA, gl.ONE, gl.ONE_MINUS_SRC_ALPHA)

	for _, v := range c.wCache {
		v.Redraw()
	}

	gl.Disable(gl.BLEND)
	gl.Enable(gl.DEPTH_TEST)
}

func (c *Controller) Resize() {
	if c.GameObject() != nil {
		RectTransformComponent(c.GameObject()).SetSize(forge.GetWindow().Resolution().Vec2())
	}
}

func (c *Controller) Start() {
	c.Resize()
	c.UpdateCache()
}

func (c *Controller) Update() {
	if forge.GetWindow().WindowResized() {
		c.Resize()
	}
	if forge.GetWindow().HasEvents() {
		c.raycast()
	}
}

func (c *Controller) raycast() {
	var target Widget
	pos := forge.GetWindow().MousePosition()

	for _, v := range c.wCache {
		if v.Raycast(pos) {
			target = v
			break
		}
	}

	c.processInteractions(target)
}

func (c *Controller) processInteractions(w Widget) {
	// Dragging Check
	//-------------------------------------------------------------------------
	// An object which is dragging will always be the selected object. This
	// step does NOT start dragging, but rather keeps the drag process going
	// or stops it if there is a mouse_up event.
	if c.selected != nil {
		if c.selected.Dragging() {
			if forge.GetWindow().MouseUp(glfw.MouseButton1) {
				c.selected.HandleEvent(EventDragEnd)
			} else {
				c.selected.HandleEvent(EventDrag)
				return
			}
		}
	}

	// Highlighting Check
	//-------------------------------------------------------------------------
	// This step checks for highlighting changes.
	if w != c.highlighted {
		prev := c.highlighted
		if prev != nil {
			prev.HandleEvent(EventMouseLeave)
		}

		c.highlighted = w
		if c.highlighted != nil {
			c.highlighted.HandleEvent(EventMouseEnter)
		}
	}

	// Selection/Dragging Start Check
	//-------------------------------------------------------------------------
	// This step does selection handling, and starts a dragging sequence if the
	// target object allows dragging. Selection changes are triggered by
	// mouse_down events.
	if forge.GetWindow().MouseDown(glfw.MouseButton1) {
		if w != nil {
			if w != c.selected {
				prev := c.selected
				if prev != nil {
					prev.HandleEvent(EventDeselect)
				}

				c.selected = w

				c.highlighted.HandleEvent(EventSelect)
				c.highlighted.HandleEvent(EventDragStart)
			} else {
				c.highlighted.HandleEvent(EventDragStart)
			}
		} else {
			if c.selected != nil {
				c.selected.HandleEvent(EventDeselect)
			}
			c.selected = nil
		}
	} else if forge.GetWindow().MouseUp(glfw.MouseButton1) {
		// Click Check
		//---------------------------------------------------------------------
		// If we got this far and a mouse_up event is detected, it should be
		// assumed a click event just took place.
		if w != nil {
			c.highlighted.HandleEvent(EventClick)
		}
	} else if forge.GetWindow().MouseWheel() {
		if w != nil {
			c.highlighted.HandleEvent(EventMouseWheel)
		}
	}
}

func NewController() *Controller {
	c := &Controller{}

	c.SetName("UIController")
	forge.GetInstance().MustAssign(c)

	return c
}

func CreateController(name string) *forge.GameObject {
	object := CreateGenericObject(name)

	controller := NewController()

	object.AddComponent(controller)

	return object
}
