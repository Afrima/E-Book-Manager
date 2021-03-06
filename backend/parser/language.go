package parser

import (
	"e-book-manager/epub/epubReader"
	"errors"
)

func GetLanguage(metaData epubReader.Metadata) (string, error) {
	if metaData.Language == nil || len(*metaData.Language) == 0 {
		return "", errors.New("lang not found")
	} else if len(*metaData.Language) > 1 {
		return "", errors.New("multi lang not supported")
	}
	lang := *metaData.Language
	return lang[0].Text, nil
}
