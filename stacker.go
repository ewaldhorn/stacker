package stacker

import (
	"errors"
	"sync"
)

// ----------------------------------------------------------------------------
type Stacker[Type any] struct {
	mutex sync.Mutex
	Data  []Type
}

// ----------------------------------------------------------------------------
func NewStacker[Type any]() *Stacker[Type] {
	return &Stacker[Type]{mutex: sync.Mutex{}, Data: []Type{}}
}

// ----------------------------------------------------------------------------
// Returns the number of elements in the stack
func (stacker *Stacker[Type]) Length() int {
	stacker.mutex.Lock()
	defer stacker.mutex.Unlock()
	length := len(stacker.Data)

	return length
}

// ----------------------------------------------------------------------------
// Pushed a new item into the stack
func (stacker *Stacker[Type]) Push(element Type) {
	stacker.mutex.Lock()
	defer stacker.mutex.Unlock()
	stacker.Data = append(stacker.Data, element)
}

// ----------------------------------------------------------------------------
// Pops and returns the first item on the stack, possibly error if the stack is empty
func (stacker *Stacker[Type]) Pop() (Type, error) {
	stacker.mutex.Lock()
	defer stacker.mutex.Unlock()
	length := len(stacker.Data)

	if length == 0 {
		var empty Type
		return empty, errors.New("empty Stack")
	}
	element := stacker.Data[length-1]
	stacker.Data = stacker.Data[:length-1]

	return element, nil
}

// ----------------------------------------------------------------------------
// Returns the first entry in the stack, does not remove it, can error on empty
func (stacker *Stacker[Type]) Peek() (Type, error) {
	stacker.mutex.Lock()
	defer stacker.mutex.Unlock()
	length := len(stacker.Data)

	if length == 0 {
		var empty Type
		return empty, errors.New("empty Stack")
	}

	return stacker.Data[length-1], nil
}
