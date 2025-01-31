package stacker

import "testing"

// ----------------------------------------------------------------------------
func Test_Stacker_Empty(t *testing.T) {
	myStack := NewStacker[bool]()

	_, err := myStack.Peek()

	if err == nil {
		t.Errorf("[empty] expected an error on peek, got nil")
	}

	_, err = myStack.Pop()
	if err == nil {
		t.Errorf("[empty] expected an error on pop, got nil")
	}

}

// ----------------------------------------------------------------------------
func Test_Stacker_Smoke(t *testing.T) {

}

// ----------------------------------------------------------------------------
func Test_Stacker_Large(t *testing.T) {

}
