package tree

import (
	"fmt"
	"politics/go/entry"
	"politics/go/entry/helper"
)

func (t *Tree) LookupTreeHash(hash string) (*Tree, error) {
	id, err := helper.ParseHash(hash)
	if err != nil {
		return nil, fmt.Errorf("LookupTreeHash: Couldn’t parse hash %v.", err)
	}
	return t.LookupTree(id)
}

func (t *Tree) LookupTree(id int64) (*Tree, error) {
	e, err := t.LookupEntry(id)
	if err != nil {
		return nil, err
	}
	tree, ok := e.(*Tree)
	if !ok {
		return nil, fmt.Errorf("Entry with id %v (%v) found, but isn’t a tree.", id, helper.ToTimestamp(id))
	}
	return tree, nil
}

func (t *Tree) LookupEntryHash(hash string) (entry.Entry, error) {
	id, err := helper.ParseHash(hash)
	if err != nil {
		return nil, fmt.Errorf("LookupEntryHash: Couldn’t parse hash %v.", err)
	}
	return t.LookupEntry(id)
}

// Starting recursive function
func (t *Tree) LookupEntry(id int64) (entry.Entry, error) {
	return t.lookup([]*Tree{}, id)
}

// Recursive function
func (t *Tree) lookup(stack []*Tree, id int64) (entry.Entry, error) {
	if t.Id() == id {
		return t, nil
	}
	for _, e := range t.Entries() {
		if e.Id() == id {
			return e, nil
		}
	}
	for i, h := range t.Trees {
		if i == 0 {
			return h.lookup(append(stack, t.Trees[1:]...), id)
		}
	}
	for i, h := range stack {
		if i == 0 {
			return h.lookup(stack[1:], id)
		}
	}
	return nil, fmt.Errorf("lookupEntry: Id %v (%v) not found.", id, helper.ToTimestamp(id))
}

// search

func (t *Tree) SearchTree(slug, lang string) (*Tree, error) {
	return t.search([]*Tree{}, slug, lang)
}

func (t *Tree) search(stack []*Tree, slug, lang string) (*Tree, error) {
	if t.Slug(lang) == slug {
		return t, nil
	}
	//	for _, e := range t.Els {
	//		if e.Id() == id {
	//			return e, nil
	//		}
	//	}
	for i, h := range t.Trees {
		if i == 0 {
			return h.search(append(stack, t.Trees[1:]...), slug, lang)
		}
	}
	for i, h := range stack {
		if i == 0 {
			return h.search(stack[1:], slug, lang)
		}
	}
	return nil, fmt.Errorf("Couldn’t find slug %v in Tree.", slug)
}
