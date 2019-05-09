// Copyright (c) 2019 Coinbase, Inc. See LICENSE

package util

// StringSet - an unordered collection of unique strings
type StringSet map[string]struct{}

// NewStringSet - initialize a new StringSet
func NewStringSet() StringSet {
	return StringSet{}
}

// Add - add value to the set
func (s StringSet) Add(value string) {
	s[value] = struct{}{}
}

// Remove - remove value from the set
func (s StringSet) Remove(value string) {
	delete(s, value)
}

// Contains - checks whether the set contains a given value
func (s StringSet) Contains(value string) bool {
	_, ok := s[value]
	return ok
}