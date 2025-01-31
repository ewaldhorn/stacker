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

	if err.Error() != "empty Stack" {
		t.Errorf("[empty] expected 'empty Stack' error on pop, got %s", err.Error())
	}
}

// ----------------------------------------------------------------------------
func Test_Stacker_Smoke(t *testing.T) {
	myStack := NewStacker[int]()

	if myStack.Length() != 0 {
		t.Errorf("[smoke] Expected a length of 0, got %d", myStack.Length())
	}

	// add 1..5
	for i := range 5 {
		myStack.Push(i + 1)
	}

	result, err := myStack.Peek()

	if err != nil {
		t.Errorf("[smoke] didn't expect '%s' I tell you!", err.Error())
	}

	if result != 5 {
		t.Errorf("[smoke] I expected 5, got %d instead", result)
	}

	if myStack.Length() != 5 {
		t.Errorf("[smoke] Expected a length of 5, got %d", myStack.Length())
	}

}

// ----------------------------------------------------------------------------
func Test_Stacker_Large(t *testing.T) {

}
