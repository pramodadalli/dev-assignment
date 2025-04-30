package models

import "time"

type FileMeta struct {
	Filename     string
	Size         int64
	UploadTime   time.Time
	OriginalName string
}

var FileMetadata = map[string][]FileMeta{} // key: username
