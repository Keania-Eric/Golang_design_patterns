package singleton

import "testing"

func TestGetInstance(t *testing.T) {

	counter1 := GetInstance()

	if counter1 == nil {
		// Test of acceptance criteria 1 failed
		t.Error("Expected pointer to be singleton after calling GetInstance, not nil")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Errorf("After calling for the first time count should be 1 but it was %d", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		// Test 2 failed
		t.Error("Expected same counter for instance 2 but got a different one")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling addone using the second counter the result should be 2 but we got %d", currentCount)
	}
}
