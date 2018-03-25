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

package ui

import (
	"github.com/go-gl/mathgl/mgl32"

	"fmt"

	"github.com/haakenlabs/arc/core"
	"github.com/haakenlabs/arc/scene"
	"github.com/haakenlabs/arc/system/instance"
)

type AnchorPreset uint8
type PivotPreset uint8

const (
	AnchorTopLeft AnchorPreset = iota
	AnchorTopCenter
	AnchorTopRight
	AnchorMiddleLeft
	AnchorMiddleCenter
	AnchorMiddleRight
	AnchorBottomLeft
	AnchorBottomCenter
	AnchorBottomRight
	StretchAnchorLeft
	StretchAnchorCenter
	StretchAnchorRight
	StretchAnchorTop
	StretchAnchorMiddle
	StretchAnchorBottom
	StretchAnchorAll
)

const (
	PivotTopLeft PivotPreset = iota
	PivotTopCenter
	PivotTopRight
	PivotMiddleLeft
	PivotMiddleCenter
	PivotMiddleRight
	PivotBottomLeft
	PivotBottomCenter
	PivotBottomRight
)

type RectTransform struct {
	scene.BaseTransform

	rect      core.Rect
	anchorMax mgl32.Vec2
	anchorMin mgl32.Vec2
	offsetMax mgl32.Vec2
	offsetMin mgl32.Vec2
	pivot     mgl32.Vec2
	autoSize  bool
}

func NewRectTransform() *RectTransform {
	t := &RectTransform{
		autoSize: true,
	}

	t.SetRotationN(mgl32.QuatIdent())
	t.SetScaleN(mgl32.Vec3{1.0, 1.0, 1.0})

	t.SetName("RectTransform")
	instance.MustAssign(t)

	return t
}

func RectTransformComponent(g *scene.GameObject) *RectTransform {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*RectTransform); ok {
			return ct
		}
	}

	return nil
}

func (t *RectTransform) Rect() core.Rect {
	return t.rect
}

func (t *RectTransform) AnchorMax() mgl32.Vec2 {
	return t.anchorMax
}

func (t *RectTransform) AnchorMin() mgl32.Vec2 {
	return t.anchorMin
}

func (t *RectTransform) OffsetMax() mgl32.Vec2 {
	return t.offsetMax
}

func (t *RectTransform) OffsetMin() mgl32.Vec2 {
	return t.offsetMin
}

func (t *RectTransform) Pivot() mgl32.Vec2 {
	return t.pivot
}

func (t *RectTransform) Size() mgl32.Vec2 {
	return t.rect.Size()
}

func (t *RectTransform) Autosize() bool {
	return t.autoSize
}

func (t *RectTransform) SetRect(rect core.Rect) {
	t.rect = rect
	t.ComputeOffsets()
	t.Recompute(true)
}

func (t *RectTransform) SetPosition2D(position mgl32.Vec2) {
	t.rect.SetOrigin(position)
	t.ComputeOffsets()
	t.Recompute(true)
}

func (t *RectTransform) SetSize(size mgl32.Vec2) {
	t.rect.SetSize(size)
	t.ComputeOffsets()

	t.Recompute(true)
}

func (t *RectTransform) SetAnchorMax(anchor mgl32.Vec2) {
	t.anchorMax = anchor
	t.ComputeOffsets()
}

func (t *RectTransform) SetAnchorMin(anchor mgl32.Vec2) {
	t.anchorMin = anchor
	t.ComputeOffsets()
}

func (t *RectTransform) SetPivot(pivot mgl32.Vec2) {
	t.pivot = pivot
	t.ComputeOffsets()
}

func (t *RectTransform) SetAnchorPreset(preset AnchorPreset) {
	switch preset {
	case AnchorTopLeft:
		t.anchorMin = mgl32.Vec2{}
		t.anchorMax = mgl32.Vec2{}
		break
	case AnchorTopCenter:
		t.anchorMin = mgl32.Vec2{0.5, 0}
		t.anchorMax = mgl32.Vec2{0.5, 0}
		break
	case AnchorTopRight:
		t.anchorMin = mgl32.Vec2{1, 0}
		t.anchorMax = mgl32.Vec2{1, 0}
		break
	case AnchorMiddleLeft:
		t.anchorMin = mgl32.Vec2{0, 0.5}
		t.anchorMax = mgl32.Vec2{0, 0.5}
		break
	case AnchorMiddleCenter:
		t.anchorMin = mgl32.Vec2{0.5, 0.5}
		t.anchorMax = mgl32.Vec2{0.5, 0.5}
		break
	case AnchorMiddleRight:
		t.anchorMin = mgl32.Vec2{1, 0.5}
		t.anchorMax = mgl32.Vec2{1, 0.5}
		break
	case AnchorBottomLeft:
		t.anchorMin = mgl32.Vec2{0, 1}
		t.anchorMax = mgl32.Vec2{0, 1}
		break
	case AnchorBottomCenter:
		t.anchorMin = mgl32.Vec2{0.5, 1}
		t.anchorMax = mgl32.Vec2{0.5, 1}
		break
	case AnchorBottomRight:
		t.anchorMin = mgl32.Vec2{1, 1}
		t.anchorMax = mgl32.Vec2{1, 1}
		break
	case StretchAnchorLeft:
		t.anchorMin = mgl32.Vec2{0, 0}
		t.anchorMax = mgl32.Vec2{0, 1}
		break
	case StretchAnchorCenter:
		t.anchorMin = mgl32.Vec2{0.5, 0}
		t.anchorMax = mgl32.Vec2{0.5, 1}
		break
	case StretchAnchorRight:
		t.anchorMin = mgl32.Vec2{1, 0}
		t.anchorMax = mgl32.Vec2{1, 1}
		break
	case StretchAnchorTop:
		t.anchorMin = mgl32.Vec2{0, 0}
		t.anchorMax = mgl32.Vec2{1, 0}
		break
	case StretchAnchorMiddle:
		t.anchorMin = mgl32.Vec2{0, 0.5}
		t.anchorMax = mgl32.Vec2{1, 0.5}
		break
	case StretchAnchorBottom:
		t.anchorMin = mgl32.Vec2{0, 1}
		t.anchorMax = mgl32.Vec2{1, 1}
		break
	case StretchAnchorAll:
		t.anchorMin = mgl32.Vec2{0, 0}
		t.anchorMax = mgl32.Vec2{1, 1}
		break
	default:
		break
	}

	t.ComputeOffsets()
}

func (t *RectTransform) SetPivotPreset(preset PivotPreset) {
	switch preset {
	case PivotTopLeft:
		t.pivot = mgl32.Vec2{0, 0}
		break
	case PivotTopCenter:
		t.pivot = mgl32.Vec2{0.5, 0}
		break
	case PivotTopRight:
		t.pivot = mgl32.Vec2{1, 0}
		break
	case PivotMiddleLeft:
		t.pivot = mgl32.Vec2{0, 0.5}
		break
	case PivotMiddleCenter:
		t.pivot = mgl32.Vec2{0.5, 0.5}
		break
	case PivotMiddleRight:
		t.pivot = mgl32.Vec2{1, 0.5}
		break
	case PivotBottomLeft:
		t.pivot = mgl32.Vec2{0, 1}
		break
	case PivotBottomCenter:
		t.pivot = mgl32.Vec2{0.5, 1}
		break
	case PivotBottomRight:
		t.pivot = mgl32.Vec2{1, 1}
		break
	default:
		break
	}

	t.ComputeOffsets()
}

func (t *RectTransform) SetPresets(anchor AnchorPreset, pivot PivotPreset) {
	t.SetAnchorPreset(anchor)
	t.SetPivotPreset(pivot)
}

func (t *RectTransform) SetAutosize(autosize bool) {
	t.autoSize = autosize
}

func (t *RectTransform) Start() {
	t.ComputeOffsets()
	t.Recompute(false)
}

func (t *RectTransform) ComputeOffsets() {
	var parentSize mgl32.Vec2
	var pivotSkew mgl32.Vec2

	parent := t.ParentTransform()
	if parent == nil {
		t.offsetMin = t.rect.Min()
		t.offsetMax = t.rect.Max()
		return
	}

	parentSize = parent.Size()

	pivotSkew[0] = t.rect.Width() * t.pivot.X()
	pivotSkew[1] = t.rect.Height() * t.pivot.Y()

	t.offsetMin[0] = (parentSize.X()*t.anchorMin.X() - pivotSkew.X() + t.rect.Left()) - parentSize.X()*t.anchorMin.X()
	t.offsetMin[1] = (parentSize.Y()*t.anchorMin.Y() - pivotSkew.Y() + t.rect.Top()) - parentSize.Y()*t.anchorMin.Y()

	t.offsetMax[0] = (parentSize.X()*t.anchorMin.X() + t.offsetMin.X()) + t.Size().X() - parentSize.X()*t.anchorMax.X()
	t.offsetMax[1] = (parentSize.Y()*t.anchorMin.Y() + t.offsetMin.Y()) + t.Size().Y() - parentSize.Y()*t.anchorMax.Y()

	if t.pivot[1] == 0.5 {
		fmt.Printf("--------------\n")
		fmt.Printf("parentSize: %v\n", parentSize)
		fmt.Printf("      size: %v\n", t.Size())
		fmt.Printf("  position: %v\n", t.Position())
		fmt.Printf("      left: %v top: %v\n", t.rect.Left(), t.rect.Top())
		fmt.Printf("      rect: %v\n", t.rect)
		fmt.Printf("     pivot: %v\n", t.pivot)
		fmt.Printf(" pivotSkew: %v\n", pivotSkew)
		fmt.Printf(" anchorMin: %v\n", t.anchorMin)
		fmt.Printf(" anchorMax: %v\n", t.anchorMax)
		fmt.Printf(" offsetMin: %v\n", t.offsetMin)
		fmt.Printf(" offsetMax: %v\n", t.offsetMax)
	}
}

func (t *RectTransform) Recompute(updateChildren bool) {
	var aMin mgl32.Vec2
	var aMax mgl32.Vec2
	var aSize mgl32.Vec2
	var pSize mgl32.Vec2

	if t.ParentTransform() != nil {
		pSize = t.ParentTransform().Size()
	}

	aMin = mgl32.Vec2{
		pSize.X() * t.anchorMin.X(),
		pSize.Y() * t.anchorMin.Y(),
	}
	aMax = mgl32.Vec2{
		pSize.X() * t.anchorMax.X(),
		pSize.Y() * t.anchorMax.Y(),
	}
	aSize = aMax.Add(t.offsetMax).Sub(aMin.Add(t.offsetMin))

	pos := aMin.Add(t.offsetMin)

	if t.autoSize {
		t.rect.SetSize(aSize)
	}

	if t.pivot[1] == 0.5 {
		fmt.Printf("--------------\n")
		fmt.Printf("parentSize: %v\n", pSize)
		fmt.Printf("      size: %v\n", t.Size())
		fmt.Printf("  position: %v\n", t.Position())
		fmt.Printf("      left: %v top: %v\n", t.rect.Left(), t.rect.Top())
		fmt.Printf("      rect: %v\n", t.rect)
		fmt.Printf("     pivot: %v\n", t.pivot)
		fmt.Printf(" anchorMin: %v\n", t.anchorMin)
		fmt.Printf(" anchorMax: %v\n", t.anchorMax)
		fmt.Printf(" offsetMin: %v\n", t.offsetMin)
		fmt.Printf(" offsetMax: %v\n", t.offsetMax)

		fmt.Printf(" recompute position: %v\n", t.Position())
	}

	t.BaseTransform.SetPositionN(pos.Vec3(0.0))
	t.BaseTransform.Recompute(updateChildren)
}

func (t *RectTransform) WorldPosition() mgl32.Vec2 {
	return t.ActiveMatrix().Col(3).Vec2()
}

func (t *RectTransform) ContainsWorldPosition(position mgl32.Vec2) bool {
	return core.NewRect(t.WorldPosition(), t.Size()).Contains(position)
}

func (t *RectTransform) ParentTransform() *RectTransform {
	if t.GameObject() != nil {
		if parent := t.GameObject().Parent(); parent != nil {
			if obj, ok := parent.Transform().(*RectTransform); ok {
				return obj
			}
		}
	}

	return nil
}
