package parser

import (
	"e-book-manager/epub/epubReader"
	"errors"
)

func GetPublisher(metaData epubReader.Metadata) (*string, error) {
	if metaData.Publisher == nil || len(*metaData.Publisher) == 0 {
		return nil, errors.New("no publisher found")
	} else if len(*metaData.Publisher) > 1 {
		return nil, errors.New("multi publisher not supported")
	}
	pub := *metaData.Publisher
	return &pub[0].Text, nil
}
