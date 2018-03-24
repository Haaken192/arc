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

	"github.com/haakenlabs/arc/core"
	"github.com/haakenlabs/arc/scene"
	"github.com/haakenlabs/arc/system/instance"
	"github.com/haakenlabs/arc/ui"
)

var _ ui.Widget = &Label{}

type Label struct {
	ui.BaseComponent

	TextColor core.Color

	text *ui.Text
}

func NewLabel() *Label {
	w := &Label{
		text: ui.NewText(),
	}

	w.TextColor = ui.Styles.TextColor

	w.SetName("UILabel")
	instance.MustAssign(w)

	w.text.SetValue("Label")
	w.text.SetFontSize(ui.Styles.TextSize)
	w.text.SetColor(w.TextColor)

	return w
}

func (w *Label) SetValue(value string) {
	w.text.SetValue(value)
	w.Rearrange()
}

func (w *Label) Value() string {
	return w.text.Value()
}

func (w *Label) SetFontSize(size int32) {
	w.text.SetFontSize(size)
	w.Rearrange()
}

func (w *Label) FontSize(size int32) int32 {
	return w.text.FontSize()
}

func (w *Label) OnActivate() {
	w.Rearrange()
}

func (w *Label) Start() {
	w.Rearrange()
}

func (w *Label) Raycast(pos mgl32.Vec2) bool {
	return false
}

func (w *Label) Dragging() bool {
	return false
}

func (w *Label) HandleEvent(event ui.EventType) {}

func (w *Label) Rearrange() {
	// DANGER
	w.text.Refresh()
	w.RectTransform().SetSize(w.text.Size())
}

func (w *Label) OnTransformChanged() {

}

func (w *Label) Redraw() {
	m := w.RectTransform().ActiveMatrix()

	w.text.Draw(m)
}

func LabelComponent(g *scene.GameObject) *Label {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Label); ok {
			return ct
		}
	}

	return nil
}

func CreateLabel(name string) *scene.GameObject {
	object := ui.CreateGenericObject(name)

	object.AddComponent(NewLabel())

	return object
}
