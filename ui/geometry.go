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
	"github.com/haakenlabs/forge"
)

type Alignment uint8

const (
	AlignmentTopLeft Alignment = iota
	AlignmentTopCenter
	AlignmentTopRight
	AlignmentMiddleLeft
	AlignmentMiddleCenter
	AlignmentMiddleRight
	AlignmentBottomLeft
	AlignmentBottomCenter
	AlignmentBottomRight
)

func Align(rect, parent forge.Rect, alignment Alignment) mgl32.Vec2 {
	offset := mgl32.Vec2{}
	dot := mgl32.Vec2{}

	switch alignment {
	case AlignmentTopLeft:
	case AlignmentTopCenter:
		offset[0] = 0.5
	case AlignmentTopRight:
		offset[0] = 1.0
	case AlignmentMiddleLeft:
		offset[1] = 0.5
	case AlignmentMiddleCenter:
		offset[0] = 0.5
		offset[1] = 0.5
	case AlignmentMiddleRight:
		offset[0] = 1.0
		offset[1] = 0.5
	case AlignmentBottomLeft:
		offset[1] = 1.0
	case AlignmentBottomCenter:
		offset[0] = 0.5
		offset[1] = 1.0
	case AlignmentBottomRight:
		offset[0] = 1.0
		offset[1] = 1.0
	}

	dot[0] = (parent.Width() - rect.Width()) * offset[0]
	dot[1] = (parent.Height() - rect.Height()) * offset[1]

	dot.Add(rect.Origin())

	return dot
}
