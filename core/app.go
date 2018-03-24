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
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

const (
	fixedTime    = float64(0.05)
	maxFrameSkip = 5

	builtinAssets = "<builtin>:builtin.json"
)

var (
	// appInst is the currently running app.
	appInst *App

	// appOnce is used to set the appInst variable only once.
	appOnce sync.Once
)

// App is the backbone of any Arc app.
type App struct {
	// Name is the name of this app.
	Name string

	// Company is the name of the company responsible for this app.
	Company string

	// PreSetupFunc is a callback invoked prior to app setup.
	PreSetupFunc func() error

	// PostSetupFunc is a callback invoked after app setup.
	PostSetupFunc func() error

	// PreTeardownFunc is a callback invoked prior to app teardown.
	PreTeardownFunc func()

	// PostTeardownFunc is a callback invoked after app teardown.
	PostTeardownFunc func()

	systems []System
	running bool
}

// Setup sets up the App.
func (a *App) Setup() error {
	if appInst != nil {
		return errors.New("app already created")
	}
	setApp(a)

	// Register required systems.
	a.RegisterSystem(NewWindowSystem())
	a.RegisterSystem(NewInstanceSystem())
	a.RegisterSystem(NewAssetSystem())
	a.RegisterSystem(NewTimeSystem())

	if a.PreSetupFunc != nil {
		if err := a.PreSetupFunc(); err != nil {
			return err
		}
	}

	for i := range a.systems {
		logrus.Debug("Setting up system: ", a.systems[i].Name())

		if err := a.systems[i].Setup(); err != nil {
			return err
		}
	}

	if a.PostSetupFunc != nil {
		if err := a.PostSetupFunc(); err != nil {
			return err
		}
	}

	return nil
}

// Teardown tears down the app.
func (a *App) Teardown() {
	if a.PreTeardownFunc != nil {
		a.PreTeardownFunc()
	}

	for i := len(a.systems) - 1; i >= 0; i-- {
		logrus.Debug("Tearing down system: ", a.systems[i].Name())

		a.systems[i].Teardown()
	}

	if a.PostTeardownFunc != nil {
		a.PostTeardownFunc()
	}
}

func (a *App) Run() error {
	a.running = true

	a.setupSignalHandler()

	frame := 0
	loops := 0

	time := a.MustSystem(SysNameTime).(*TimeSystem)
	window := a.MustSystem(SysNameWindow).(*WindowSystem)

	for a.running {
		a.running = !window.ShouldClose()

		time.FrameStart()

		frame++

		a.onUpdate()

		loops = 0
		for time.LogicUpdate() && loops < maxFrameSkip {
			time.LogicTick()
			a.onFixedUpdate()
			loops++
		}

		window.ClearBuffers()
		a.onDisplay()
		window.SwapBuffers()

		window.HandleEvents()
		time.FrameEnd()
	}

	return nil
}

// Quit instructs the App to shutdown by setting the running variable to false.
func (a *App) Quit() {
	a.running = false
}

// RegisterSystem registers a system with the App. A system can only be added
// once, it is an error to add a system more than once. Systems are initialized
// in the order they are added and torn down in the reverse order.
func (a *App) RegisterSystem(s System) {
	// Check for existing system.
	if a.SystemRegistered(s.Name()) {
		panic(ErrSystemExists(s.Name()))
	}

	// Add system to the systems slice.
	a.systems = append(a.systems, s)

	logrus.Debug("Registered system: ", s.Name())
}

// SystemRegistered returns true if the system with the given name is registered
// with this App.
func (a *App) SystemRegistered(name string) bool {
	for i := range a.systems {
		if a.systems[i].Name() == name {
			return true
		}
	}

	return false
}

// System returns a system by the given name.
func (a *App) System(name string) (System, error) {
	for i := range a.systems {
		if a.systems[i].Name() == name {
			return a.systems[i], nil
		}
	}

	return nil, ErrSystemNotFound(name)
}

// MustSystem is like System, but panics if the system cannot be found.
func (a *App) MustSystem(name string) System {
	s, err := a.System(name)
	if err != nil {
		panic(err)
	}

	return s
}

func (a *App) setupSignalHandler() {
	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	go handleSignal(s, a)
}

func (a *App) onDisplay() {
	//if s := a.ActiveScene(); s != nil {
	//	sg := s.Graph()
	//	if sg.Dirty() {
	//		sg.Update()
	//	}
	//
	//	cameras := s.cameras
	//	for i := range cameras {
	//		cameras[i].Render()
	//	}
	//
	//	sg.SendMessage(MessageGUIRender)
	//}
}

func (a *App) onUpdate() {
	//if s := a.ActiveScene(); s != nil {
	//	sg := s.Graph()
	//	if sg.Dirty() {
	//		sg.Update()
	//	}
	//
	//	if !s.started {
	//		s.started = true
	//		sg.SendMessage(MessageStart)
	//	}
	//
	//	sg.SendMessage(MessageUpdate)
	//	sg.SendMessage(MessageLateUpdate)
	//}
}

func (a *App) onFixedUpdate() {
	//if s := a.ActiveScene(); s != nil {
	//	sg := s.Graph()
	//	if sg.Dirty() {
	//		sg.Update()
	//	}
	//
	//	sg.SendMessage(MessageFixedUpdate)
	//}
}

func handleSignal(s chan os.Signal, a *App) {
	<-s
	a.Quit()
}

/// / NewApp creates a new application.
func NewApp() *App {
	a := &App{}

	return a
}

// CurrentApp returns the currently running app.
func CurrentApp() *App {
	return appInst
}

// setApp sets the App, but only once.
func setApp(a *App) {
	appOnce.Do(func() {
		appInst = a
	})
}
