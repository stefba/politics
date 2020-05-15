package set

import (
	"politics/go/entry"
	"politics/go/entry/helper"
	"politics/go/entry/helper/read"
	"politics/go/entry/helper/sort"
	"politics/go/entry/types/media"
)

func readEntries(path string, parent entry.Entry) (entry.Entries, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "readEntries",
	}

	files, err := read.GetFiles(path, false)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	entries, err := readEntryFiles(files, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	return sort.SortEntries(path, entries)
}

func readEntryFiles(files []*read.FileInfo, parent entry.Entry) (entry.Entries, error) {
	entries := entry.Entries{}
	for _, fi := range files {
		entry, err := media.NewMediaEntry(fi.Path, parent)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

