package inmem

import (
	"testing"
)

func TestStringSet(t *testing.T) {
	tolkeinSelection := []string{
		"The Hobbit",
		"The Adventures Of Tom Bombadil",
		"The Silmarillion",
	}

	cases := map[string]bool{
		"The Hobbit":                            true,
		"The Silmarillion":                      true,
		"The Lion, the Witch, and the Wardrobe": false,
	}

	set := NewStringSet(tolkeinSelection)

	runStringSetCases(cases, &set, t)
}

func TestStringSetEmpty(t *testing.T) {
	cases := map[string]bool{
		"Something":      false,
		"Something else": false,
	}

	set := NewStringSet(make([]string, 0))

	runStringSetCases(cases, &set, t)
}

func runStringSetCases(cases map[string]bool, set *StringSet, t *testing.T) {
	for input, expected := range cases {
		actual := set.Contains(input)

		if actual != expected {
			t.Errorf("Expected set.Contains(%v) == %v, but was %v.", input, expected, actual)
		}
	}
}
