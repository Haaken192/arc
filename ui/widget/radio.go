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

type RadioState int

const (
	RadioStateOff RadioState = iota
	RadioStateMixed
	RadioStateOn
)

type RadioGroup struct {
	ui.BaseComponent

	radios []*Radio
}

var defaultRadioSize = mgl32.Vec2{16, 16}
var defaultRadioCheckSize = mgl32.Vec2{9, 9}

var _ ui.Widget = &Radio{}

type Radio struct {
	ui.BaseComponent

	state      RadioState
	eventState ui.EventType

	BgColor       forge.Color
	BgColorActive forge.Color
	RadioMixColor forge.Color
	RadioOnColor  forge.Color

	backgroundColor forge.Color
	tint            forge.Color

	onChangeFunc func(RadioState)

	background *ui.Graphic
	check      *ui.Graphic
	text       *ui.Text
}

func (w *Radio) Dragging() bool {
	return false
}

func (w *Radio) HandleEvent(event ui.EventType) {
	switch event {
	case ui.EventClick:
		if w.state == RadioStateOn {
			w.state = RadioStateOff
		} else {
			w.state = RadioStateOn
		}
		if w.onChangeFunc != nil {
			w.onChangeFunc(w.state)
		}
	}

	w.eventState = event
}

func (w *Radio) Redraw() {
	switch w.eventState {
	case ui.EventClick:
		fallthrough
	case ui.EventMouseEnter:
		w.background.SetColor(w.BgColorActive)
	default:
		w.background.SetColor(w.BgColor)
	}
	switch w.state {
	case RadioStateMixed:
		w.check.SetColor(w.RadioMixColor)
	case RadioStateOn:
		w.check.SetColor(w.RadioOnColor)
	}

	m := w.RectTransform().ActiveMatrix()

	w.background.Draw(m)
	if w.state == RadioStateOn {
		w.check.Draw(m)
	}
	w.text.Draw(m)
}

func (w *Radio) Raycast(pos mgl32.Vec2) bool {
	bounding := forge.NewRect(
		w.RectTransform().WorldPosition().Add(w.background.Position()),
		w.background.Size(),
	)

	return bounding.Contains(pos)
}

func (w *Radio) Rearrange() {
	w.check.Refresh()
	w.check.SetPosition(ui.Align(w.check.Rect(), w.background.Rect(), ui.AlignmentMiddleCenter))
	w.background.Refresh()

	w.text.Refresh()
	textPos := ui.Align(w.text.Rect(), w.background.Rect(), ui.AlignmentMiddleLeft)
	textPos = textPos.Add(mgl32.Vec2{w.background.Size().X() + 8, 0})
	w.text.SetPosition(textPos)
}

func (w *Radio) SetValue(value string) {
	w.text.SetValue(value)
	w.Rearrange()
}

func (w *Radio) Start() {
	w.Rearrange()
}

func (w *RadioGroup) AddRadio(radio ...*Radio) {
	w.radios = append(w.radios, radio...)
}

func NewRadio() *Radio {
	w := &Radio{}

	w.SetName("UIRadio")
	forge.GetInstance().MustAssign(w)

	return w
}

func NewRadioGroup() *RadioGroup {
	w := &RadioGroup{}

	w.SetName("UIRadioGroup")
	forge.GetInstance().MustAssign(w)

	return w
}

func CreateRadio(name string) *forge.GameObject {
	object := ui.CreateGenericObject(name)

	radio := NewRadio()

	radio.background = ui.NewGraphic()
	radio.background.SetSize(defaultRadioSize)
	radio.check = ui.NewGraphic()
	radio.check.SetSize(defaultRadioCheckSize)
	radio.text = ui.NewText()
	radio.text.SetValue("Radio")

	object.AddComponent(radio)

	return object
}

func CreateRadioGroup(name string, radios ...*Radio) *forge.GameObject {
	object := ui.CreateGenericObject(name)

	group := NewRadioGroup()
	group.AddRadio(radios...)

	object.AddComponent(group)

	return object
}
