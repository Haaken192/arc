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

import (
	"reflect"
	"testing"
)

func TestParseFloat32Slice(t *testing.T) {
	var tests = []struct {
		in      string
		sep     string
		count   int
		want    []float32
		wantErr bool
	}{
		{in: "", sep: "", want: nil, wantErr: false},
		{in: "1.1", sep: ",", want: []float32{1.1}, wantErr: false},
		{in: "1,2.25", sep: ",", want: []float32{1, 2.25}, wantErr: false},
		{in: "1.25,2.5", sep: ",", count: 2, want: []float32{1.25, 2.5}, wantErr: false},
		{in: "1.25,2.5,32.4", sep: ",", count: 2, wantErr: true},
		{in: "1.25,2.53,false", sep: ",", wantErr: true},
	}

	for i, v := range tests {
		got, err := ParseFloat32Slice(v.in, v.sep, v.count)
		if (err != nil) != v.wantErr {
			t.Errorf("%s failed test case %d. err: %v wantErr: %v", t.Name(), i, err, v.wantErr)
		} else if !v.wantErr {
			if !reflect.DeepEqual(v.want, got) {
				t.Errorf("%s case %d value mismatch. want: %v got: %v", t.Name(), i, v.want, got)
			}
		}
	}
}

func TestParseInt32Slice(t *testing.T) {
	var tests = []struct {
		in      string
		sep     string
		count   int
		want    []int32
		wantErr bool
	}{
		{in: "", sep: "", want: nil, wantErr: false},
		{in: "1", sep: ",", want: []int32{1}, wantErr: false},
		{in: "1,2", sep: ",", want: []int32{1, 2}, wantErr: false},
		{in: "1,2", sep: ",", count: 2, want: []int32{1, 2}, wantErr: false},
		{in: "1,2,3", sep: ",", count: 2, wantErr: true},
		{in: "1,2,false", sep: ",", wantErr: true},
	}

	for i, v := range tests {
		got, err := ParseInt32Slice(v.in, v.sep, v.count)
		if (err != nil) != v.wantErr {
			t.Errorf("%s failed test case %d. err: %v wantErr: %v", t.Name(), i, err, v.wantErr)
		} else if !v.wantErr {
			if !reflect.DeepEqual(v.want, got) {
				t.Errorf("%s case %d value mismatch. want: %v got: %v", t.Name(), i, v.want, got)
			}
		}
	}
}

func TestParseUint32Slice(t *testing.T) {
	var tests = []struct {
		in      string
		sep     string
		count   int
		want    []uint32
		wantErr bool
	}{
		{in: "", sep: "", want: nil, wantErr: false},
		{in: "1", sep: ",", want: []uint32{1}, wantErr: false},
		{in: "1,2", sep: ",", want: []uint32{1, 2}, wantErr: false},
		{in: "1,2", sep: ",", count: 2, want: []uint32{1, 2}, wantErr: false},
		{in: "1,2,3", sep: ",", count: 2, wantErr: true},
		{in: "1,2,false", sep: ",", wantErr: true},
	}

	for i, v := range tests {
		got, err := ParseUint32Slice(v.in, v.sep, v.count)
		if (err != nil) != v.wantErr {
			t.Errorf("%s failed test case %d. err: %v wantErr: %v", t.Name(), i, err, v.wantErr)
		} else if !v.wantErr {
			if !reflect.DeepEqual(v.want, got) {
				t.Errorf("%s case %d value mismatch. want: %v got: %v", t.Name(), i, v.want, got)
			}
		}
	}
}

func TestParseVec2(t *testing.T) {

}

func TestParseVec3(t *testing.T) {

}

func TestParseVec4(t *testing.T) {

}
