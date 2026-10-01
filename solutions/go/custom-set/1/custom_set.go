package customset

import (
	"sort"
	"strconv"
	"strings"
)
// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

type Set map[string]struct{}

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, v := range l {
		set[v] = struct{}{}
	}
	return set
}

func (s Set) String() string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	parts := make([]string, len(keys))
	for i, k := range keys {
		parts[i] = strconv.Quote(k)
	}

	return "{" + strings.Join(parts, ", ") + "}"
}

func (s Set) IsEmpty() bool {
    return len(s) == 0
}

func (s Set) Has(elem string) bool {
	if _, ok := s[elem]; ok {
        return true
    }
    return false
}

func (s Set) Add(elem string) {
    s[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
    for key := range s1 {
        if !s2.Has(key) {
            return false
        }
    }
    return true
}

func Disjoint(s1, s2 Set) bool {
    for key := range s1 {
        if s2.Has(key) {
            return false
        }
    }
    return true
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func Intersection(s1, s2 Set) Set {
	result := New()
	for k := range s1 {
		if s2.Has(k) {
			result.Add(k)
		}
	}
	return result
}

func Difference(s1, s2 Set) Set {
	result := New()
	for k := range s1 {
		if !s2.Has(k) {
			result.Add(k)
		}
	}
	return result
}

func Union(s1, s2 Set) Set {
	result := New()
	for k := range s1 {
		result.Add(k)
	}
	for k := range s2 {
		result.Add(k)
	}
	return result
}
