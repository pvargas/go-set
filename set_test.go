package set_test

import (
	"reflect"
	"testing"

	"github.com/pvargas/go-set"
)

func TestNewSetEmpty(t *testing.T) {
	expectedLength := 0
	testSet := set.NewSet[string]()
	actualLength := len(testSet)

	if expectedLength != actualLength {
		t.Errorf("\nResult %+v\nExpected %+v",
			actualLength, expectedLength)
	}
}

func TestNewSetPopulated(t *testing.T) {
	expectedMembers := []int{12, 8, 6, 4}
	expectedLength := 4
	testSet := set.NewSet(expectedMembers...)
	actualLength := len(testSet)

	if expectedLength != actualLength {
		t.Errorf("\nActual length: %+v\nExpected length: %+v",
			actualLength, expectedLength)
	}

	for _, element := range expectedMembers {
		_, ok := testSet[element]

		if !ok {
			t.Errorf("\nElement %+v not in set", element)
		}
	}
}

func TestNewSetUnique(t *testing.T) {
	expectedMembers := []string{"Bruce Wayne", "Peter Parker", "Peter Parker"}
	expectedLength := 2
	testSet := set.NewSet(expectedMembers...)
	actualLength := len(testSet)

	if expectedLength != actualLength {
		t.Errorf("\nActual length: %+v\nExpected length: %+v",
			actualLength, expectedLength)
	}
}

func TestContains(t *testing.T) {
	tempElement := "007"
	expectedMembers := []string{"James", "Bond", tempElement}
	testSet := set.NewSet(expectedMembers...)

	for _, element := range expectedMembers {
		if !testSet.Contains(element) {
			t.Errorf("\nElement %+v not in set", element)
		}
	}

	delete(testSet, tempElement)

	if testSet.Contains(tempElement) {
		t.Errorf("\nUnexpected element found in set:  %+v ", tempElement)
	}

	if !testSet.Contains(expectedMembers[0]) {
		t.Errorf("\nExpected element not found in set:  %+v ", expectedMembers[0])
	}

	if !testSet.Contains(expectedMembers[1]) {
		t.Errorf("\nExpected element not found in set:  %+v ", expectedMembers[1])
	}
}

func TestElementsPopulated(t *testing.T) {
	expectedMembers := []string{"Frodo", "Sam", "Pippin", "Merry"}
	testSet := set.NewSet(expectedMembers...)

	for _, element := range expectedMembers {
		if !testSet.Contains(element) {
			t.Errorf("\nExpected element not in set: %+v", element)
		}
	}
}

func TestElementsEmpty(t *testing.T) {
	emptySet := set.NewSet[rune]()
	expectedLength := 0
	actualLength := len(emptySet.Elements())

	if expectedLength != actualLength {
		t.Errorf("\nActual length: %+v\nExpected length: %+v",
			actualLength, expectedLength)
	}
}

func TestInsert(t *testing.T) {
	expectedMembers := []int{1, 2, 3, 4}
	testSet := set.NewSet[int]()

	for _, element := range expectedMembers {
		testSet.Insert(element)
	}

	for _, element := range expectedMembers {
		if !testSet.Contains(element) {
			t.Errorf("\nExpected element not in set: %+v", element)
		}
	}
}

func TestRemovePresent(t *testing.T) {
	expectedMembers := []string{"Ithkuil", "Toki Pona", "Quenya"}
	testSet := set.NewSet(expectedMembers...)
	expectedLength := 0

	for _, element := range expectedMembers {
		testSet.Remove(element)
	}

	actualLength := len(testSet)

	if expectedLength != actualLength {
		t.Errorf("\nActual length: %+v\nExpected length: %+v",
			actualLength, expectedLength)
	}
}

func TestRemoveMissing(t *testing.T) {
	missingElement := "Esperanto"
	expectedMembers := []string{"Ithkuil", "Toki Pona", "Quenya"}
	testSet := set.NewSet(expectedMembers...)
	expectedLength := len(expectedMembers)

	testSet.Remove(missingElement)

	actualLength := len(testSet)

	if expectedLength != actualLength {
		t.Errorf("\nActual length: %+v\nExpected length: %+v",
			actualLength, expectedLength)
	}
}

func TestUnion(t *testing.T) {
	setA := set.NewSet(1, 3, 5)
	setB := set.NewSet(2, 4, 6)
	expected := set.NewSet(1, 2, 3, 4, 5, 6)

	actualA := setA.Union(setB)

	if !reflect.DeepEqual(actualA, expected) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v",
			actualA, expected)
	}

	actualB := setB.Union(setA)

	if !reflect.DeepEqual(actualB, expected) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v",
			actualB, expected)
	}

	// the original sets should not be modified
	expectedSetA := set.NewSet(1, 3, 5)
	expectedSetB := set.NewSet(2, 4, 6)
	if !reflect.DeepEqual(setA, expectedSetA) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v",
			setA, expectedSetA)
	}
	if !reflect.DeepEqual(setB, expectedSetB) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v",
			setB, expectedSetB)
	}
}
