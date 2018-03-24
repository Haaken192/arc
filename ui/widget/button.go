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

var _ ui.Widget = &Button{}

var defaultButtonSize = mgl32.Vec2{96, 32}

type Button struct {
	ui.BaseComponent

	BgColor         forge.Color
	BgColorActive   forge.Color
	TextColor       forge.Color
	TextColorActive forge.Color

	value      string
	eventState ui.EventType

	onPressedFunc func()

	background *ui.Graphic
	text       *ui.Text
}

func NewButton() *Button {
	w := &Button{
		value:      "Button",
		background: ui.NewGraphic(),
		text:       ui.NewText(),
	}

	w.TextColor = ui.Styles.TextColor
	w.TextColorActive = ui.Styles.TextColorActive
	w.BgColor = ui.Styles.WidgetColor
	w.BgColorActive = ui.Styles.WidgetColorActive

	w.SetName("UIButton")
	forge.GetInstance().MustAssign(w)

	w.background.SetColor(w.BgColor)
	w.background.SetSize(defaultButtonSize)

	w.text.SetFontSize(12)
	w.text.SetValue(w.value)
	w.text.SetColor(w.TextColor)

	return w
}

func (w *Button) SetValue(value string) {
	w.value = value
	w.text.SetValue(value)
}

func (w *Button) Value() string {
	return w.value
}

func (w *Button) SetOnPressedFunc(fn func()) {
	w.onPressedFunc = fn
}

func (w *Button) Dragging() bool {
	return false
}

func (w *Button) HandleEvent(event ui.EventType) {
	switch event {
	case ui.EventClick:
		if w.onPressedFunc != nil {
			w.onPressedFunc()
		}
	}

	w.eventState = event
}

func (w *Button) Start() {
	w.Rearrange()
}

func (w *Button) Raycast(pos mgl32.Vec2) bool {
	return w.RectTransform().ContainsWorldPosition(pos)
}

func (w *Button) Redraw() {
	switch w.eventState {
	case ui.EventClick:
		fallthrough
	case ui.EventMouseEnter:
		w.background.SetColor(w.BgColorActive)
		w.text.SetColor(w.TextColorActive)
	default:
		w.background.SetColor(w.BgColor)
		w.text.SetColor(w.TextColor)
	}

	m := w.GetTransform().ActiveMatrix()

	w.background.Draw(m)
	w.text.Draw(m)
}

func (w *Button) Rearrange() {
	w.text.Refresh()
	textPos := ui.Align(w.text.Rect(), w.RectTransform().Rect(), ui.AlignmentMiddleCenter)
	w.text.SetPosition(textPos)

	w.background.Refresh()
}

func ButtonComponent(g *forge.GameObject) *Button {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Button); ok {
			return ct
		}
	}

	return nil
}

func CreateButton(name string) *forge.GameObject {
	object := ui.CreateGenericObject(name)
	rt := ui.RectTransformComponent(object)
	rt.SetSize(defaultButtonSize)

	button := NewButton()

	object.AddComponent(button)

	return object
}
