package models

import (
	"os"
	"sync"
	"time"
)

type FileInfo struct {
	sync.Mutex
	Name          string
	NameExt       string
	Size          float64
	Mode          os.FileMode
	ModTime       time.Time
	IsDir         bool
	NameSize      string
	Srt           int
	SizeSrt       float64
	NameTailleSrt string
}
