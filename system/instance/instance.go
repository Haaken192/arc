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

package instance

import (
	"github.com/haakenlabs/arc/core"
)

// Assign registers an object that conforms to Object with the instance database.
func Assign(o core.Object) error {
	return core.GetInstanceSystem().Assign(o)
}

// MustAssign is like Assign, but panics if an error is encountered.
func MustAssign(o core.Object) {
	core.GetInstanceSystem().MustAssign(o)
}

// Release releases an object with given ID from the instance database. The ID
// will be freed and available for reuse.
func Release(id ...int32) {
	core.GetInstanceSystem().Release(id...)
}

// ReleaseAll releases all objects in the instance database.
func ReleaseAll() {
	core.GetInstanceSystem().ReleaseAll()
}

// Get retrieves an object with given from the instance database.
func Get(id int32) (core.Object, error) {
	return core.GetInstanceSystem().Get(id)
}
