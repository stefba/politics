package text

/*
import (
	"github.com/grokify/html-strip-tags-go"
	"politics/go/entry/hyph"
	"politics/go/entry/parser"
)

func markupLangs(langs map[string]string, paragraphstyle string) (map[string]string, error) {
	html := map[string]string{}
	for lang, text := range langs {
		h, err := parser.MarkupText(text, paragraphstyle)
		if err != nil {
			return nil, err
		}
		html[lang] = h
	}
	return html, nil
}

func stripLangs(langs map[string]string) (map[string]string, error) {
	blank := map[string]string{}
	for lang, text := range langs {
		b := strip.StripTags(text)
		blank[lang] = b
	}
	return blank, nil
}

func hyphenateLangs(langs map[string]string) (map[string]string, error) {
	hyphed := map[string]string{}
	for lang, text := range langs {
		h, err := hyph.HyphenateText(text, lang)
		if err != nil {
			return nil, err
		}
		hyphed[lang] = h
	}
	return hyphed, nil
}
*/
