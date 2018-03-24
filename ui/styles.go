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
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/haakenlabs/forge"
)

type StyleSet struct {
	BackgroundColor     forge.Color `json:"background_color"`
	TextColor           forge.Color `json:"text_color"`
	TextColorPrimary    forge.Color `json:"text_color_primary"`
	TextColorActive     forge.Color `json:"text_color_active"`
	TextColorDisabled   forge.Color `json:"text_color_disabled"`
	WidgetColor         forge.Color `json:"widget_color"`
	WidgetColorPrimary  forge.Color `json:"widget_color_primary"`
	WidgetColorActive   forge.Color `json:"widget_color_active"`
	WidgetColorDisabled forge.Color `json:"widget_color_disabled"`
	TextSize            int32       `json:"text_size"`
}

var Styles = StyleSet{
	BackgroundColor:     forge.Color{0.1, 0.1, 0.1, 0.9},
	TextColor:           forge.Color{1.0, 1.0, 1.0, 0.9},
	TextColorActive:     forge.Color{1.0, 1.0, 1.0, 0.9},
	TextColorPrimary:    forge.Color{0.0, 0.27, 0.68, 0.9},
	TextColorDisabled:   forge.Color{0.5, 0.5, 0.5, 0.5},
	WidgetColor:         forge.Color{0.15, 0.15, 0.15, 0.9},
	WidgetColorPrimary:  forge.Color{0.0, 0.27, 0.68, 0.9},
	WidgetColorActive:   forge.Color{0.17, 0.17, 0.17, 1.0},
	WidgetColorDisabled: forge.Color{0.1, 0.1, 0.1, 0.5},
	TextSize:            11,
}

func LoadStyle(r io.Reader) error {
	var s StyleSet

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	Styles = s

	return nil
}
