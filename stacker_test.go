package stacker

import (
	"fmt"
	"testing"
)

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

	// just peek for now
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

	// now pop one
	result, err = myStack.Pop()

	if err != nil {
		t.Errorf("[smoke] didn't expect '%s' I tell you!", err.Error())
	}

	if result != 5 {
		t.Errorf("[smoke] I expected 5, got %d instead", result)
	}

	if myStack.Length() != 4 {
		t.Errorf("[smoke] Expected a length of 4, got %d", myStack.Length())
	}

}

// ----------------------------------------------------------------------------
func Test_Stacker_Large(t *testing.T) {
	myStack := NewStacker[string]()

	if myStack.Length() != 0 {
		t.Errorf("[large] Expected a length of 0, got %d", myStack.Length())
	}

	// add 1..1000
	for i := range 1000 {
		myStack.Push(fmt.Sprintf("Sentence %d", i+1))
	}

	if myStack.Length() != 1000 {
		t.Errorf("[large] Expected a length of 1000, got %d", myStack.Length())
	}

	// now remove 100 items
	for range 100 {
		_, _ = myStack.Pop()
	}

	if myStack.Length() != 900 {
		t.Errorf("[large] Expected a length of 900, got %d", myStack.Length())
	}

	msg, err := myStack.Peek()

	if err != nil {
		t.Errorf("[large] failed with '%s', expected no error", err.Error())
	}

	if msg != "Sentence 900" {
		t.Errorf("[large] Wanted 'Sentence 900', got '%s'", msg)
	}

	// do a number of peeks
	for range 10000 {
		_, _ = myStack.Peek()
	}

	if myStack.Length() != 900 {
		t.Errorf("[large] Expected a length of 900, got %d (peeks)", myStack.Length())
	}

	for myStack.Length() > 0 {
		_, err = myStack.Pop()

		if err != nil {
			t.Errorf("[large] error popping all '%s'", err.Error())
		}
	}

	if myStack.Length() != 0 {
		t.Errorf("[large] stack should be empty now, it contains %d elements", myStack.Length())
	}
}
