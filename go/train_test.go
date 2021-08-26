package main

import (
	"testing"
)

func TestPassengerTrain(t *testing.T) {
	var train = &Train{"HPP"}
	expected := "<HHHH::|OOOO|::|OOOO|"
	actual := train.print()
	if expected != actual {
		t.Errorf("Expected: %s Actual: %s", expected, actual)
	}
}

func TestRestaurantTrain(t *testing.T) {
	var train = &Train{"HPRP"}
	expected := "<HHHH::|OOOO|::|hThT|::|OOOO|"
	actual := train.print()
	if expected != actual {
		t.Errorf("Expected: %s Actual: %s", expected, actual)
	}
}

func TestDoubleHeadedTrain(t *testing.T) {
	var train = &Train{"HPRPH"}
	expected := "<HHHH::|OOOO|::|hThT|::|OOOO|::HHHH>"
	actual := train.print()
	if expected != actual {
		t.Errorf("Expected: %s Actual: %s", expected, actual)
	}
}

func TestDetachment(t *testing.T) {
	var train = &Train{"HPRPH"}

	train.detachEnd()
	expectedPostEndDetach := "<HHHH::|OOOO|::|hThT|::|OOOO|::HHHH>"
	actualPostEndDetach := train.print()
	if expectedPostEndDetach != actualPostEndDetach {
		t.Errorf("Expected: %s Actual: %s", expectedPostEndDetach, actualPostEndDetach)
	}

	train.detachHead()
	expectedPostHeadDetach := "|OOOO|::|hThT|::|OOOO|"
	actualPostHeadDetach := train.print()
	if expectedPostHeadDetach != actualPostHeadDetach {
		t.Errorf("Expected: %s Actual: %s", expectedPostHeadDetach, actualPostHeadDetach)
	}
}

func TestErrorOnFullCargoTrain(t *testing.T) {
	var train = &Train{"HPRPH"}
	expectedInitial := "<HHHH::|____|::|____|::|____|"
	actualInitial := train.print()
	if expectedInitial != actualInitial {
		t.Errorf("Expected: %s Actual: %s", expectedInitial, actualInitial)
	}

	train.fill()
	expectedFirstFill := "<HHHH::|^^^^|::|____|::|____|"
	actualFirstFill := train.print()
	if expectedFirstFill != actualFirstFill {
		t.Errorf("Expected: %s Actual: %s", expectedFirstFill, actualFirstFill)
	}

	train.fill()
	expectedSecondFill := "<HHHH::|^^^^|::|^^^^|::|____|"
	actualSecondFill := train.print()
	if expectedSecondFill != actualSecondFill {
		t.Errorf("Expected: %s Actual: %s", expectedSecondFill, actualSecondFill)
	}

	train.fill()
	expectedThirdFill := "<HHHH::|^^^^|::|^^^^|::|^^^^|"
	actualThirdFill := train.print()
	if expectedThirdFill != actualThirdFill {
		t.Errorf("Expected: %s Actual: %s", expectedThirdFill, actualThirdFill)
	}

	_, err := train.fill()
	if err == nil {
		t.Errorf("Filling a train that is already full should return an error")
	}
}

func TestErrorOnFullMixedTrain(t *testing.T) {
	var train = &Train{"HPCPC"}
	expectedInitial := "<HHHH::|OOOO|::|____|::|OOOO|::|____|"
	actualInitial := train.print()
	if expectedInitial != actualInitial {
		t.Errorf("Expected: %s Actual: %s", expectedInitial, actualInitial)
	}

	train.fill()
	expectedFirstFill := "<HHHH::|OOOO|::|^^^^|::|OOOO|::|____|"
	actualFirstFill := train.print()
	if expectedFirstFill != actualFirstFill {
		t.Errorf("Expected: %s Actual: %s", expectedFirstFill, actualFirstFill)
	}

	train.fill()
	expectedSecondFill := "<HHHH::|OOOO|::|^^^^|::|OOOO|::|^^^^|"
	actualSecondFill := train.print()
	if expectedSecondFill != actualSecondFill {
		t.Errorf("Expected: %s Actual: %s", expectedSecondFill, actualSecondFill)
	}

	_, err := train.fill()
	if err == nil {
		t.Errorf("Filling a train that is already full should return an error")
	}
}
