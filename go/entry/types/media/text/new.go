package text

import (
	"politics/go/entry"
	"politics/go/entry/helper"
	"politics/go/entry/parts/file"
	"politics/go/entry/parts/info"
	bf "gopkg.in/russross/blackfriday.v2"
	"time"
)

type Text struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	TextLangs map[string]string
	blank map[string]string
}

func NewText(path string, parent entry.Entry) (*Text, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "NewText",
	}

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	inf, langs, err := ReadTextFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	date, err := helper.ParseDate(inf["date"])
	if err != nil {
		date, err = helper.ParseDatePath(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
	}

	rendered := renderLangs(langs)

	return &Text{
		parent: parent,
		file:   file,

		date: date,
		info: inf,

		TextLangs: rendered,
		blank: langs,
	}, nil
}

func renderLangs(langs map[string]string) (map[string]string) {
	for _, l := range []string{"de", "en"} {
		text := string(bf.Run([]byte(langs[l]),bf.WithExtensions(bf.HardLineBreak)))
		langs[l] = text
	}
	return langs
}

func ReadTextFile(path string) (info.Info, map[string]string, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "readTextFile",
	}

	parts, err := splitTextFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, nil, fnErr
	}

	inf, err := info.UnmarshalInfo([]byte(parts["info"]))
	if err != nil {
		fnErr.Err = err
		return nil, nil, fnErr
	}

	delete(parts, "info")
	return inf, parts, nil
}
