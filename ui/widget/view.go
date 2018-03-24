package widget

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/haakenlabs/forge/ui"
)

var _ ui.Widget = &View{}

type View struct {
	ui.BaseComponent
}

func (w *View) HandleEvent(event ui.EventType) {
}

func (w *View) Dragging() bool {
	return false
}

func (w *View) Raycast(pos mgl32.Vec2) bool {
	return false
}

func (w *View) Redraw() {
}

func (w *View) Rearrange() {
}