package widget

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/haakenlabs/forge/ui"
)

var _ ui.Widget = &Scrollbar{
}

type Scrollbar struct {
	ui.BaseComponent
}

func (w *Scrollbar) HandleEvent(event ui.EventType) {
}

func (w *Scrollbar) Dragging() bool {
	return false
}

func (w *Scrollbar) Raycast(pos mgl32.Vec2) bool {
	return false
}

func (w *Scrollbar) Redraw() {
}

func (w *Scrollbar) Rearrange() {
}
