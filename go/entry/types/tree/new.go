package tree

import (
	"politics/go/entry"
	"politics/go/entry/helper"
	"politics/go/entry/parts/file"
	"politics/go/entry/parts/info"
	"time"
)

type Tree struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	entries entry.Entries
	Trees   Trees
}

func (t *Tree) Copy() *Tree {
	return &Tree{
		parent: t.parent,
		file:   t.file,

		date: t.date,
		info: t.info,

		entries: t.entries,
		Trees:   t.Trees,
	}
}

type Trees []*Tree

func ReadTree(path string, parent *Tree) (*Tree, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "ReadTree",
	}

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	inf, err := readTreeInfo(path, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	date, err := helper.ParseDate(inf["date"])
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s := &Tree{
		parent: parent,
		file:   file,

		date: date,
		info: inf,
	}

	trees, err := readTrees(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	entries, err := readEntries(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s.entries = entries
	s.Trees = trees

	return s, nil
}

func readTreeInfo(path string, parent *Tree) (info.Info, error) {
	return readPolInfo(path, parent)
}
