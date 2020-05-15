package set

import (
	"politics/go/entry"
)

func (s *Set) SetEntries(es entry.Entries) {
	s.entries = es
}
