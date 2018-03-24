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

const (
	defaultTextboxCursorSize = float32(2)
	defaultTextboxPadding    = float32(4)
)

var _ ui.Widget = &Textbox{}

type Textbox struct {
	ui.BaseComponent

	value string

	state ui.EventType

	WidgetColor       forge.Color
	WidgetColorActive forge.Color
	TextColor         forge.Color

	onChangeFunc func(string)

	background *ui.Graphic
	cursor     *ui.Graphic
	text       *ui.Text

	dragging bool
	focus    bool
}

func (w *Textbox) SetOnChangeFunc(fn func(string)) {
	w.onChangeFunc = fn
}

func (w *Textbox) Rearrange() {

}

func (w *Textbox) Redraw() {
	m := w.RectTransform().ActiveMatrix()

	w.background.Draw(m)
	w.text.Draw(m)
}

func (w *Textbox) Raycast(pos mgl32.Vec2) bool {
	return w.RectTransform().ContainsWorldPosition(pos)
}

func (w *Textbox) Dragging() bool {
	return w.dragging
}

func (w *Textbox) HandleEvent(event ui.EventType) {
	switch event {
	case ui.EventSelect:
		w.focus = true
	case ui.EventDeselect:
		w.focus = false
	}

	w.state = event
}

func NewTextbox() *Textbox {
	w := &Textbox{
		value: "Text",
	}

	w.SetName("UITextbox")
	forge.GetInstance().MustAssign(w)

	return w
}

func CreateTextbox(name string) *forge.GameObject {
	object := ui.CreateGenericObject(name)

	textbox := NewTextbox()

	textbox.background = ui.NewGraphic()
	textbox.text = ui.NewText()

	object.AddComponent(textbox)

	return object
}
