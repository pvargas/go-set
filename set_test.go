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
		t.Errorf("\nResult %+v\nExpected %+v", actualLength, expectedLength)
	}
}

func TestNewSetPopulated(t *testing.T) {
	expectedMembers := []int{12, 8, 6, 4}
	expectedLength := 4
	testSet := set.NewSet(expectedMembers...)
	actualLength := len(testSet)

	if expectedLength != actualLength {
		t.Errorf("\nActual length: %+v\nExpected length: %+v", actualLength, expectedLength)
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
		t.Errorf("\nActual length: %+v\nExpected length: %+v", actualLength, expectedLength)
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
		t.Errorf("\nActual length: %+v\nExpected length: %+v", actualLength, expectedLength)
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
		t.Errorf("\nActual length: %+v\nExpected length: %+v", actualLength, expectedLength)
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
		t.Errorf("\nActual length: %+v\nExpected length: %+v", actualLength, expectedLength)
	}
}

func TestUnion(t *testing.T) {
	setA := set.NewSet(0, 1, 3, 5)
	setB := set.NewSet(0, 2, 4, 6)
	expected := set.NewSet(0, 1, 2, 3, 4, 5, 6)

	actualA := setA.Union(setB)

	if !reflect.DeepEqual(actualA, expected) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", actualA, expected)
	}

	actualB := setB.Union(setA)

	if !reflect.DeepEqual(actualB, expected) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", actualB, expected)
	}

	// the original sets should not have been modified
	expectedSetA := set.NewSet(0, 1, 3, 5)
	expectedSetB := set.NewSet(0, 2, 4, 6)
	if !reflect.DeepEqual(setA, expectedSetA) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", setA, expectedSetA)
	}
	if !reflect.DeepEqual(setB, expectedSetB) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", setB, expectedSetB)
	}
}

func TestDifference(t *testing.T) {
	setA := set.NewSet(1, 2, 3, 4)
	setB := set.NewSet(3, 4, 5, 6)
	expectedAB := set.NewSet(1, 2)

	actualAB := setA.Difference(setB)

	if !reflect.DeepEqual(actualAB, expectedAB) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", actualAB, expectedAB)
	}

	expectedBA := set.NewSet(5, 6)
	actualBA := setB.Difference(setA)

	if !reflect.DeepEqual(actualBA, expectedBA) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", actualBA, expectedBA)
	}
}

func TestSymmetricDifference(t *testing.T) {
	setA := set.NewSet(1, 2, 3, 4)
	setB := set.NewSet(0, 2, 3, 5)
	expected := set.NewSet(0, 1, 4, 5)

	actualA := setA.SymmetricDifference(setB)

	if !reflect.DeepEqual(actualA, expected) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", actualA, expected)
	}

	actualB := setB.SymmetricDifference(setA)

	if !reflect.DeepEqual(actualB, expected) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", actualB, expected)
	}

	setC := set.NewSet(9, 28, 44)
	actualC := setA.SymmetricDifference(setC)
	expectedC := set.NewSet(1, 2, 3, 4, 9, 28, 44)

	if !reflect.DeepEqual(actualC, expectedC) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", actualC, expectedC)
	}
}

func TestIntersection(t *testing.T) {
	setA := set.NewSet(1, 2, 3, 4)
	setB := set.NewSet(3, 4, 5, 6)
	expected := set.NewSet(3, 4)

	actualAB := setA.Intersection(setB)

	if !reflect.DeepEqual(actualAB, expected) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", actualAB, expected)
	}

	actualBA := setB.Intersection(setA)

	if !reflect.DeepEqual(actualBA, expected) {
		t.Errorf("\nActual set: %+v\nExpected set: %+v", actualBA, expected)
	}
}

func TestIsSubset(t *testing.T) {
	setA := set.NewSet("C", "D", "Eb", "F", "G", "Ab", "Bb")
	setB := set.NewSet("C", "Eb", "G")
	setC := set.NewSet("C", "Eb", "G", "F#")
	expectedA := false
	expectedB := true
	expectedC := false

	actualA := setA.IsSubset(setB)

	if !reflect.DeepEqual(actualA, expectedA) {
		t.Errorf("\nActual: %+v\nExpected: %+v", actualA, expectedA)
	}

	actualB := setB.IsSubset(setA)

	if actualB != expectedB {
		t.Errorf("\nActual: %+v\nExpected: %+v", actualB, expectedB)
	}

	actualC := setC.IsSubset(setA)

	if actualC != expectedC {
		t.Errorf("\nActual: %+v\nExpected: %+v", actualC, expectedC)
	}
}

func TestIsProperSubset(t *testing.T) {
	setA := set.NewSet("C", "Bb", "G", "Eb")
	setB := set.NewSet("C", "Bb", "G", "Eb")
	setC := set.NewSet("C", "Eb", "G")
	expectedA := false
	expectedC := true

	actualA := setA.IsProperSubset(setB)

	if actualA != expectedA {
		t.Errorf("\nActual: %+v\nExpected: %+v", actualA, expectedA)
	}

	actualC := setC.IsProperSubset(setA)

	if actualC != expectedC {
		t.Errorf("\nActual: %+v\nExpected: %+v", actualC, expectedC)
	}
}

func TestIsDisjoint(t *testing.T) {
	setA := set.NewSet("Eb", "G", "Bb")
	setB := set.NewSet("D", "F#", "A")
	setC := set.NewSet("Bb", "D", "F")

	expected := true
	expectedC := false

	actualA := setA.IsDisjoint(setB)

	if actualA != expected {
		t.Errorf("\nActual: %+v\nExpected: %+v", actualA, expected)
	}

	actualB := setB.IsDisjoint(setA)

	if actualB != expected {
		t.Errorf("\nActual: %+v\nExpected: %+v", actualB, expected)
	}

	actualC := setC.IsDisjoint(setA)

	if actualC != expectedC {
		t.Errorf("\nActual: %+v\nExpected: %+v", actualC, expectedC)
	}
}
