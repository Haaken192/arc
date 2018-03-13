package sg

import (
	"errors"
	"fmt"
)

// ErrDescriptorLimit describes an error in which a descriptor is invalid.
type ErrDescriptorInvalid Descriptor

func (e ErrDescriptorInvalid) Error() string {
	return fmt.Sprintf("descriptor invalid: %d", e)
}

// ErrDescriptorLimit describes an error in which a descriptor cannot be found.
type ErrDescriptorNotFound Descriptor

func (e ErrDescriptorNotFound) Error() string {
	return fmt.Sprintf("descriptor not found: %d", e)
}

// ErrDescriptorLimit describes an error in which a descriptor already exists.
type ErrDescriptorExists Descriptor

func (e ErrDescriptorExists) Error() string {
	return fmt.Sprintf("descriptor already exists: %d", e)
}

// ErrDescriptorLimit describes an error in which the number descriptors
// has exceeded a limit.
var ErrDescriptorLimit = errors.New("descriptor limit reached")

// ErrDescriptorLimit describes an error in which a descriptor has no parent.
type ErrDescriptorNoParent Descriptor

func (e ErrDescriptorNoParent) Error() string {
	return fmt.Sprintf("descriptor has no parent: %d", e)
}

type ErrDescriptorDescendant struct {
	d Descriptor
	p Descriptor
}

func (e ErrDescriptorDescendant) Error() string {
	return fmt.Sprintf("parent descriptor %d is a descendant of %d", e.p, e.d)
}

type ErrEdgeExists struct {
	d Descriptor
	p Descriptor
}

func (e ErrEdgeExists) Error() string {
	return fmt.Sprintf("edge already exists: %d->%d ", e.p, e.d)
}
