package main

// #cgo pkg-config: taglib
// #cgo LDFLAGS: -ltag_c
// #include <stdlib.h>
// #include <tag_c.h>
import "C"
import (
	"errors"
	"unsafe"
)

type musicTag struct {
	title     string
	performer string
	album     string
}

func ReadFileAndTag(filename string) (*musicTag, error) {

	cFileName := C.CString(filename)
	defer C.free(unsafe.Pointer(cFileName))

	cFile := C.taglib_file_new(cFileName)
	defer C.taglib_file_free(cFile)

	if cFile == nil {
		return nil, errors.New("invalid")
	}

	cTag := C.taglib_file_tag(cFile)

	cTitleName := C.taglib_tag_title(cTag)
	titleName := C.GoString(cTitleName)
	defer C.free(unsafe.Pointer(cTitleName))

	cPerformerName := C.taglib_tag_artist(cTag)
	performerName := C.GoString(cPerformerName)
	defer C.free(unsafe.Pointer(cPerformerName))

	cAlbumName := C.taglib_tag_album(cTag)
	albumName := C.GoString(cAlbumName)
	defer C.free(unsafe.Pointer(cAlbumName))

	return &musicTag{
		title:     titleName,
		performer: performerName,
		album:     albumName,
	}, nil
}
