package download

import (
	"thor/internal/models"

	"github.com/anacrolix/torrent"
)

type IDownload interface {
	Close() error
	ListDownloads() []models.Download
	AddDownload(hashString string) (*torrent.Torrent, error)
}
