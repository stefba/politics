package set

import (
	// "log"
	"politics/go/entry"
	"politics/go/entry/helper"
	"politics/go/entry/parts/file"
	"politics/go/entry/parts/info"
	"time"
)

type Set struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	entries entry.Entries
	//Cover *Image
}

func (s *Set) Copy() *Set {
	return &Set{
		parent: s.parent,
		file:   s.file,

		date: s.date,
		info: s.info,

		entries: s.entries,
	}
}

type Sets []*Set

func NewSet(path string, parent entry.Entry) (*Set, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "NewSet",
	}

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	info, err := info.ReadDirInfo(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	date, err := helper.ParseDate(info["date"])
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s := &Set{
		parent: parent,
		file:   file,

		date: date,
		info: info,
	}

	entries, err := readEntries(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s.entries = entries

	return s, nil
}

/*
	cover, err := ReadCover(path, h)
	if err != nil {
		// log.Println(err)
	}
*/
