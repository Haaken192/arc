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

package widget

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/haakenlabs/forge"
	"github.com/haakenlabs/forge/ui"
)

var _ ui.Widget = &Image{}

type Image struct {
	ui.BaseComponent

	graphic *ui.Graphic
}

func (w *Image) Color() forge.Color {
	return w.graphic.Color()
}

func (w *Image) Texture() *forge.Texture2D {
	return w.graphic.Texture()
}

func (w *Image) SetColor(color forge.Color) {
	w.graphic.SetColor(color)
}

func (w *Image) SetTexture(texture *forge.Texture2D) {
	w.graphic.SetTexture(texture)
}

func (w *Image) OnActivate() {
	w.Rearrange()
}

func (w *Image) OnTransformChanged() {
	w.Rearrange()
}

func (w *Image) Start() {
	w.Rearrange()
}

func (w *Image) Dragging() bool {
	return false
}

func (w *Image) HandleEvent(event ui.EventType) {}

func (w *Image) Raycast(pos mgl32.Vec2) bool {
	return false
}

func (w *Image) Redraw() {
	m := w.GetTransform().ActiveMatrix()

	w.graphic.Draw(m)
}

func (w *Image) Rearrange() {
	w.graphic.SetSize(w.RectTransform().Size())
	w.graphic.Refresh()
}

func NewImage() *Image {
	w := &Image{
		graphic: ui.NewGraphic(),
	}

	w.SetName("UIImage")
	forge.GetInstance().MustAssign(w)

	w.graphic.SetColor(ui.Styles.BackgroundColor)

	return w
}

func ImageComponent(g *forge.GameObject) *Image {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Image); ok {
			return ct
		}
	}

	return nil
}

func CreateImage(name string) *forge.GameObject {
	object := ui.CreateGenericObject(name)

	image := NewImage()

	object.AddComponent(image)

	return object
}

func CreatePanel(name string) *forge.GameObject {
	object := CreateImage(name)

	rt := ui.RectTransformComponent(object)
	rt.SetSize(mgl32.Vec2{480, 320})

	return object
}
