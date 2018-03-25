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
	"fmt"

	"github.com/sirupsen/logrus"
)

type Scene interface {
	// OnActivate is called when the scene transitions to the active state.
	OnActivate()

	// OnDeactivate is called when the scene transitions to the inactive state.
	OnDeactivate()

	// Display is called when the scene should render.
	Display()

	// Display is called at fixed intervals for logic updates.
	FixedUpdate()

	// Display is called every frame for logic updates.
	Update()

	// Load is called when the scene is being initialized.
	Load() error

	// Loaded reports if the scene has been loaded.
	Loaded() bool

	// Name returns the name of the scene.
	Name() string
}

var sceneInst *SceneSystem

const SysNameScene = "scene"

var _ System = &SceneSystem{}

type SceneSystem struct {
	scenes map[string]Scene
	active []string
}

// Setup sets up the System.
func (s *SceneSystem) Setup() error {
	if sceneInst != nil {
		return ErrSystemInit(SysNameScene)
	}
	sceneInst = s

	return nil
}

// Teardown tears down the System.
func (s *SceneSystem) Teardown() {

}

// Name returns the name of the System.
func (s *SceneSystem) Name() string {
	return SysNameScene
}

func (s *SceneSystem) Register(scene Scene) error {
	if s.Registered(scene.Name()) {
		return fmt.Errorf("register scene: '%s' already registered", scene.Name())
	}

	s.scenes[scene.Name()] = scene

	logrus.Debug("Registered new scene: ", scene.Name())

	return nil
}

func (s *SceneSystem) Load(name string) error {
	if !s.Registered(name) {
		return fmt.Errorf("load scene: '%s' not registered", name)
	}

	if !s.scenes[name].Loaded() {
		return s.scenes[name].Load()
	}

	return nil
}

func (s *SceneSystem) PurgePush(name string) error {
	if !s.Registered(name) {
		return fmt.Errorf("purge push scene: '%s' not registered", name)
	}

	s.active = s.active[:0]
	s.Push(name)

	return nil
}

func (s *SceneSystem) Replace(name string) error {
	if !s.Registered(name) {
		return fmt.Errorf("replace error: '%s' not registered", name)
	}

	s.Pop()
	s.Push(name)

	return nil
}

func (s *SceneSystem) Push(name string) error {
	if !s.Registered(name) {
		return fmt.Errorf("push: '%s' not registered", name)
	}

	if err := s.Load(name); err != nil {
		return err
	}

	s.active = append(s.active, name)
	s.scenes[name].OnActivate()

	return nil
}

func (s *SceneSystem) Pop() string {
	if len(s.active) != 0 {
		last := s.active[len(s.active)-1]

		s.active = s.active[:len(s.active)-1]
		s.scenes[last].OnDeactivate()

		return last
	}

	return ""
}

func (s *SceneSystem) RemoveAll() {
	for key := range s.scenes {
		delete(s.scenes, key)
	}
	s.active = s.active[:0]
}

func (s *SceneSystem) Unregister(name string) error {
	if !s.Registered(name) {
		return fmt.Errorf("unregister scene: '%s' not registered", name)
	}

	delete(s.scenes, name)

	return nil
}

func (s *SceneSystem) Registered(name string) bool {
	_, ok := s.scenes[name]

	return ok
}

func (s *SceneSystem) Active() Scene {
	if s.ActiveCount() != 0 {
		name := s.active[len(s.active)-1]
		return s.scenes[name]
	}

	return nil
}

func (s *SceneSystem) ActiveName() string {
	if sc := s.Active(); sc != nil {
		return sc.Name()
	}

	return ""
}

func (s *SceneSystem) Count() int {
	return len(s.scenes)
}

func (s *SceneSystem) ActiveCount() int {
	return len(s.active)
}

func (s *SceneSystem) OnDisplay() {
	if sc := s.Active(); sc != nil {
		sc.Display()
	}
}

func (s *SceneSystem) OnUpdate() {
	if sc := s.Active(); sc != nil {
		sc.Update()
	}
}

func (s *SceneSystem) OnFixedUpdate() {
	if sc := s.Active(); sc != nil {
		sc.FixedUpdate()
	}
}

// NewSceneSystem creates a new scene system.
func NewSceneSystem() *SceneSystem {
	return &SceneSystem{
		scenes: make(map[string]Scene),
	}
}

// GetSceneSystem gets the scene system from the current app.
func GetSceneSystem() *SceneSystem {
	return sceneInst
}
