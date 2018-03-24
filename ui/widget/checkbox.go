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

type CheckState int

const (
	CheckStateOff CheckState = iota
	CheckStateMixed
	CheckStateOn
)

var _ ui.Widget = &Checkbox{}

var defaultCheckboxSize = mgl32.Vec2{16, 16}
var defaultCheckboxCheckSize = mgl32.Vec2{9, 9}

type CheckboxGroup struct {
	ui.BaseComponent

	checkboxes []*Checkbox
}

type Checkbox struct {
	ui.BaseComponent

	state      CheckState
	eventState ui.EventType

	BgColor       forge.Color
	BgColorActive forge.Color
	CheckMixColor forge.Color
	CheckOnColor  forge.Color

	onChangeFunc func(CheckState)

	background *ui.Graphic
	check      *ui.Graphic
	text       *ui.Text
}

func (w *Checkbox) SetOnChangeFunc(fn func(CheckState)) {
	w.onChangeFunc = fn
}

func (w *Checkbox) Dragging() bool {
	return false
}

func (w *Checkbox) HandleEvent(event ui.EventType) {
	switch event {
	case ui.EventClick:
		if w.state == CheckStateOn {
			w.state = CheckStateOff
		} else {
			w.state = CheckStateOn
		}
		if w.onChangeFunc != nil {
			w.onChangeFunc(w.state)
		}
	}

	w.eventState = event
}

func (w *Checkbox) Redraw() {
	switch w.eventState {
	case ui.EventClick:
		fallthrough
	case ui.EventMouseEnter:
		w.background.SetColor(w.BgColorActive)
	default:
		w.background.SetColor(w.BgColor)
	}
	switch w.state {
	case CheckStateMixed:
		w.check.SetColor(w.CheckMixColor)
	case CheckStateOn:
		w.check.SetColor(w.CheckOnColor)
	}

	m := w.RectTransform().ActiveMatrix()

	w.background.Draw(m)
	if w.state == CheckStateOn {
		w.check.Draw(m)
	}
	w.text.Draw(m)
}

func (w *Checkbox) Raycast(pos mgl32.Vec2) bool {
	bounding := forge.NewRect(
		w.RectTransform().WorldPosition().Add(w.background.Position()),
		w.background.Size(),
	)

	return bounding.Contains(pos)
}

func (w *Checkbox) Start() {
	w.Rearrange()
}

func (w *Checkbox) Rearrange() {
	w.check.Refresh()
	w.check.SetPosition(ui.Align(w.check.Rect(), w.background.Rect(), ui.AlignmentMiddleCenter))
	w.background.Refresh()

	w.text.Refresh()
	textPos := ui.Align(w.text.Rect(), w.background.Rect(), ui.AlignmentMiddleLeft)
	textPos = textPos.Add(mgl32.Vec2{w.background.Size().X() + 8, 0})
	w.text.SetPosition(textPos)
}

func (w *Checkbox) SetValue(value string) {
	w.text.SetValue(value)
	w.Rearrange()
}

func (w *CheckboxGroup) AddCheckbox(checkbox ...*Checkbox) {
	w.checkboxes = append(w.checkboxes, checkbox...)
}

func NewCheckbox() *Checkbox {
	w := &Checkbox{}

	w.BgColor = ui.Styles.WidgetColor
	w.BgColorActive = ui.Styles.WidgetColorActive
	w.CheckMixColor = ui.Styles.WidgetColor
	w.CheckOnColor = ui.Styles.WidgetColorPrimary

	w.SetName("UICheckbox")
	forge.GetInstance().MustAssign(w)

	return w
}

func NewCheckboxGroup() *CheckboxGroup {
	w := &CheckboxGroup{}

	w.SetName("UICheckboxGroup")
	forge.GetInstance().MustAssign(w)

	return w
}

func CheckboxComponent(g *forge.GameObject) *Checkbox {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Checkbox); ok {
			return ct
		}
	}

	return nil
}

func CreateCheckbox(name string) *forge.GameObject {
	object := ui.CreateGenericObject(name)

	checkbox := NewCheckbox()

	checkbox.background = ui.NewGraphic()
	checkbox.background.SetSize(defaultCheckboxSize)
	checkbox.check = ui.NewGraphic()
	checkbox.check.SetSize(defaultCheckboxCheckSize)
	checkbox.text = ui.NewText()
	checkbox.text.SetValue("Checkbox")

	object.AddComponent(checkbox)

	return object
}

func CreateCheckboxGroup(name string, checkboxes ...*Checkbox) *forge.GameObject {
	object := ui.CreateGenericObject(name)

	group := NewCheckboxGroup()
	group.AddCheckbox(checkboxes...)

	object.AddComponent(group)

	return object
}
