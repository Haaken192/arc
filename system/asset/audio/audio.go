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

package audio

import (
	"path/filepath"
	"sync"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"

	"fmt"

	"github.com/haakenlabs/arc/core"
	"github.com/haakenlabs/arc/system/asset"
)

const AssetNameAudio = "audio"

var _ core.AssetHandler = &Handler{}

type Handler struct {
	core.BaseAssetHandler
}

func (h *Handler) Load(r *core.Resource) error {
	var streamer beep.Streamer
	var format beep.Format
	var err error

	name := r.Base()
	ext := filepath.Ext(name)

	if _, dup := h.Items[name]; dup {
		return core.ErrAssetExists(name)
	}

	switch ext {
	case "mp3":
		streamer, format, err = mp3.Decode(r.ReadCloser())
	case "wav":
		streamer, format, err = wav.Decode(r.ReadCloser())
	case "flac":
		streamer, format, err = flac.Decode(r.ReadCloser())
	default:
		return fmt.Errorf("unknown audio type: %s", ext)
	}

	if err != nil {
		return err
	}

	s := core.NewSound(streamer, format)
	s.SetName(name)

	return h.Add(name, s)
}

func (h *Handler) Add(name string, sound *core.Sound) error {
	if _, dup := h.Items[name]; dup {
		return core.ErrAssetExists(name)
	}

	if err := sound.Alloc(); err != nil {
		return err
	}

	h.Items[name] = sound.ID()

	return nil
}

func (h *Handler) Get(name string) (*core.Sound, error) {
	a, err := h.GetAsset(name)
	if err != nil {
		return nil, err
	}

	a2, ok := a.(*core.Sound)
	if !ok {
		return nil, core.ErrAssetType(name)
	}

	return a2, nil
}

func (h *Handler) MustGet(name string) *core.Sound {
	a, err := h.Get(name)
	if err != nil {
		panic(err)
	}

	return a
}

func (h *Handler) Name() string {
	return AssetNameAudio
}

func NewHandler() *Handler {
	h := &Handler{}
	h.Items = make(map[string]int32)
	h.Mu = &sync.RWMutex{}

	return h
}

func Get(name string) (*core.Sound, error) {
	return mustHandler().Get(name)
}

func MustGet(name string) *core.Sound {
	return mustHandler().MustGet(name)
}

func mustHandler() *Handler {
	h, err := asset.GetHandler(AssetNameAudio)
	if err != nil {
		panic(err)
	}

	return h.(*Handler)
}
