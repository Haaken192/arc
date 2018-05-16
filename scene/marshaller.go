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

package scene

import (
	"encoding/json"

	"github.com/haakenlabs/arc/core"
	"github.com/juju/errors"
)

type ErrUnmarshallerNotFound string
type ErrUnmarshallerExists string
type ErrMarshallerNotFound string
type ErrMarshallerExists string

func (e ErrUnmarshallerNotFound) Error() string {
	return "unmarshaller for type " + string(e) + " not found"
}

func (e ErrUnmarshallerExists) Error() string {
	return "unmarshaller for type " + string(e) + " already exists"
}

func (e ErrMarshallerNotFound) Error() string {
	return "marshaller for type " + string(e) + " not found"
}

func (e ErrMarshallerExists) Error() string {
	return "marshaller for type " + string(e) + " already exists"
}

type UnmarshalComponentFunc func([]byte) (Component, error)

var (
	componentUnmarshallers = make(map[string]UnmarshalComponentFunc)
)

type JSONComponent struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type JSONGameObject struct {
	Name       string            `json:"name"`
	Objects    []*JSONGameObject `json:"objects"`
	Components []*JSONComponent  `json:"components"`
}

type JSONScene struct {
	Name    string            `json:"name"`
	Objects []*JSONGameObject `json:"objects"`
}

func ComponentUnmarshaller(t string) (UnmarshalComponentFunc, error) {
	if fn, ok := componentUnmarshallers[t]; ok {
		return fn, nil
	}

	return nil, ErrMarshallerNotFound(t)
}

func AddComponentUnmarshaller(t string, fn UnmarshalComponentFunc) error {
	if _, dup := componentUnmarshallers[t]; dup {
		return ErrMarshallerExists(t)
	}

	componentUnmarshallers[t] = fn

	return nil
}

func BuildGameObject(data *JSONGameObject, parent *GameObject) (*GameObject, error) {
	o := NewGameObject(data.Name)
	o.parent = parent

	for _, v := range data.Objects {
		c, err := BuildGameObject(v, o)
		if err != nil {
			return nil, err
		}

		o.children = append(o.children, c)
	}

	for _, v := range data.Components {
		fn, err := ComponentUnmarshaller(v.Type)
		if err != nil {
			return nil, err
		}

		c, err := fn(v.Data)
		if err != nil {
			return nil, err
		}

		o.AddComponent(c)
	}

	return o, nil
}

func BuildScene(r *core.Resource) (*Scene, error) {
	data := &JSONScene{}
	if err := json.Unmarshal(r.Bytes(), data); err != nil {
		return nil, errors.Annotate(err, "json unmarshal error")
	}

	s := NewScene(data.Name)

	for _, v := range data.Objects {
		o, err := BuildGameObject(v, nil)
		if err != nil {
			return nil, errors.Annotate(err, "build object error")
		}

		if err := s.AddObject(o, nil); err != nil {
			return nil, errors.Annotate(err, "add object error")
		}
	}

	return s, nil
}
