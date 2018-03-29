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
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/haakenlabs/arc/pkg/math"
)

var _ System = &AudioSystem{}

var audioInst *AudioSystem

const SysNameAudio = "audio"

type AudioChannel uint8

type AudioSystem struct {
	volume     float64
	channels   AudioChannel
	sampleRate beep.SampleRate
	mute       bool
}

// Setup sets up the System.
func (s *AudioSystem) Setup() error {
	if timeInst != nil {
		return ErrSystemInit(SysNameAudio)
	}
	audioInst = s

	speaker.Init(s.sampleRate, s.sampleRate.N(time.Second/10))

	return nil
}

// Teardown tears down the System.
func (s *AudioSystem) Teardown() {
	audioInst = nil
}

// Name returns the name of the System.
func (s *AudioSystem) Name() string {
	return SysNameAudio
}

func (s *AudioSystem) Volume() float64 {
	return s.volume
}

func (s *AudioSystem) SetVolume(volume float64) {
	s.volume = math.Clamp(volume, 0.0, 1.0)
}

func (s *AudioSystem) Mute() {
	s.SetMute(true)
}

func (s *AudioSystem) Unmute() {
	s.SetMute(false)
}

func (s *AudioSystem) SetMute(mute bool) {
	s.mute = mute
}

func (s *AudioSystem) PlaySound(sound *Sound) {
	speaker.Play(sound.streamer)
}

func NewAudioSystem(rate beep.SampleRate) *AudioSystem {
	return &AudioSystem{
		sampleRate: rate,
	}
}

// GetTime gets the time system from the current app.
func GetAudioSystem() *AudioSystem {
	return audioInst
}
