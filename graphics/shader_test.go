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

package graphics

import (
	"reflect"
	"testing"

	"github.com/go-gl/mathgl/mgl32"
)

func TestShaderProperty_UnmarshalJSON(t *testing.T) {
	var tests = []struct {
		in      []byte
		want    interface{}
		wantErr bool
	}{
		{in: []byte("bool:false"), want: false},
		{in: []byte("float:34.1256"), want: float32(34.1256)},
		{in: []byte("int:1024"), want: int32(1024)},
		{in: []byte("vec2:25.3,342.1"), want: mgl32.Vec2{25.3, 342.1}},
		{in: []byte("vec3:255,55.2,0.855"), want: mgl32.Vec3{255, 55.2, 0.855}},
		{in: []byte("vec4:0,1,2,3"), want: mgl32.Vec4{0, 1, 2, 3}},
		{in: []byte("mat2:0.1,1.1,2.1,3.1"), want: mgl32.Mat2{0.1, 1.1, 2.1, 3.1}},
		{in: []byte("mat3:0,1,2,3,4,5,6,7,8"), want: mgl32.Mat3{0, 1, 2, 3, 4, 5, 6, 7, 8}},
		{in: []byte("mat4:0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15"), want: mgl32.Mat4{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
		{in: []byte("vec:0,1,2,3"), wantErr: true},
	}

	var p ShaderProperty

	for i, v := range tests {
		err := p.UnmarshalJSON(v.in)
		if (err != nil) != v.wantErr {
			t.Errorf("%s failed test case %d. err: %v wantErr: %v", t.Name(), i, err, v.wantErr)
		} else if !v.wantErr {
			if !reflect.DeepEqual(p.Value, v.want) {
				t.Errorf("%s case %d value mismatch. want: %v got: %v", t.Name(), i, v.want, p.Value)
			}
		}
	}
}
