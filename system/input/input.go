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

package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/haakenlabs/arc/core"
)

func KeyDown(key glfw.Key) bool {
	return core.GetWindowSystem().KeyDown(key)
}

func KeyUp(key glfw.Key) bool {
	return core.GetWindowSystem().KeyUp(key)
}

func KeyPressed() bool {
	return core.GetWindowSystem().KeyPressed()
}

func MouseDown(button glfw.MouseButton) bool {
	return core.GetWindowSystem().MouseDown(button)
}

func MouseUp(button glfw.MouseButton) bool {
	return core.GetWindowSystem().MouseUp(button)
}

func MouseWheelX() float64 {
	return core.GetWindowSystem().MouseWheelX()
}

func MouseWheelY() float64 {
	return core.GetWindowSystem().MouseWheelY()
}

func MouseWheel() bool {
	return core.GetWindowSystem().MouseWheel()
}

func MouseMoved() bool {
	return core.GetWindowSystem().MouseMoved()
}

func MousePressed() bool {
	return core.GetWindowSystem().MousePressed()
}

func MousePosition() mgl32.Vec2 {
	return core.GetWindowSystem().MousePosition()
}

func WindowResized() bool {
	return core.GetWindowSystem().WindowResized()
}

func ShouldClose() bool {
	return core.GetWindowSystem().ShouldClose()
}

func HandleEvents() {
	core.GetWindowSystem().HandleEvents()
}

func HasEvents() bool {
	return core.GetWindowSystem().HasEvents()
}
