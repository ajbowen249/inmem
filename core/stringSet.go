package inmem

// StringSet is a container for a map[string]bool to leverage the
// O(1) average-case lookup time of a map, as Go does not have a
// true set type in its standard library.
type StringSet struct {
	container map[string]bool
}

// NewStringSet creates a StringSet from a collection of values.
func NewStringSet(values []string) StringSet {
	set := StringSet{make(map[string]bool)}

	for _, value := range values {
		set.container[value] = false
	}

	return set
}

// Contains returns true if an item exists within the set.
func (s *StringSet) Contains(item string) bool {
	_, exists := (*s).container[item]
	return exists
}
