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

package core

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var _ System = &TimeSystem{}

var timeInst *TimeSystem

const SysNameTime = "time"

const fixedTime = float64(0.05)

// TimeSystem implements a time system.
type TimeSystem struct {
	frameTime     float64
	deltaTime     float64
	nextLogicTick float64
	frame         uint64
}

// Setup sets up the System.
func (t *TimeSystem) Setup() error {
	if timeInst != nil {
		return ErrSystemInit(SysNameTime)
	}
	timeInst = t

	return nil
}

// Teardown tears down the System.
func (t *TimeSystem) Teardown() {

}

// Name returns the name of the System.
func (t *TimeSystem) Name() string {
	return SysNameTime
}

func (t *TimeSystem) FrameTime() float64 {
	return t.frameTime
}

func (t *TimeSystem) DeltaTime() float64 {
	return t.deltaTime
}

func (t *TimeSystem) FixedTime() float64 {
	return fixedTime
}

func (t *TimeSystem) Delta() float64 {
	return t.deltaTime
}

func (t *TimeSystem) Now() float64 {
	return glfw.GetTime()
}

func (t *TimeSystem) FrameStart() {
	t.frameTime = t.Now()
}

func (t *TimeSystem) FrameEnd() {
	t.deltaTime = t.Now() - t.frameTime
	t.frame++
}

func (t *TimeSystem) Frame() uint64 {
	return t.frame
}

func (t *TimeSystem) LogicTick() {
	t.nextLogicTick += fixedTime
}

func (t *TimeSystem) LogicUpdate() bool {
	return t.Now() > t.nextLogicTick
}

// NewTime creates a new time system.
func NewTimeSystem() *TimeSystem {
	return &TimeSystem{}
}

// GetTime gets the time system from the current app.
func GetTimeSystem() *TimeSystem {
	return timeInst
}
