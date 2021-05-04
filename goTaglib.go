package main

import (
	"errors"
	"os"
)

import "github.com/dhowden/tag"

type musicTag struct {
	title     string
	performer string
	album     string
}

func ReadFileAndTag(filename string) (*musicTag, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("error opening : " + filename)
	}
	defer file.Close()
	tags, err2 := tag.ReadFrom(file)
	if err2 != nil {
		return nil, errors.New("error reading tag : " + filename)
	}

	return &musicTag{
		title:     tags.Title(),
		performer: tags.Artist(),
		album:     tags.Album(),
	}, nil
}
