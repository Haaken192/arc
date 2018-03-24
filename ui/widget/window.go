package widget

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/haakenlabs/forge"
	"github.com/haakenlabs/forge/ui"
)

var _ ui.Widget = &Window{}

type Window struct {
	ui.BaseComponent

	background *ui.Graphic
	titlebar   *ui.Graphic
	btnClose   *ui.Graphic
	btnCloseBg *ui.Graphic
	title      *ui.Text

	BgColor   forge.Color
	TextColor forge.Color

	dragging       bool
	draggable      bool
	btnCloseEnable bool
	btnCloseFocus  bool
}

func (w *Window) HandleEvent(event ui.EventType) {
	if w.btnCloseEnable {

	}
}

func (w *Window) Dragging() bool {
	return w.dragging
}

func (w *Window) Raycast(pos mgl32.Vec2) bool {
	bounding := forge.NewRect(
		w.RectTransform().WorldPosition().Add(w.background.Position()),
		w.background.Size(),
	)

	return bounding.Contains(pos)
}

func (w *Window) Redraw() {
	w.background.SetColor(w.BgColor)
	w.title.SetColor(w.TextColor)
	w.btnClose.SetColor(w.TextColor)
	w.btnCloseBg.SetColor(w.BgColor)
}

func (w *Window) Rearrange() {
	w.background.Refresh()
	w.btnClose.Refresh()
	w.btnCloseBg.Refresh()
	w.titlebar.Refresh()
	w.title.Refresh()
}

func (w *Window) OnTransformChanged() {
	w.Rearrange()
}

func (w *Window) SetTitle(title string) {
	w.title.SetValue(title)
}

func (w *Window) Title() string {
	return w.title.Value()
}

func (w *Window) SetCloseButtonEnabled(enabled bool) {
	w.btnCloseEnable = enabled
}

func (w *Window) CloseButtonEnabled() bool {
	return w.btnCloseEnable
}
