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

package prefabs

import (
	"github.com/haakenlabs/arc/scene"
	"github.com/haakenlabs/arc/system/asset/texture"
	"github.com/haakenlabs/arc/system/window"
	"github.com/haakenlabs/arc/ui"
	"github.com/haakenlabs/arc/ui/widget"
)

func CreateSplash() *scene.GameObject {
	c := ui.CreateController("splash")

	bg := widget.CreatePanel("splash-bg")
	ui.RectTransformComponent(bg).SetSize(window.Resolution().Vec2())
	ui.RectTransformComponent(bg).SetAnchorPreset(ui.StretchAnchorAll)
	widget.ImageComponent(bg).SetColor(ui.Styles.BackgroundColor)

	logo := widget.CreateImage("splash-logo")
	ui.RectTransformComponent(logo).SetPresets(ui.AnchorMiddleCenter, ui.PivotMiddleCenter)
	widget.ImageComponent(logo).SetTexture(texture.MustGet("arc-logo.png"))

	bg.AddChild(logo)
	c.AddChild(bg)

	return c
}
