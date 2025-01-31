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
func (stack *Stacker[Type]) Push(element Type) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.Data = append(stack.Data, element)
}

// ----------------------------------------------------------------------------
func (stack *Stacker[Type]) Pop() (Type, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	l := len(stack.Data)

	if l == 0 {
		var empty Type
		return empty, errors.New("empty Stack")
	}
	element := stack.Data[l-1]
	stack.Data = stack.Data[:l-1]

	return element, nil
}

// ----------------------------------------------------------------------------
func (stack *Stacker[Type]) Peek() (Type, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	l := len(stack.Data)

	if l == 0 {
		var empty Type
		return empty, errors.New("empty Stack")
	}

	return stack.Data[l-1], nil
}
