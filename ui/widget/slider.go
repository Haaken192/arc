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
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/haakenlabs/forge"
	"github.com/haakenlabs/forge/system/input"
	"github.com/haakenlabs/forge/ui"
)

const (
	defaultSliderThumbSize = float32(16)
	defaultSliderHeight    = float32(10)
)

var _ ui.Widget = &Slider{}

type Slider struct {
	ui.BaseComponent

	value float64
	min   float64
	max   float64

	state ui.EventType

	intMode bool

	WidgetColor        forge.Color
	WidgetColorActive  forge.Color
	WidgetColorPrimary forge.Color

	onChangeFunc func(float64)

	background  *ui.Graphic
	activeTrack *ui.Graphic
	thumb       *ui.Graphic

	dragging bool
}

func (w *Slider) SetValue(value float64) {
	if w.intMode {
		value = math.Round(value)
	}
	w.value = mgl64.Clamp(value, w.min, w.max)
	w.Rearrange()

	if w.onChangeFunc != nil {
		w.onChangeFunc(w.value)
	}
}

func (w *Slider) SetRelValue(value float64) {
	newValue := (w.max - w.min) * mgl64.Clamp(value, 0.0, 1.0)
	if w.intMode {
		newValue = math.Round(newValue)
	}

	w.SetValue(newValue)
}

func (w *Slider) SetMinValue(value float64) {
	if w.intMode {
		value = math.Round(value)
	}
	if value > w.max {
		return
	}

	w.min = value

	w.SetValue(w.value)
}

func (w *Slider) SetMaxValue(value float64) {
	if w.intMode {
		value = math.Round(value)
	}
	if value < w.min {
		return
	}

	w.max = value

	w.SetValue(w.value)
}

func (w *Slider) Value() float64 {
	if w.intMode {
		return math.Round(w.value)
	}

	return w.value
}

func (w *Slider) MaxValue() float64 {
	if w.intMode {
		return math.Round(w.max)
	}

	return w.max
}

func (w *Slider) MinValue() float64 {
	if w.intMode {
		return math.Round(w.min)
	}

	return w.min
}

func (w *Slider) SetOnChangeFunc(fn func(float64)) {
	w.onChangeFunc = fn
}

func (w *Slider) Dragging() bool {
	return w.dragging
}

func (w *Slider) Raycast(pos mgl32.Vec2) bool {
	return w.RectTransform().ContainsWorldPosition(pos)
}

func (w *Slider) HandleEvent(event ui.EventType) {
	pos := input.MousePosition()
	relPos := w.RectTransform().WorldPosition()
	size := w.RectTransform().Size()

	rel := (pos.X() - relPos.X()) / (relPos.X() + size.X() - relPos.X())

	switch event {
	case ui.EventDragStart:
		fallthrough
	case ui.EventDrag:
		w.dragging = true
		w.SetRelValue(float64(rel))
	case ui.EventClick:
		w.dragging = false
		w.SetRelValue(float64(rel))
	default:
		w.dragging = false
	}

	w.state = event
}

func (w *Slider) Redraw() {
	switch w.state {
	case ui.EventMouseEnter:
		w.background.SetColor(w.WidgetColorActive)
	default:
		w.background.SetColor(w.WidgetColor)
	}

	w.activeTrack.SetColor(w.WidgetColorPrimary)
	w.thumb.SetColor(w.WidgetColorActive)

	m := w.RectTransform().ActiveMatrix()

	w.background.Draw(m)
	w.activeTrack.Draw(m)
	w.thumb.Draw(m)
}

func (w *Slider) Rearrange() {
	width := w.RectTransform().Size().X()
	//w.RectTransform().SetSize(mgl32.Vec2{width, defaultSliderThumbSize})

	activeWidth := float32(math.Floor(float64(width) * w.relativeValue()))

	w.activeTrack.SetSize(mgl32.Vec2{activeWidth, defaultSliderHeight})
	w.background.SetSize(mgl32.Vec2{width, defaultSliderHeight})
	w.background.SetPosition(mgl32.Vec2{0, 0})

	w.thumb.SetSize(mgl32.Vec2{defaultSliderThumbSize, defaultSliderThumbSize})
	thumbPos := ui.Align(w.thumb.Rect(), w.activeTrack.Rect(), ui.AlignmentMiddleLeft)
	thumbPos = thumbPos.Add(mgl32.Vec2{activeWidth - w.thumb.Size().X()/2, 0})
	w.thumb.SetPosition(thumbPos)

	w.background.Refresh()
	w.activeTrack.Refresh()
	w.thumb.Refresh()
}

func (w *Slider) Start() {
	w.Rearrange()
}

func (w *Slider) relativeValue() float64 {
	return (w.value - w.min) / (w.max - w.min)
}

func NewSlider() *Slider {
	w := &Slider{
		value: 0.5,
		min:   0.0,
		max:   1.0,
	}

	w.WidgetColor = ui.Styles.WidgetColor
	w.WidgetColorActive = ui.Styles.WidgetColorActive
	w.WidgetColorPrimary = ui.Styles.WidgetColorPrimary

	w.SetName("UISlider")
	forge.GetInstance().MustAssign(w)

	return w
}

func SliderComponent(g *forge.GameObject) *Slider {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Slider); ok {
			return ct
		}
	}

	return nil
}

func CreateSlider(name string) *forge.GameObject {
	object := ui.CreateGenericObject(name)

	slider := NewSlider()

	slider.background = ui.NewGraphic()
	slider.activeTrack = ui.NewGraphic()
	slider.thumb = ui.NewGraphic()

	object.AddComponent(slider)

	return object
}
