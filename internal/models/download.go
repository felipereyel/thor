package models

import "github.com/anacrolix/torrent"

type Download struct {
	Name   string
	Hash   string
	Status string
}

func DownloadFromTorrent(torr *torrent.Torrent) Download {
	download := Download{
		Name: torr.Name(),
		Hash: torr.InfoHash().HexString(),
	}

	if torr.Complete.Bool() {
		download.Status = "Complete"
	} else {
		download.Status = "Downloading"
	}

	return download
}
