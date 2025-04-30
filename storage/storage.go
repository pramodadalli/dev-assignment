package storage

import (
	"dev-assignment/models"
	"sync"
)

var (
	Users     = make(map[string]*models.User)
	UserFiles = make(map[string][]FileMeta)
	Mutex     = sync.Mutex{}
)

type FileMeta struct {
	Filename     string
	OriginalName string
	Size         int64
	UploadTime   string
}
