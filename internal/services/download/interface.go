package download

import (
	"github.com/anacrolix/torrent"
	"github.com/felipereyel/thor/internal/models"
)

type IDownload interface {
	Close() error
	ListDownloads() []models.Download
	AddDownload(hashString string) (*torrent.Torrent, error)
}
