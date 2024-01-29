package download

import (
	"thor/internal/models"
)

type IDownload interface {
	Close() error
	ListDownloads() []models.Download
	AddDownload(hashString string) (models.Download, error)
}
