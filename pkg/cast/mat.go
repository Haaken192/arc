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

package cast

import "github.com/go-gl/mathgl/mgl32"

func ParseMat2(str string) (mgl32.Mat2, error) {
	v, err := ParseFloat32Slice(str, ",", 4)
	if err != nil {
		return mgl32.Mat2{}, err
	}

	return mgl32.Mat2{
		v[0], v[1],
		v[2], v[3],
	}, nil
}

func ParseMat3(str string) (mgl32.Mat3, error) {
	v, err := ParseFloat32Slice(str, ",", 9)
	if err != nil {
		return mgl32.Mat3{}, err
	}

	return mgl32.Mat3{
		v[0], v[1], v[2],
		v[3], v[4], v[5],
		v[6], v[7], v[8],
	}, nil
}

func ParseMat4(str string) (mgl32.Mat4, error) {
	v, err := ParseFloat32Slice(str, ",", 16)
	if err != nil {
		return mgl32.Mat4{}, err
	}

	return mgl32.Mat4{
		v[0], v[1], v[2], v[3],
		v[4], v[5], v[6], v[7],
		v[8], v[9], v[10], v[11],
		v[12], v[13], v[14], v[15],
	}, nil
}
