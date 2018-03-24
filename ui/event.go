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

type EventType int

const (
	EventUnknown EventType = iota
	EventInput
	EventValueChanged
	EventClick
	EventMouseEnter
	EventMouseLeave
	EventMouseWheel
	EventDragStart
	EventDrag
	EventDragEnd
	EventSelect
	EventDeselect
)

func (e EventType) String() string {
	switch e {
	case EventUnknown:
		return "EventUnknown"
	case EventInput:
		return "EventInput"
	case EventValueChanged:
		return "EventValueChanged"
	case EventClick:
		return "EventClick"
	case EventMouseEnter:
		return "EventMouseEnter"
	case EventMouseLeave:
		return "EventMouseLeave"
	case EventMouseWheel:
		return "EventMouseWheel"
	case EventDragStart:
		return "EventDragStart"
	case EventDrag:
		return "EventDrag"
	case EventDragEnd:
		return "EventDragEnd"
	case EventSelect:
		return "EventSelect"
	case EventDeselect:
		return "EventDeselect"
	default:
		return "Unrecognized EventType: " + string(int(e))
	}
}
